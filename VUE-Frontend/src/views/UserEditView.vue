<template>
  <div class="user-edit-view">
    <div class="header-section">
      <div class="back-button">
        <button @click="$router.back()" class="btn btn-secondary">
          <i class="icon-arrow-left"></i> Back
        </button>
      </div>
      <h1 class="page-title">Edit User</h1>
    </div>

    <div v-if="loading" class="loading">
      <p>Loading user data...</p>
    </div>

    <div v-else-if="error" class="error">
      <p>{{ error }}</p>
      <button @click="loadUser" class="btn btn-primary">Retry</button>
    </div>

    <UserEditForm 
      v-else-if="user"
      :userData="user"
      @updated="handleUserUpdated"
      @cancelled="handleCancel"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getUser } from '../services/userService'
import UserEditForm from '../components/UserEditForm.vue'

const route = useRoute()
const router = useRouter()

const user = ref(null)
const loading = ref(true)
const error = ref('')

const userId = route.params.id

const loadUser = async () => {
  try {
    loading.value = true
    error.value = ''
    const response = await getUser(userId)
    user.value = response.data
  } catch (err) {
    console.error('Error loading user:', err)
    error.value = 'Failed to load user data: ' + (err.response?.data?.error || err.message)
  } finally {
    loading.value = false
  }
}

const handleUserUpdated = (updatedUser) => {
  // Show success message
  alert('User updated successfully!')
  
  // Navigate back to user detail view or user list
  router.push(`/users/${userId}`)
}

const handleCancel = () => {
  // Navigate back to user detail view
  router.push(`/users/${userId}`)
}

onMounted(() => {
  loadUser()
})
</script>

<style scoped>
.user-edit-view {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.header-section {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 2rem;
}

.back-button .btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background-color: #6c757d;
  color: white;
  border: none;
  border-radius: 4px;
  text-decoration: none;
  cursor: pointer;
  transition: background-color 0.3s;
}

.back-button .btn:hover {
  background-color: #545b62;
}

.page-title {
  color: #333;
  margin: 0;
}

.loading,
.error {
  text-align: center;
  padding: 2rem;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.error {
  color: #dc3545;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background-color: #545b62;
}

@media (max-width: 768px) {
  .user-edit-view {
    padding: 1rem;
  }

  .header-section {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
