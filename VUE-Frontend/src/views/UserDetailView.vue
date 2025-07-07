<template>
  <div class="user-detail-container">
    <!-- Header Section -->
    <div class="header-section">
      <div class="back-button">
        <button @click="$router.back()" class="btn btn-secondary">
          <i class="icon-arrow-left"></i> Back
        </button>
      </div>
      <h1 class="page-title">User Details</h1>
    </div>

    <!-- User Info Card -->
    <div class="user-card">
      <div class="user-header">
        <div class="user-avatar">
          <div class="avatar-placeholder">
            <i class="icon-user"></i>
          </div>
        </div>
        <div class="user-title">
          <h2 class="user-name">{{ user.name || `User #${userId}` }}</h2>
          <p class="user-subtitle">User Profile</p>
        </div>
      </div>
      
      <div class="user-details">
        <div class="detail-column">
          <div class="user-detail-item">
            <span class="label">ID</span>
            <span class="value">{{ userId }}</span>
          </div>
          <div v-if="user.email" class="user-detail-item">
            <span class="label">Email</span>
            <span class="value">{{ user.email }}</span>
          </div>
          <div v-if="user.nic" class="user-detail-item">
            <span class="label">NIC</span>
            <span class="value">{{ user.nic }}</span>
          </div>
        </div>
        
        <div class="detail-column">
          <div v-if="user.address" class="user-detail-item">
            <span class="label">Address</span>
            <span class="value">{{ user.address }}</span>
          </div>
          <div v-if="user.gender" class="user-detail-item">
            <span class="label">Gender</span>
            <span class="value">{{ user.gender }}</span>
          </div>
          <div v-if="user.birthday" class="user-detail-item">
            <span class="label">Birthday</span>
            <span class="value">{{ formatDate(user.birthday) }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Phone Numbers Section -->
    <div class="phones-section">
      <div class="section-header">
        <h3 class="section-title">Phone Numbers</h3>
        <span class="phone-count">{{ phones.length }} phone(s)</span>
      </div>
      
      <div class="phones-container">
        <PhoneList :phones="phones" :userId="userId" @deleted="loadPhones" />
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="loading-state">
      <div class="spinner"></div>
      <p>Loading phone numbers...</p>
    </div>

    <!-- Empty State -->
    <div v-if="!loading && phones.length === 0" class="empty-state">
      <div class="empty-icon">
        <i class="icon-phone-off"></i>
      </div>
      <h3>No phone numbers found</h3>
      <p>This user doesn't have any phone numbers yet.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import PhoneList from '../components/PhoneList.vue'
import { getPhonesByUser } from '../services/phoneService'
import { getUser } from '../services/userService'

const route = useRoute()
const userId = route.params.id
const phones = ref([])
const user = ref({})
const loading = ref(true)

const loadPhones = async () => {
  try {
    loading.value = true
    const res = await getPhonesByUser(userId)
    phones.value = res.data || []
  } catch (error) {
    console.error('Error loading phones:', error)
    phones.value = []
  } finally {
    loading.value = false
  }
}

const loadUser = async () => {
  try {
    const res = await getUser(userId)
    user.value = res.data || {}
  } catch (error) {
    console.error('Error loading user:', error)
    user.value = {}
  }
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

onMounted(() => {
  loadUser()
  loadPhones()
})
</script>

<style scoped>
.user-detail-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

/* Header Section */
.header-section {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
  gap: 20px;
}

.back-button .btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  border: none;
  border-radius: 8px;
  background: #f8f9fa;
  color: #495057;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.2s ease;
}

.back-button .btn:hover {
  background: #e9ecef;
  transform: translateY(-1px);
}

.page-title {
  font-size: 28px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

/* User Card */
.user-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 16px;
  padding: 30px;
  margin-bottom: 30px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  color: white;
}

.user-header {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
}

.user-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  backdrop-filter: blur(10px);
}

.user-title h2 {
  margin: 0 0 4px 0;
  font-size: 28px;
  font-weight: 600;
}

.user-subtitle {
  margin: 0;
  opacity: 0.7;
  font-size: 14px;
  font-weight: 400;
}

.user-details {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 30px;
}

.detail-column {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.user-detail-item {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.user-detail-item .label {
  font-size: 12px;
  font-weight: 600;
  opacity: 0.7;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.user-detail-item .value {
  font-size: 16px;
  font-weight: 500;
  opacity: 0.95;
  word-break: break-word;
}

/* Phones Section */
.phones-section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 16px rgba(0, 0, 0, 0.05);
  border: 1px solid #e9ecef;
}

.section-header {
  display: flex;
  justify-content: between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 16px;
  border-bottom: 2px solid #f8f9fa;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.phone-count {
  background: #e3f2fd;
  color: #1976d2;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
}

.phones-container {
  margin-top: 20px;
}

/* Loading State */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px;
  color: #6c757d;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 3px solid #f3f3f3;
  border-top: 3px solid #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: #6c757d;
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 20px;
  opacity: 0.5;
}

.empty-state h3 {
  font-size: 20px;
  margin-bottom: 8px;
  color: #495057;
}

.empty-state p {
  font-size: 14px;
  margin: 0;
}

/* Icons (you can replace these with actual icon fonts like Font Awesome) */
.icon-arrow-left::before { content: '←'; }
.icon-user::before { content: '👤'; }
.icon-phone-off::before { content: '📵'; }

/* Responsive Design */
@media (max-width: 768px) {
  .user-detail-container {
    padding: 16px;
  }
  
  .header-section {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
  
  .user-card {
    padding: 24px;
  }
  
  .user-header {
    flex-direction: column;
    text-align: center;
    gap: 16px;
    margin-bottom: 24px;
  }
  
  .user-details {
    grid-template-columns: 1fr;
    gap: 20px;
  }
  
  .section-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
}
</style>
