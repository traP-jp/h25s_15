import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue' // 仮

const routes = [{ path: '/', name: 'home', component: HomePage }]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
