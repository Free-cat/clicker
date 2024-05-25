import { createRouter, createWebHistory } from 'vue-router';
import GameUI from '../components/GameUI.vue';
import BoostPage from '../components/BoostPage.vue';

const routes = [
    {
        path: '/',
        name: 'GameUI',
        component: GameUI
    },
    {
        path: '/boost',
        name: 'BoostPage',
        component: BoostPage
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;
