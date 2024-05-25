const lc = window.location;
const wsStart = lc.protocol === 'https:' ? 'wss://' : 'ws://';
// const port = lc.port ? `:${lc.port}` : '8081';

const state = {
    tokensLeft: 0,
    tokensLimit: 500,
    totalClicks: 0,
    dailyClicks: 0,
    level: 1,
    progress: 0,
    levelImage: '',
    userId: null,
    energy: 956,
    energyLimit: 1500,
    socket: null,  // Initialize socket as null
    reconnectAttempts: 0
};

const mutations = {
    SET_SOCKET(state, socket) {
        state.socket = socket;
    },
    SET_INITIAL_DATA(state, data) {
        state.tokensLeft = data.tokens_left;
        state.tokensLimit = data.tokens_limit;
        state.totalClicks = data.total_clicks;
        state.dailyClicks = data.daily_clicks;
        state.level = data.level;
        state.progress = data.progress;
        state.levelImage = data.level_image;
        state.energy = data.energy;
        state.energyLimit = data.energy_limit;
    },
    INCREMENT_TOKEN(state) {
        // Select max tokens limit based on level
        const maxTokens = state.tokensLeft + 1;

        state.tokensLeft = Math.min(maxTokens, state.tokensLimit);
    },
    UPDATE_DATA(state, data) {
        state.tokensLeft = data.tokens_left;
        state.totalClicks = data.total_clicks;
        state.dailyClicks = data.daily_clicks;
        state.level = data.level;
        state.progress = data.progress;
        state.levelImage = data.level_image;
        state.energy = data.energy;
        state.energyLimit = data.energy_limit;
        state.progress = data.progress;
    },
    SET_USER_ID(state, userId) {
        state.userId = userId;
    },
    DECREMENT_TOKENS(state) {
        if (state.tokensLeft > 0) {
            state.tokensLeft--;
        }
    },
    BOOST_ENERGY(state) {
        state.energy += 100; // Example increment, adjust as needed
        if (state.energy > state.energyLimit) {
            state.energy = state.energyLimit;
        }
    },
    RESET_SOCKET(state) {
        state.socket = null;
    },
    INCREMENT_RECONNECT_ATTEMPTS(state) {
        state.reconnectAttempts++;
    },
    RESET_RECONNECT_ATTEMPTS(state) {
        state.reconnectAttempts = 0;
    }
};

const actions = {
    connectWebSocket({ commit, state, dispatch }) {
        return new Promise((resolve, reject) => {
            if (state.socket) {
                console.log('WebSocket already connected');
                resolve(); // WebSocket already connected
                return;
            }

            let userId = state.userId || prompt("Please enter your id", "1");
            commit('SET_USER_ID', userId);

            // const socket = new WebSocket(`${wsStart}${loc.host}/ws`);
            const socket = new WebSocket(`${wsStart}${lc.hostname}:8081/ws`);
            // const socket = new WebSocket('ws://localhost:8081/ws'); // Hardcoded URL for local testing

            socket.onopen = () => {
                console.log('Connected to WebSocket');
                const initialMessage = { user_id: userId };
                socket.send(JSON.stringify(initialMessage));
                commit('SET_SOCKET', socket);
                commit('RESET_RECONNECT_ATTEMPTS');
                resolve(); // Resolve the promise after connection
            };

            socket.onmessage = (event) => {
                const data = JSON.parse(event.data);
                if (data.action === "increment_token") {
                    commit('INCREMENT_TOKEN');
                } else if (data.status === "initial_data") {
                    commit('SET_INITIAL_DATA', data);
                } else {
                    commit('UPDATE_DATA', data);
                }
            };

            const handleReconnect = () => {
                commit('RESET_SOCKET');
                let reconnectAttempts = state.reconnectAttempts;
                if (reconnectAttempts < 10) { // Limit to 10 attempts
                    commit('INCREMENT_RECONNECT_ATTEMPTS');
                    let timeout = Math.min(1000 * Math.pow(2, reconnectAttempts), 30000); // Exponential backoff with a max delay of 30 seconds
                    setTimeout(() => {
                        console.log(`Reconnecting... Attempt ${reconnectAttempts + 1}`);
                        dispatch('connectWebSocket').catch(err => console.error('Reconnect failed', err));
                    }, timeout);
                }
            };

            socket.onclose = () => {
                console.log('Disconnected from WebSocket');
                handleReconnect();
            };

            socket.onerror = (error) => {
                console.error('WebSocket error:', error);
                handleReconnect();
                reject(error);
            };
        });
    },
    incrementToken({ commit, state }) {
        if (state.tokensLeft > 0) {
            commit('INCREMENT_TOKEN');
        } else {
            alert('No more tokens left!');
        }
    },
    sendClickEvent({ commit, state }) {
        if (state.tokensLeft > 0) {
            commit('DECREMENT_TOKENS');
            const message = { user_id: state.userId, action: "click" };
            if (state.socket && state.socket.readyState === WebSocket.OPEN) {
                state.socket.send(JSON.stringify(message));
            } else {
                console.error('WebSocket is not open');
            }
        } else {
            alert('No more tokens left!');
        }
    },
    boostEnergy({ commit }) {
        commit('BOOST_ENERGY');
    }
};

const getters = {
    tokensLeft: state => state.tokensLeft,
    tokensLimit: state => state.tokensLimit,
    totalClicks: state => state.totalClicks,
    dailyClicks: state => state.dailyClicks,
    level: state => state.level,
    progress: state => state.progress,
    levelImage: state => state.levelImage,
    energy: state => state.energy,
    energyLimit: state => state.energyLimit,
    socket: state => state.socket // Getter for socket
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
