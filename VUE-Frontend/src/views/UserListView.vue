<template>
  <div class="user-list-view">
    <div class="header">
      <h2>User Details</h2>
      <button @click="goBack" class="btn-back">
        ← Back to Add User
      </button>
    </div>
    <div class="user-list-container">
      <h3>Select a user to view details:</h3>
      <UserList :users="users" @deleted="loadUsers" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import UserList from '../components/UserDetails.vue'
import { getUsers } from '../services/userService'

const router = useRouter()
const users = ref([])

const loadUsers = async () => {
  const res = await getUsers()
  users.value = res.data
}

const goBack = () => {
  router.push('/users')
}

onMounted(loadUsers)
</script>

<style scoped>
.user-list-view {
  padding: 1rem;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 2px solid #e0e0e0;
}

.header h2 {
  margin: 0;
  color: #333;
}

.btn-back {
  background-color: #6c757d;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: background-color 0.2s;
}

.btn-back:hover {
  background-color: #5a6268;
}

.user-list-container {
  background-color: #f8f9fa;
  border: 2px solid #4CAF50;
  border-radius: 8px;
  padding: 1.5rem;
}

.user-list-container h3 {
  color: #333;
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  font-weight: 600;
  text-align: center;
  background-color: #4CAF50;
  color: white;
  padding: 0.75rem;
  border-radius: 4px;
  margin: -1.5rem -1.5rem 1rem -1.5rem;
}

@media (max-width: 768px) {
  .user-list-view {
    padding: 0.5rem;
  }
  
  .header {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }
  
  .user-list-container {
    padding: 1rem;
  }
  
  .user-list-container h3 {
    margin: -1rem -1rem 1rem -1rem;
  }
}
</style>
