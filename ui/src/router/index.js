import { createRouter, createWebHistory } from 'vue-router'
import BackenLayout from '../views/backend/BackendLayout.vue'
import FrontendLayout from '../views/frontend/FrontendLayout.vue'

const router = createRouter({
  // /backend/page
  // /backend#page1
  // https://cn.vitejs.dev/guide/env-and-mode.html
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'LoginView',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/backend',
      name: 'BackendLayout',
      component: BackenLayout
    },
    {
      path: '/frontend',
      name: 'FrontendLayout',
      component: FrontendLayout
    }
  ]
})

export default router
