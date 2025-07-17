<template>
  <div class="login-container">
    <div class="login-card">
      <h2 class="login-title">Welcome Back</h2>
      <p class="login-subtitle">Sign in to your account</p>
      
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label for="email">Email</label>
          <input
            id="email"
            v-model="email"
            type="email"
            required
            placeholder="Enter your email"
            :disabled="loading"
          />
        </div>
        
        <div class="form-group">
          <label for="password">Password</label>
          <input
            id="password"
            v-model="password"
            type="password"
            required
            placeholder="Enter your password"
            :disabled="loading"
          />
        </div>
        
        <div v-if="error" class="error-message">
          {{ error }}
        </div>
        
        <button
          type="submit"
          class="login-button"
          :disabled="loading"
        >
          <span v-if="loading">Signing in...</span>
          <span v-else>Sign In</span>
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import authService from '../services/authService'

const router = useRouter()

const email = ref('admin@example.com')
const password = ref('password123')
const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
  if (!email.value || !password.value) {
    error.value = 'Please fill in all fields'
    return
  }

  loading.value = true
  error.value = ''

  try {
    await authService.login(email.value, password.value)
    
    // Redirect to dashboard
    router.push('/dashboard')
  } catch (err) {
    error.value = err.message || 'Login failed. Please try again.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1e293b 0%, #334155 25%, #475569 50%, #64748b 75%, #94a3b8 100%);
  padding: 20px;
}

.login-card {
  background: linear-gradient(145deg, #1f2937 0%, #374151 25%, #4b5563 50%, #6b7280 75%, #9ca3af 100%);
  padding: 2.5rem;
  border-radius: 16px;
  box-shadow: 0 10px 40px rgba(31, 41, 55, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.2);
  width: 100%;
  max-width: 450px;
  position: relative;
  overflow: hidden;
}

.login-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0.05) 100%);
  backdrop-filter: blur(10px);
  z-index: 0;
}

.login-card > * {
  position: relative;
  z-index: 1;
}

.login-title {
  font-size: 2rem;
  font-weight: 700;
  text-align: center;
  color: #ffffff;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.5);
  margin-bottom: 0.5rem;
  letter-spacing: 0.5px;
}

.login-subtitle {
  text-align: center;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 2.5rem;
  font-size: 1.1rem;
  font-weight: 500;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1.8rem;
}

.form-group {
  margin-bottom: 1.8rem;
}

.form-group label {
  display: block;
  font-size: 1rem;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 0.6rem;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);
  letter-spacing: 0.3px;
}

.form-group input {
  width: 100%;
  padding: 0.9rem 1rem;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.3s ease;
  background-color: rgba(255, 255, 255, 0.95);
  color: #1f2937;
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  box-sizing: border-box;
}

.form-group input:focus {
  outline: none;
  border-color: #ffffff;
  box-shadow: 0 0 0 3px rgba(255, 255, 255, 0.3);
  background-color: #ffffff;
  transform: translateY(-1px);
}

.form-group input:hover {
  border-color: rgba(255, 255, 255, 0.6);
  background-color: #ffffff;
}

.form-group input:disabled {
  background-color: rgba(249, 250, 251, 0.8);
  color: #6b7280;
  cursor: not-allowed;
  border-color: rgba(209, 213, 219, 0.6);
}

.error-message {
  background: linear-gradient(135deg, #fef2f2 0%, #fee2e2 100%);
  border: 2px solid #fca5a5;
  color: #991b1b;
  padding: 1rem 1.5rem;
  border-radius: 8px;
  font-size: 1rem;
  margin-bottom: 1.5rem;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(220, 38, 38, 0.2);
  text-align: center;
  letter-spacing: 0.2px;
}

.login-button {
  width: 100%;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%);
  color: #1f2937;
  padding: 1rem 1.5rem;
  border: 2px solid rgba(255, 255, 255, 0.8);
  border-radius: 10px;
  font-size: 1.1rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  letter-spacing: 0.3px;
}

.login-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.login-button:hover:not(:disabled) {
  background: linear-gradient(135deg, #f8fafc 0%, #e5e7eb 100%);
  color: #111827;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
  border-color: #ffffff;
}

.login-button:hover:not(:disabled)::before {
  left: 100%;
}

.login-button:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #6b7280 100%);
  color: #ffffff;
  cursor: not-allowed;
  transform: none;
  box-shadow: 0 2px 8px rgba(156, 163, 175, 0.2);
  border-color: #9ca3af;
}

.login-button:active:not(:disabled) {
  transform: translateY(0);
}

/* Responsive design */
@media (max-width: 480px) {
  .login-container {
    padding: 1rem;
  }
  
  .login-card {
    padding: 2rem 1.5rem;
  }
  
  .login-title {
    font-size: 1.7rem;
  }
  
  .form-group input {
    padding: 0.8rem;
  }
  
  .login-button {
    padding: 0.8rem 1.2rem;
    font-size: 1rem;
  }
}
</style>
