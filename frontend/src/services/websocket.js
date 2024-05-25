class WebSocketService {
    constructor(url) {
        this.url = url;
        this.socket = null;
    }

    connect() {
        this.socket = new WebSocket(this.url);
        this.socket.onopen = () => {
            console.log('WebSocket connection established');
        };
        this.socket.onmessage = (event) => {
            console.log('WebSocket message received:', event.data);
        };
        this.socket.onclose = () => {
            console.log('WebSocket connection closed');
        };
        this.socket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    }

    send(message) {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(message);
        } else {
            console.error('WebSocket is not open. Ready state:', this.socket.readyState);
        }
    }

    close() {
        if (this.socket) {
            this.socket.close();
        }
    }
}

export default new WebSocketService('wss://your-websocket-url');
