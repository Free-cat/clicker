# Clicker APP

## Description

There is a simple clicker app. You can click on the button and increase the counter. 
The counter is stored in the Redis database. You can see how much you're clicking in real-time.
Every second energy for clicking is add +1, and you can't click if you don't have enough energy.
You can't get more energy than set as limit.

## Prerequisites

Before you begin, ensure you have met the following requirements:

* You have installed the latest version of `Go`, `Node.js`, `npm`, `yarn`, and `Docker`.
* You have a `MacOS` machine. This guide is tailored for `MacOS`.

## Installing Backend

To install and start by local Golang, follow these steps:
```bash
# Clone the repository

# Navigate into the directory
cd clicker

# Install backend dependencies
go mod download
```

To install and start by Docker, follow these steps:
```bash
# Clone the repository

# Navigate into the directory
cd clicker

# Build the docker compose
docker-compose up --build

# Run the docker compose
docker-compose up

# Access the Websocket server
ws://localhost:8081/ws
```

## Installing Frontend

```bash
# Navigate into the frontend directory
cd frontend

# Install frontend dependencies
npm install

# Start the frontend
npm run serve
```

## Using Clicker APP

To use Clicker APP, follow these steps:
1. Open the browser and navigate to `http://localhost:8080/`.
2. Put user id to alert and click `OK`.
3. Click on the button to increase the counter.
4. You can see the counter in real-time.
5. Every second energy for clicking is add +1.
6. You can't click if you don't have enough energy.