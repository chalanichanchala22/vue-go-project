<template>
  <div class="users-view">
    <UserForm @refresh="loadUsers" @viewUserDetails="handleViewUserDetails" />
    <div v-if="showUserList" class="users-list-section">
      <h3>Select a user to view details:</h3>
      <UserList :users="users" @deleted="loadUsers" />
    </div>
    <UserList v-else :users="users" @deleted="loadUsers" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import UserForm from '../components/UserForm.vue'
import UserList from '../components/UserList.vue'
import { getUsers } from '../services/userService'

const users = ref([])
const showUserList = ref(false)

const loadUsers = async () => {
  const res = await getUsers()
  users.value = res.data
}

const handleViewUserDetails = () => {
  showUserList.value = true
  // Scroll to the user list section
  setTimeout(() => {
    const userListSection = document.querySelector('.users-list-section')
    if (userListSection) {
      userListSection.scrollIntoView({ behavior: 'smooth' })
    }
  }, 100)
}

onMounted(loadUsers)
</script>

<style scoped>
.users-view {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  padding: 1rem;
}

.users-list-section {
  background-color: #f8f9fa;
  border: 2px solid #4CAF50;
  border-radius: 8px;
  padding: 1.5rem;
  margin-top: 1rem;
}

.users-list-section h3 {
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
  .users-view {
    padding: 0.5rem;
  }
  
  .users-list-section {
    padding: 1rem;
  }
  
  .users-list-section h3 {
    margin: -1rem -1rem 1rem -1rem;
  }
}
</style>
