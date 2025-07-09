<template>
  <div class="users-with-phones-view">
    <div class="header">
      <h2>Users with Phone Details</h2>
      <div class="header-actions">
        <button @click="toggleView" class="btn btn-secondary">
          {{ showWithPhones ? 'Simple View' : 'Detailed View' }}
        </button>
        <button @click="goBack" class="btn-back">
          ← Back to Add User
        </button>
      </div>
    </div>
    
    <div class="loading" v-if="loading">
      <p>Loading users...</p>
    </div>
    
    <div v-else class="user-list-container">
      <h3 v-if="showWithPhones">Users with their phone numbers:</h3>
      <h3 v-else>Select a user to view details:</h3>
      
      <UserWithPhonesList 
        v-if="showWithPhones" 
        :usersWithPhones="usersWithPhones" 
        @deleted="loadUsersWithPhones"
        @phoneUpdated="loadUsersWithPhones"
        @phoneDeleted="loadUsersWithPhones"
      />
      <UserList 
        v-else 
        :users="users" 
        @deleted="loadUsers" 
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import UserList from '../components/UserDetails.vue'
import UserWithPhonesList from '../components/UserWithPhonesList.vue'
import { getUsers, getUsersWithPhones } from '../services/userService'

const router = useRouter()
const users = ref([])
const usersWithPhones = ref([])
const loading = ref(false)
const showWithPhones = ref(true)

const loadUsers = async () => {
  loading.value = true
  try {
    const res = await getUsers()
    users.value = res.data
  } catch (error) {
    console.error('Error loading users:', error)
  } finally {
    loading.value = false
  }
}

const loadUsersWithPhones = async () => {
  loading.value = true
  try {
    const res = await getUsersWithPhones()
    usersWithPhones.value = res.data
  } catch (error) {
    console.error('Error loading users with phones:', error)
  } finally {
    loading.value = false
  }
}

const toggleView = () => {
  showWithPhones.value = !showWithPhones.value
  if (showWithPhones.value && usersWithPhones.value.length === 0) {
    loadUsersWithPhones()
  } else if (!showWithPhones.value && users.value.length === 0) {
    loadUsers()
  }
}

const goBack = () => {
  router.push('/users')
}

onMounted(() => {
  loadUsersWithPhones()
})
</script>

<style scoped>
.users-with-phones-view {
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

.header-actions {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.user-list-container {
  margin-top: 1rem;
}

.loading {
  text-align: center;
  padding: 2rem;
  color: #666;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  text-decoration: none;
  font-size: 0.9rem;
  transition: background-color 0.2s;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background-color: #545b62;
}

.btn-back {
  background-color: #007bff;
  color: white;
}

.btn-back:hover {
  background-color: #0056b3;
}
</style>
