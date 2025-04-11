import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { useDeviceStore } from '@/stores/device'
import { checkAuth, isLocal } from '@/services/ApiService'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/panels',
      name: 'panels',
      component: () => import('../views/PanelsView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/panel/edit/:dirname',
      name: 'panel-edit',
      component: () => import('../views/PanelsView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/panel/view/:dirname',
      name: 'panel-view',
      component: () => import('../views/PanelsView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/macros',
      name: 'macros',
      component: () => import('../views/MacrosView.vue'),
      meta: { localOnly: true },
    },
    {
      path: '/devices',
      name: 'devices',
      component: () => import('../views/DevicesView.vue'),
    },
    // {
    //   path: '/settings',
    //   name: 'settings',
    //   component: () => import('../views/SettingsView.vue'),
    // },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue'),
    // },
  ],
})

router.beforeEach(async (to, from, next) => {
  const auth = await checkAuth()

  if (to.meta.requiresAuth && !auth && !isLocal()) next('/devices')
  else if (to.meta.localOnly && !isLocal()) next('/')
  else next()
})

export default router
