import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/crawl',
      name: 'crawl',
      component: () => import('../views/crawl.vue')
    }
  ]
})

export default router
