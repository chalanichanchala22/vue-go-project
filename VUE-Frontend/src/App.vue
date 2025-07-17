<template>
  <div class="app">
    <Navbar v-if="isAuthenticated" @logout="handleLogout" />
    
    <main class="main-content">
      <router-view />
    </main>
    
    <Footer v-if="isAuthenticated" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import authService from './services/authService'
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'

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
  display: flex;
  flex-direction: column;
}

.main-content {
  padding: 0;
  flex: 1;
}

/* Remove default padding for login page */
.main-content:has(.login-container) {
  padding: 0;
}
</style>
