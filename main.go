package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
	upgrader    = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins for now
		},
	}
	maxTokens               = 500
	refillInterval          = 1 * time.Second
	globalLastRefillTimeKey = "global_last_refill_time"
	clients                 = make(map[*websocket.Conn]bool)
	clientsMutex            sync.Mutex
	// fibonacci levels
	levels = []int{0, 10, 20, 30, 50, 80, 130, 210, 340, 550, 890, 1440, 2330, 3770, 6100, 9870}
	// fibonacci levels images delimeted by 4 levels
	levelsImages = []string{"/images/level1.webp", "/images/level1.webp", "/images/level1.webp", "/images/level1.webp"}
)

type ClickResponse struct {
	Status        string  `json:"status"`
	TokensLeft    int     `json:"tokens_left"`
	TokensLimit   int     `json:"tokens_limit"`
	TotalClicks   int     `json:"total_clicks"`
	DailyClicks   int     `json:"daily_clicks"`
	AllTimeClicks int     `json:"all_time_clicks"`
	Level         int     `json:"level"`
	LevelProgress float64 `json:"progress"`
	LevelImage    string  `json:"level_image"`
}

func currentTimeSeconds() int64 {
	return time.Now().Unix()
}

func refillTokens(userID string) {
	userTokensExists, _ := redisClient.Exists(ctx, userID+"_tokens").Result()
	currentTokens, _ := redisClient.Get(ctx, userID+"_tokens").Int()
	lastRefillTime, _ := redisClient.Get(ctx, userID+"_last_refill_at").Int64()

	fmt.Println("User tokens exists", userTokensExists)
	fmt.Println("Current tokens", currentTokens)
	fmt.Println("Last refill time", lastRefillTime)

	elapsedIntervals := (currentTimeSeconds() - lastRefillTime) / int64(refillInterval.Seconds())
	if elapsedIntervals > 0 {
		newTokens := currentTokens + int(elapsedIntervals)
		if newTokens > maxTokens {
			newTokens = maxTokens
		}
		redisClient.Set(ctx, userID+"_tokens", newTokens, 0)
		redisClient.Set(ctx, userID+"_last_refill_at", currentTimeSeconds(), 0)
	}
}

func handleClick(userID string) ClickResponse {
	refillTokens(userID)

	pipe := redisClient.TxPipeline()
	currentTokens := pipe.Get(ctx, userID+"_tokens")
	totalClicks := pipe.Get(ctx, userID+"_clicks")
	dailyClicks := pipe.Get(ctx, userID+"_"+time.Now().Format("2006-01-02")+"_clicks")
	allTimeClicks := pipe.Get(ctx, userID+"_all_time_clicks")

	_, err := pipe.Exec(ctx)

	if err != nil {
		log.Println("Error executing Redis pipeline:", err)
	}

	tokens, _ := currentTokens.Int()
	clicks, _ := totalClicks.Int()
	dailyClicksCount, _ := dailyClicks.Int()
	allTimeClicksCount, _ := allTimeClicks.Int()

	if tokens < 0 {
		tokens = 0
	}

	if tokens > 0 {
		pipe.Decr(ctx, userID+"_tokens")
		pipe.Incr(ctx, userID+"_clicks")
		pipe.Incr(ctx, userID+"_all_time_clicks")
		pipe.Incr(ctx, userID+"_"+time.Now().Format("2006-01-02")+"_clicks")
		_, err = pipe.Exec(ctx)
		if err != nil {
			log.Println("Error executing Redis pipeline:", err)
			return ClickResponse{Status: "error", TokensLeft: 0, TotalClicks: 0}
		}
		levelIndex := getLevelIndex(clicks)
		levelImage := getLevelImage(levelIndex)
		levelProgress := getLevelProgress(clicks, levelIndex)

		return ClickResponse{Status: "success",
			TokensLeft:    tokens - 1,
			TotalClicks:   clicks + 1,
			Level:         levelIndex,
			LevelProgress: levelProgress,
			LevelImage:    levelImage,
			DailyClicks:   dailyClicksCount,
			AllTimeClicks: allTimeClicksCount,
		}
	}
	return ClickResponse{Status: "error", TokensLeft: tokens, TotalClicks: clicks}
}

func getLevelImage(levelIndex int) string {
	if levelIndex < 4 {
		return levelsImages[0]
	}
	if levelIndex < 8 {
		return levelsImages[1]
	}
	if levelIndex < 12 {
		return levelsImages[2]
	}

	return levelsImages[3]
}

// Progress increment when user clicks
func getLevelProgress(tokens int, levelIndex int) float64 {
	if levelIndex == 0 {
		return 0
	}
	tokensToEarnNewLevel := levels[levelIndex] - levels[levelIndex-1]

	return float64(tokens) / float64(tokensToEarnNewLevel) * 100
}

func getLevelIndex(tokens int) int {
	for i, level := range levels {
		if tokens < level {
			return i
		}
	}
	return len(levels)
}

func sendInitialData(conn *websocket.Conn, userID string) {
	refillTokens(userID)

	pipe := redisClient.TxPipeline()
	currentTokens := pipe.Get(ctx, userID+"_tokens")
	totalClicks := pipe.Get(ctx, userID+"_clicks")
	dailyClicks := pipe.Get(ctx, userID+"_"+time.Now().Format("2006-01-02")+"_clicks")
	allTimeClicks := pipe.Get(ctx, userID+"_all_time_clicks")

	_, err := pipe.Exec(ctx)

	if err != nil {
		log.Println("Error executing Redis pipeline:", err)
	}

	tokens, _ := currentTokens.Int()
	clicks, _ := totalClicks.Int()
	dailyClicksCount, _ := dailyClicks.Int()
	allTimeClicksCount, _ := allTimeClicks.Int()
	level := getLevelIndex(clicks)
	levelImage := getLevelImage(level)
	levelProgress := getLevelProgress(clicks, level)

	initialData := ClickResponse{
		Status:        "initial_data",
		TokensLeft:    tokens,
		TokensLimit:   maxTokens,
		TotalClicks:   clicks,
		Level:         level,
		LevelProgress: levelProgress,
		LevelImage:    levelImage,
		DailyClicks:   dailyClicksCount,
		AllTimeClicks: allTimeClicksCount,
	}

	responseJSON, err := json.Marshal(initialData)
	if err != nil {
		log.Println("JSON marshal error:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, responseJSON); err != nil {
		log.Println("Write error:", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Read error:", err)
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		return
	}

	var data map[string]string
	if err := json.Unmarshal(message, &data); err != nil {
		log.Println("JSON unmarshal error:", err)
		return
	}

	userID, ok := data["user_id"]
	if !ok {
		log.Println("No user_id in message")
		return
	}

	sendInitialData(conn, userID)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			clientsMutex.Lock()
			delete(clients, conn)
			clientsMutex.Unlock()
			break
		}

		if err := json.Unmarshal(message, &data); err != nil {
			log.Println("JSON unmarshal error:", err)
			continue
		}

		userID, ok := data["user_id"]
		if !ok {
			log.Println("No user_id in message")
			continue
		}

		response := handleClick(userID)
		responseJSON, err := json.Marshal(response)
		if err != nil {
			log.Println("JSON marshal error:", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, responseJSON); err != nil {
			log.Println("Write error:", err)
			clientsMutex.Lock()
			delete(clients, conn)
			clientsMutex.Unlock()
			break
		}
	}
}

func broadcastTokenIncrement() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		incrementMessage := map[string]string{"action": "increment_token"}
		message, _ := json.Marshal(incrementMessage)

		clientsMutex.Lock()
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Broadcast error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMutex.Unlock()
	}
}

func main() {
	redisOpts, err := redis.ParseURL(os.Getenv("REDISCLOUD_URL"))
	if err != nil {
		log.Fatalf("Could not parse REDIS_URL: %v", err)
	}
	redisClient = redis.NewClient(redisOpts)

	indexPage := http.FileServer(http.Dir("./public"))
	http.Handle("/", indexPage)

	go broadcastTokenIncrement()

	http.HandleFunc("/ws", wsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Println("Server started at " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
