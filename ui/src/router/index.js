import { createRouter, createWebHistory } from 'vue-router'
import BackenLayout from '../views/backend/BackendLayout.vue'
import FrontendLayout from '../views/frontend/FrontendLayout.vue'
import { state } from '@/stores/app'

const router = createRouter({
  // /backend/page
  // /backend#page1
  // https://cn.vitejs.dev/guide/env-and-mode.html
  // https://router.vuejs.org/zh/guide/essentials/nested-routes.html
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
      component: BackenLayout,
      children: [
        {
          // blogs 相对路径 /backend/blogs
          // /blogs 绝对路径
          path: 'blogs/list',
          name: 'BackendListBlog',
          component: () => import('../views/backend/blog/ListView.vue')
        },
        {
          //   /backend/blogs/22  id=22
          path: 'blogs/detail/:id',
          name: 'BackendDetailBlog',
          component: () => import('../views/backend/blog/DetailView.vue')
        },
        {
          //   /backend/blogs/22  id=22
          path: 'blogs/edit/:id',
          name: 'BackendEditBlog',
          component: () => import('../views/backend/blog/EditView.vue')
        },
        {
          // blogs 相对路径 /backend/blogs
          // /blogs 绝对路径
          path: 'comments/list',
          name: 'BackendListComment',
          component: () => import('../views/backend/comment/ListView.vue')
        }
      ]
    },
    {
      path: '/frontend',
      name: 'FrontendLayout',
      component: FrontendLayout
    }
  ]
})

router.beforeEach((to) => {
  // 是不是访问后台页面
  if (to.fullPath.startsWith('/backend/')) {
    // 判断是否登录
    if (!state.value.token) {
      // 跳转去登录页面
      return { name: 'LoginView' }
    }
  }
})

export default router
