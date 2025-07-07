<template>
  <div class="user-list">
    <div v-if="users.length === 0" class="no-users">
      <p>No users found. Add a user to get started!</p>
    </div>
    <div v-else class="users-grid">
      <div v-for="user in users" :key="user.id" class="user-card">
        <div class="user-info">
          <h3 class="user-name">{{ user.name }}</h3>
          <p class="user-email">{{ user.email }}</p>
        </div>
        <div class="user-actions">
          <router-link :to="`/users/${user.id}`" class="btn btn-primary">
            View Details
          </router-link>
          <button @click="del(user.id)" class="btn btn-danger">
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
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

<style scoped>
.user-list {
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
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
  margin-top: 1rem;
}

.user-card {
  background: #ffffff;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 1.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.user-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.user-info {
  flex: 1;
}

.user-name {
  margin: 0 0 0.5rem 0;
  color: #333;
  font-size: 1.1rem;
  font-weight: 600;
}

.user-email {
  margin: 0;
  color: #666;
  font-size: 0.9rem;
}

.user-actions {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  text-decoration: none;
  text-align: center;
  transition: background-color 0.3s ease;
  flex: 1;
  min-width: 100px;
}

.btn-primary {
  background-color: #4CAF50;
  color: white;
}

.btn-primary:hover {
  background-color: #45a049;
}

.btn-danger {
  background-color: #dc3545;
  color: white;
}

.btn-danger:hover {
  background-color: #c82333;
}

.btn:active {
  transform: translateY(1px);
}

@media (max-width: 768px) {
  .users-grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
  
  .user-card {
    padding: 1rem;
  }
  
  .user-actions {
    flex-direction: column;
  }
  
  .btn {
    flex: none;
    width: 100%;
  }
}
</style>
