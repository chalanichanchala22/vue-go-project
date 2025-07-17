<template>
  <div class="app">
    <nav v-if="isAuthenticated" class="navbar">
      <div class="nav-container">
        <div class="nav-brand">
          <router-link to="/users" class="brand-link">User Management</router-link>
        </div>
        <div class="nav-menu">
          <router-link to="/users" class="nav-link">Users</router-link>
          <router-link to="/userlist" class="nav-link">User List</router-link>
          <button @click="handleLogout" class="logout-btn">Logout</button>
        </div>
      </div>
    </nav>
    
    <main class="main-content">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import authService from './services/authService'

const router = useRouter()
const isAuthenticated = ref(false)

const checkAuth = () => {
  isAuthenticated.value = authService.isAuthenticated()
}

const handleLogout = () => {
  authService.logout()
  isAuthenticated.value = false
  router.push('/login')
}

onMounted(() => {
  checkAuth()
  
  // Listen for route changes to update auth status
  router.afterEach(() => {
    checkAuth()
  })
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background-color: #f8fafc;
}

.app {
  min-height: 100vh;
}

.navbar {
  background-color: white;
  border-bottom: 1px solid #e2e8f0;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
}

.nav-brand .brand-link {
  font-size: 1.25rem;
  font-weight: 700;
  color: #1f2937;
  text-decoration: none;
}

.nav-menu {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.nav-link {
  color: #4b5563;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: background-color 0.15s ease;
}

.nav-link:hover,
.nav-link.router-link-active {
  background-color: #f3f4f6;
  color: #1f2937;
}

.logout-btn {
  background-color: #ef4444;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.875rem;
  transition: background-color 0.15s ease;
}

.logout-btn:hover {
  background-color: #dc2626;
}

.main-content {
  padding: 0;
}

/* Remove default padding for login page */
.main-content:has(.login-container) {
  padding: 0;
}
</style>
