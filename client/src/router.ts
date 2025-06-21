import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue' // 仮
import ResultPage from './pages/ResultPage.vue'

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/result/:gameId', name: 'resultpage', component: ResultPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
