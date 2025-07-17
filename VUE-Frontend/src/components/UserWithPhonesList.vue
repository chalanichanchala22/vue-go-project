<template>
  <div class="user-with-phones-list">
    <div v-if="usersWithPhones.length === 0" class="no-users">
      <p>No users found. Add a user to get started!</p>
    </div>
    <div v-else class="users-grid">
      <div v-for="userWithPhones in usersWithPhones" :key="userWithPhones.user.id" class="user-card">
        <div class="user-info">
          <div class="user-details">
            <h3 class="user-name">{{ userWithPhones.user.name }}</h3>
            <p class="user-email">{{ userWithPhones.user.email }}</p>
          </div>
        </div>
        
        <div class="phone-numbers">
          <div class="phone-header">
            <h4>Phone Numbers:</h4>
          </div>

          <div v-if="userWithPhones.phones && userWithPhones.phones.length > 0" class="phones-list">
            <div v-for="phone in userWithPhones.phones" :key="phone.id" class="phone-item">
              <div class="phone-info">
                <span class="phone-number">{{ phone.number }}</span>
                <span class="phone-type">{{ phone.type }}</span>
              </div>
              <div class="phone-actions">
                <button @click="editPhone(userWithPhones.user.id, phone)" class="btn-phone btn-edit">
                  ✏️
                </button>
                <button @click="deletePhone(userWithPhones.user.id, phone.id)" class="btn-phone btn-delete">
                  🗑️
                </button>
              </div>
            </div>
          </div>
          <div v-else class="no-phones">
            <p>No phone numbers added</p>
          </div>
        </div>
        
        <div class="user-actions">
          <router-link :to="`/users/${userWithPhones.user.id}`" class="btn btn-primary">
            View Details
          </router-link>
          <router-link :to="`/users/${userWithPhones.user.id}/edit`" class="btn btn-success">
            Edit
          </router-link>
          <button @click="del(userWithPhones.user.id)" class="btn btn-danger">
            Delete
          </button>
        </div>
      </div>
    </div>

    <!-- Edit Phone Modal -->
    <div v-if="editingPhone" class="modal-overlay" @click="editingPhone = null">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h4>Edit Phone Number</h4>
          <button @click="editingPhone = null" class="close-btn">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label>Phone Number:</label>
            <input
              v-model="editingPhone.number"
              type="tel"
              class="form-control"
              required
            />
          </div>
          <div class="form-group">
            <label>Type:</label>
            <select v-model="editingPhone.type" class="form-control" required>
              <option value="mobile">Mobile</option>
              <option value="home">Home</option>
              <option value="work">Work</option>
              <option value="other">Other</option>
            </select>
          </div>
        </div>
        <div class="modal-footer">
          <button @click="updatePhone" class="btn btn-success">
            Update Phone
          </button>
          <button @click="editingPhone = null" class="btn btn-secondary">
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { deleteUser } from '../services/userService'
import { updatePhone as updatePhoneService, deletePhone as deletePhoneService } from '../services/phoneService'

const props = defineProps(['usersWithPhones'])
const emit = defineEmits(['deleted', 'edit', 'phoneEdit', 'phoneDeleted', 'userProfileEdit', 'phoneUpdated'])

// Phone management state
const editingPhone = ref(null)

const edit = (user) => {
  emit('edit', user)
}

const editUserProfile = (user) => {
  emit('userProfileEdit', user)
}

const editPhone = (userId, phone) => {
  editingPhone.value = { ...phone, userId }
}

const updatePhone = async () => {
  try {
    await updatePhoneService(editingPhone.value.userId, editingPhone.value.id, editingPhone.value)
    emit('phoneUpdated')
    editingPhone.value = null
  } catch (error) {
    console.error('Error updating phone:', error)
  }
}

const deletePhone = async (userId, phoneId) => {
  if (confirm('Are you sure you want to delete this phone number?')) {
    try {
      await deletePhoneService(userId, phoneId)
      emit('phoneDeleted')
    } catch (error) {
      console.error('Error deleting phone:', error)
    }
  }
}

const del = async (id) => {
  if (confirm('Are you sure you want to delete this user?')) {
    await deleteUser(id)
    emit('deleted')
  }
}
</script>

<style scoped>
.user-with-phones-list {
  margin: 2rem 0;
}

.no-users {
  text-align: center;
  padding: 2rem;
  color: #666;
  font-style: italic;
}

.users-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 1.5rem;
}

.user-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 1.5rem;
  background-color: #f9f9f9;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
}

.user-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.user-info {
  margin-bottom: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.user-details {
  flex: 1;
}

.user-edit-actions {
  margin-left: 0.5rem;
}

.btn-profile {
  padding: 0.3rem;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.2s;
  min-width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-edit-profile {
  background-color: #17a2b8;
  color: white;
}

.btn-edit-profile:hover {
  background-color: #138496;
  transform: scale(1.1);
}

.user-name {
  margin: 0 0 0.5rem 0;
  color: #333;
  font-size: 1.2rem;
}

.user-email {
  margin: 0;
  color: #666;
  font-size: 0.9rem;
}

.phone-numbers {
  margin-bottom: 1.5rem;
}

.phone-numbers h4 {
  margin: 0 0 0.8rem 0;
  color: #555;
  font-size: 1rem;
}

.phone-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.8rem;
}

.phone-header h4 {
  margin: 0;
}

.phones-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.phone-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem;
  background-color: #fff;
  border-radius: 4px;
  border: 1px solid #e0e0e0;
}

.phone-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex: 1;
}

.phone-actions {
  display: flex;
  gap: 0.3rem;
  margin-left: 0.5rem;
}

.btn-phone {
  padding: 0.3rem;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  font-size: 0.8rem;
  transition: all 0.2s;
  min-width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-edit {
  background-color: #28a745;
  color: white;
}

.btn-edit:hover {
  background-color: #218838;
  transform: scale(1.1);
}

.btn-delete {
  background-color: #dc3545;
  color: white;
}

.btn-delete:hover {
  background-color: #c82333;
  transform: scale(1.1);
}

.phone-number {
  font-weight: 500;
  color: #333;
}

.phone-type {
  font-size: 0.8rem;
  color: #666;
  text-transform: uppercase;
  padding: 0.2rem 0.5rem;
  background-color: #f0f0f0;
  border-radius: 12px;
}

.no-phones {
  text-align: center;
  padding: 1rem;
  color: #999;
  font-style: italic;
  background-color: #f5f5f5;
  border-radius: 4px;
}

.user-actions {
  display: flex;
  gap: 0.5rem;
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

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-success {
  background-color: #28a745;
  color: white;
}

.btn-success:hover {
  background-color: #218838;
}

.btn-danger {
  background-color: #dc3545;
  color: white;
}

.btn-danger:hover {
  background-color: #c82333;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid #e9ecef;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h4 {
  margin: 0;
  color: #333;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #666;
  line-height: 1;
}

.close-btn:hover {
  color: #333;
}

.modal-body {
  padding: 1.5rem;
}

.modal-footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid #e9ecef;
  display: flex;
  gap: 0.5rem;
  justify-content: flex-end;
}
</style>
