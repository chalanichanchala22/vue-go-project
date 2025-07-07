<template>
  <ul>
    <li v-for="user in users" :key="user.id">
      <router-link :to="`/users/${user.id}`">{{ user.name }}</router-link>
      <button @click="del(user.id)">Delete</button>
    </li>
  </ul>
</template>

<script setup>
import { deleteUser } from '../services/userService'
const props = defineProps(['users'])
const emit = defineEmits(['deleted'])

const del = async (id) => {
  await deleteUser(id)
  emit('deleted')
}
</script>
