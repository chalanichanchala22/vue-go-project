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
            <div v-if="editingPhone === phone.id" class="phone-edit-form">
              <input 
                v-model="editForm.number" 
                type="text" 
                placeholder="Phone Number"
                class="phone-input"
                @keyup.enter="saveEdit(phone.id)"
                @keyup.escape="cancelEdit()"
                ref="phoneInput"
              />
              <select v-model="editForm.type" class="type-select">
                <option value="home">Home</option>
                <option value="work">Work</option>
                <option value="mobile">Mobile</option>
                <option value="other">Other</option>
              </select>
            </div>
            <div v-else class="phone-display">
              <span class="phone-number">{{ formatPhoneNumber(phone.number) }}</span>
              <span class="phone-type">{{ phone.type }}</span>
            </div>
          </div>
        </div>
        
        <div class="phone-actions">
          <div v-if="editingPhone === phone.id" class="edit-actions">
            <button 
              @click="saveEdit(phone.id)" 
              class="save-btn"
              :disabled="updating === phone.id"
            >
              <i v-if="updating === phone.id" class="icon-loading"></i>
              <i v-else class="icon-save"></i>
              {{ updating === phone.id ? 'Saving...' : 'Save' }}
            </button>
            <button 
              @click="cancelEdit()" 
              class="cancel-btn"
              :disabled="updating === phone.id"
            >
              <i class="icon-cancel"></i>
              Cancel
            </button>
          </div>
          <div v-else class="view-actions">
            <button 
              @click="startEdit(phone)" 
              class="edit-btn"
              :disabled="deleting === phone.id"
            >
              <i class="icon-edit"></i>
              Edit
            </button>
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
  </div>
</template>

<script setup>
import { ref, nextTick } from 'vue'
import { deletePhone, updatePhone } from '../services/phoneService'

const props = defineProps(['phones', 'userId'])
const emit = defineEmits(['deleted', 'edit', 'updated'])
const deleting = ref(null)
const updating = ref(null)
const editingPhone = ref(null)
const editForm = ref({
  number: '',
  type: ''
})

const edit = (phone) => {
  emit('edit', phone)
}

const startEdit = (phone) => {
  editingPhone.value = phone.id
  editForm.value = {
    number: phone.number,
    type: phone.type
  }
  
  // Focus the input after the DOM updates
  nextTick(() => {
    const input = document.querySelector('.phone-input')
    if (input) input.focus()
  })
}

const cancelEdit = () => {
  editingPhone.value = null
  editForm.value = {
    number: '',
    type: ''
  }
}

const saveEdit = async (phoneId) => {
  try {
    updating.value = phoneId
    console.log('Updating phone:', phoneId, 'with data:', editForm.value)
    const response = await updatePhone(props.userId, phoneId, editForm.value)
    console.log('Update response:', response)
    editingPhone.value = null
    editForm.value = {
      number: '',
      type: ''
    }
    emit('updated')
    console.log('Update completed, emitted updated event')
  } catch (error) {
    console.error('Error updating phone:', error)
    if (error.response) {
      console.error('Response data:', error.response.data)
      console.error('Response status:', error.response.status)
    }
  } finally {
    updating.value = null
  }
}

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
  flex: 1;
}

.phone-display {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.phone-edit-form {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.phone-input, .type-select {
  padding: 6px 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.phone-input {
  font-weight: 600;
}

.phone-input:focus, .type-select:focus {
  outline: none;
  border-color: #4CAF50;
  box-shadow: 0 0 0 2px rgba(76, 175, 80, 0.2);
}

.phone-number {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
}

.phone-type {
  font-size: 12px;
  color: #6c757d;
  text-transform: uppercase;
  background: #e9ecef;
  padding: 2px 6px;
  border-radius: 10px;
  align-self: flex-start;
}

.phone-id {
  font-size: 12px;
  color: #6c757d;
}

.phone-actions {
  display: flex;
  gap: 8px;
}

.edit-actions, .view-actions {
  display: flex;
  gap: 8px;
}

.save-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.save-btn:hover:not(:disabled) {
  background: #218838;
  transform: translateY(-1px);
}

.save-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.cancel-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: #6c757d;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.cancel-btn:hover:not(:disabled) {
  background: #5a6268;
  transform: translateY(-1px);
}

.cancel-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.edit-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.edit-btn:hover:not(:disabled) {
  background: #218838;
  transform: translateY(-1px);
}

.edit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.edit-btn:active {
  transform: translateY(0);
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
.icon-edit::before { content: '✏️'; }
.icon-save::before { content: '💾'; }
.icon-cancel::before { content: '❌'; }
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
