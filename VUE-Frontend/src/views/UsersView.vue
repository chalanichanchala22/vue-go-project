<template>
  <div class="users-view">
    <UserForm @refresh="loadUsers" />
    <UserList :users="users" @deleted="loadUsers" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import UserForm from '../components/UserForm.vue'
import { getUsers } from '../services/userService'

const users = ref([])

const loadUsers = async () => {
  const res = await getUsers()
  users.value = res.data
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

@media (max-width: 768px) {
  .users-view {
    padding: 0.5rem;
  }
}
</style>
