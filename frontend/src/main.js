import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // Import the router
import store from './store';
import './main.css'; // Import the main.css file

const app = createApp(App);

app.use(router);
app.use(store);

router.isReady().then(() => {
    app.mount('#app');
});
