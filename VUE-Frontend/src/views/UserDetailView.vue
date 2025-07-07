<template>
  <h1>User Details</h1>
  <PhoneList :phones="phones" :userId="userId" @deleted="loadPhones" />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import PhoneList from '../components/PhoneList.vue'
import { getPhonesByUser } from '../services/phoneService'

const route = useRoute()
const userId = route.params.id
const phones = ref([])

const loadPhones = async () => {
  const res = await getPhonesByUser(userId)
  phones.value = res.data
}

onMounted(loadPhones)
</script>
