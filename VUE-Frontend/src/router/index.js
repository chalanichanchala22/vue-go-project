import { createRouter, createWebHistory } from 'vue-router'
import UsersView from '../views/UsersView.vue'
import UserDetailView from '../views/UserDetailView.vue'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/DashboardView.vue'
import authService from '../services/authService'

const routes = [
  { 
    path: '/', 
    redirect: () => {
      return authService.isAuthenticated() ? '/dashboard' : '/login'
    }
  },
  { 
    path: '/login', 
    component: LoginView,
    meta: { requiresGuest: true }
  },
  { 
    path: '/dashboard', 
    component: DashboardView,
    meta: { requiresAuth: true }
  },
  { 
    path: '/users', 
    component: UsersView,
    meta: { requiresAuth: true }
  },
  { 
    path: '/users/:id', 
    component: UserDetailView,
    meta: { requiresAuth: true }
  },
  { 
    path: '/users/:id/edit', 
    component: () => import('../views/UserEditView.vue'),
    meta: { requiresAuth: true }
  },
  { 
    path: '/userlist', 
    component: () => import('../views/UserListView.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guards
router.beforeEach((to, from, next) => {
  const isAuthenticated = authService.isAuthenticated()
  
  if (to.meta.requiresAuth && !isAuthenticated) {
    // Redirect to login if trying to access protected route
    next('/login')
  } else if (to.meta.requiresGuest && isAuthenticated) {
    // Redirect to dashboard if trying to access login while authenticated
    next('/dashboard')
  } else {
    next()
  }
})

export default router
