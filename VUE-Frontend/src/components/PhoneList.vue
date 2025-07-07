<template>
  <ul>
    <li v-for="phone in phones" :key="phone.id">
      {{ phone.number }} 
      <button @click="del(phone.id)">Delete</button>
    </li>
  </ul>
</template>

<script setup>
import { deletePhone } from '../services/phoneService'
const props = defineProps(['phones', 'userId'])
const emit = defineEmits(['deleted'])

const del = async (phoneId) => {
  await deletePhone(props.userId, phoneId)
  emit('deleted')
}
</script>
