<template>
  <UserForm @refresh="loadUsers" />
  <UserList :users="users" @deleted="loadUsers" />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import UserForm from '../components/UserForm.vue'
import UserList from '../components/UserList.vue'
import { getUsers } from '../services/userService'

const users = ref([])

const loadUsers = async () => {
  const res = await getUsers()
  users.value = res.data
}

onMounted(loadUsers)
</script>
