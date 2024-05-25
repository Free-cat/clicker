// store/index.js
import { createStore } from 'vuex';
import websocket from './modules/websocket';

export default createStore({
    modules: {
        websocket
    }
});
