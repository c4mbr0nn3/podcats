import { createRouter, createWebHashHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import routes from './routes'

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

// redirect to login page if not logged in and trying to access a restricted page
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const isAuthenticated = authStore.isAuthenticated
  const needPasswordChange = authStore.needPasswordChange
  const isRoot = authStore.isRoot
  const authRequired = to.meta.requiresAuth
  const requiresRoot = to.meta.requiresRoot

  if (authRequired && !isAuthenticated)
    return next({ name: 'login' })

  if (authRequired && needPasswordChange && to.name !== 'change-password')
    return next({ name: 'change-password' })

  if (requiresRoot && !isRoot)
    return next({ name: 'home' })

  next()
})

export default router
