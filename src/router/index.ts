import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/HomePage.vue'
import About from '../views/AboutPage.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    meta: {
      template: 'BasicLayout',
    },
  },
  {
    path: '/about',
    name: 'about',
    component: About,
    meta: {
      template: 'BasicLayout',
    },
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
