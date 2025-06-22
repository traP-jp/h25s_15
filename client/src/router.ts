import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue' // 仮
import ResultPage from './pages/ResultPage.vue'
import GamePage from './pages/GamePage.vue' // 仮

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/result/:gameId', name: 'resultpage', component: ResultPage, props: true },
  { path: '/game/:gameId', name: 'game', component: GamePage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
