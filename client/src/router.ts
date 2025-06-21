import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue' // 仮
import GamePage from './pages/GamePage.vue' // 仮

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/game/:gameId', name: 'game', component: GamePage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
