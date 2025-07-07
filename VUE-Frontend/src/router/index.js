import { createRouter, createWebHistory } from 'vue-router'
import UsersView from '../views/UsersView.vue'
import UserDetailView from '../views/UserDetailView.vue'

const routes = [
  { path: '/', redirect: '/users' },
  { path: '/users', component: UsersView },
  { path: '/users/:id', component: UserDetailView },

]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
