/*
Macrame is a program that enables the user to create keyboard macros and button panels. 
The macros are saved as simple JSON files and can be linked to the button panels. The panels can 
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation, either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

import { createRouter, createWebHistory } from 'vue-router'
import DashboardView from '../views/DashboardView.vue'
import { checkAuth, isLocal } from '@/services/ApiService'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: DashboardView,
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
