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
          <div class="user-edit-actions">
            <button @click="editUserProfile(userWithPhones.user)" class="btn-profile btn-edit-profile">
              ✏️
            </button>
          </div>
        </div>
        
        <div class="phone-numbers">
          <h4>Phone Numbers:</h4>
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
          <button @click="edit(userWithPhones.user)" class="btn btn-success">
            Edit
          </button>
          <button @click="del(userWithPhones.user.id)" class="btn btn-danger">
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { deleteUser } from '../services/userService'
import { deletePhone as deletePhoneService } from '../services/phoneService'
const props = defineProps(['usersWithPhones'])
const emit = defineEmits(['deleted', 'edit', 'phoneEdit', 'phoneDeleted', 'userProfileEdit'])

const edit = (user) => {
  emit('edit', user)
}

const editUserProfile = (user) => {
  emit('userProfileEdit', user)
}

const editPhone = (userId, phone) => {
  emit('phoneEdit', { userId, phone })
}

const deletePhone = async (userId, phoneId) => {
  try {
    await deletePhoneService(userId, phoneId)
    emit('phoneDeleted')
  } catch (error) {
    console.error('Error deleting phone:', error)
  }
}

const del = async (id) => {
  await deleteUser(id)
  emit('deleted')
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
</style>
