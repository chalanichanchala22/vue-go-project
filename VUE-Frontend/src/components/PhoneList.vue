<template>
  <div class="phone-list">
    <div v-if="phones.length === 0" class="no-phones">
      <p>No phone numbers available</p>
    </div>
    
    <div v-else class="phone-items">
      <div 
        v-for="phone in phones" 
        :key="phone.id" 
        class="phone-item"
      >
        <div class="phone-info">
          <div class="phone-icon">
            <i class="icon-phone"></i>
          </div>
          <div class="phone-details">
            <span class="phone-number">{{ formatPhoneNumber(phone.number) }}</span>
           
          </div>
        </div>
        
        <div class="phone-actions">
          <button 
            @click="del(phone.id)" 
            class="delete-btn"
            :disabled="deleting === phone.id"
          >
            <i v-if="deleting === phone.id" class="icon-loading"></i>
            <i v-else class="icon-trash"></i>
            {{ deleting === phone.id ? 'Deleting...' : 'Delete' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { deletePhone } from '../services/phoneService'

const props = defineProps(['phones', 'userId'])
const emit = defineEmits(['deleted'])
const deleting = ref(null)

const del = async (phoneId) => {
  try {
    deleting.value = phoneId
    await deletePhone(props.userId, phoneId)
    emit('deleted')
  } catch (error) {
    console.error('Error deleting phone:', error)
  } finally {
    deleting.value = null
  }
}

const formatPhoneNumber = (number) => {
  // Simple phone number formatting - you can enhance this based on your needs
  if (!number) return ''
  const cleaned = number.replace(/\D/g, '')
  if (cleaned.length === 10) {
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }
  return number
}
</script>

<style scoped>
.phone-list {
  width: 100%;
}

.no-phones {
  text-align: center;
  padding: 20px;
  color: #6c757d;
  font-style: italic;
}

.phone-items {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.phone-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  transition: all 0.2s ease;
}

.phone-item:hover {
  background: #f1f3f4;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.phone-info {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.phone-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #4CAF50, #45a049);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 16px;
}

.phone-details {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.phone-number {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
}

.phone-id {
  font-size: 12px;
  color: #6c757d;
}

.phone-actions {
  display: flex;
  gap: 8px;
}

.delete-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.delete-btn:hover:not(:disabled) {
  background: #c82333;
  transform: translateY(-1px);
}

.delete-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.delete-btn:active {
  transform: translateY(0);
}

/* Icons */
.icon-phone::before { content: '📞'; }
.icon-trash::before { content: '🗑️'; }
.icon-loading::before { 
  content: '⟳'; 
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* Responsive Design */
@media (max-width: 576px) {
  .phone-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  
  .phone-info {
    width: 100%;
  }
  
  .phone-actions {
    width: 100%;
    justify-content: flex-end;
  }
  
  .delete-btn {
    width: auto;
  }
}
</style>
