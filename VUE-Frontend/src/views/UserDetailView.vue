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
        <div class="header-actions">
          <span class="phone-count">{{ phones.length }} phone(s)</span>
          <button @click="showAddForm = true" class="btn btn-primary add-phone-btn">
            <i class="icon-plus"></i> Add Phone
          </button>
        </div>
      </div>
      
      <!-- Add Phone Form -->
      <div v-if="showAddForm" class="add-phone-form">
        <h4>Add New Phone Number</h4>
        <div class="form-group">
          <label for="phoneNumber">Phone Number:</label>
          <input
            id="phoneNumber"
            v-model="newPhone.number"
            type="tel"
            class="form-control"
            placeholder="Enter phone number"
            required
          />
        </div>
        <div class="form-group">
          <label for="phoneType">Type:</label>
          <select id="phoneType" v-model="newPhone.type" class="form-control" required>
            <option value="">Select type</option>
            <option value="mobile">Mobile</option>
            <option value="home">Home</option>
            <option value="work">Work</option>
            <option value="other">Other</option>
          </select>
        </div>
        <div class="form-actions">
          <button @click="addPhone" class="btn btn-success" :disabled="!newPhone.number || !newPhone.type">
            Add Phone
          </button>
          <button @click="cancelAdd" class="btn btn-secondary">
            Cancel
          </button>
        </div>
      </div>
      
      <div class="phones-container">
        <!-- Phone List -->
        <div v-if="phones.length > 0" class="phones-list">
          <div v-for="phone in phones" :key="phone.id" class="phone-item">
            <div class="phone-info">
              <div class="phone-number">{{ phone.number }}</div>
              <div class="phone-type">{{ phone.type }}</div>
            </div>
            <div class="phone-actions">
              <button @click="editPhone(phone)" class="btn btn-sm btn-secondary">
                <i class="icon-edit"></i> Edit
              </button>
              <button @click="deletePhone(phone.id)" class="btn btn-sm btn-danger">
                <i class="icon-delete"></i> Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Edit Phone Modal -->
    <div v-if="editingPhone" class="modal-overlay" @click="cancelEdit">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h4>Edit Phone Number</h4>
          <button @click="cancelEdit" class="close-btn">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label for="editPhoneNumber">Phone Number:</label>
            <input
              id="editPhoneNumber"
              v-model="editingPhone.number"
              type="tel"
              class="form-control"
              required
            />
          </div>
          <div class="form-group">
            <label for="editPhoneType">Type:</label>
            <select id="editPhoneType" v-model="editingPhone.type" class="form-control" required>
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
          <button @click="cancelEdit" class="btn btn-secondary">
            Cancel
          </button>
        </div>
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
import { ref, reactive, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getPhonesByUser, createPhone, updatePhone as updatePhoneService, deletePhone as deletePhoneService } from '../services/phoneService'
import { getUser } from '../services/userService'

const route = useRoute()
const userId = route.params.id
const phones = ref([])
const user = ref({})
const loading = ref(true)

// Phone management state
const showAddForm = ref(false)
const editingPhone = ref(null)
const newPhone = reactive({
  number: '',
  type: ''
})

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

// Phone management functions
const addPhone = async () => {
  try {
    await createPhone({
      ...newPhone,
      user_id: userId
    })
    await loadPhones()
    cancelAdd()
  } catch (error) {
    console.error('Error adding phone:', error)
  }
}

const editPhone = (phone) => {
  editingPhone.value = { ...phone }
}

const updatePhone = async () => {
  try {
    await updatePhoneService(userId, editingPhone.value.id, editingPhone.value)
    await loadPhones()
    cancelEdit()
  } catch (error) {
    console.error('Error updating phone:', error)
  }
}

const deletePhone = async (phoneId) => {
  if (confirm('Are you sure you want to delete this phone number?')) {
    try {
      await deletePhoneService(userId, phoneId)
      await loadPhones()
    } catch (error) {
      console.error('Error deleting phone:', error)
    }
  }
}

const cancelAdd = () => {
  showAddForm.value = false
  newPhone.number = ''
  newPhone.type = ''
}

const cancelEdit = () => {
  editingPhone.value = null
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
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 16px;
  border-bottom: 2px solid #f8f9fa;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.add-phone-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  font-size: 0.9rem;
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

/* Add Phone Form */
.add-phone-form {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 8px;
  margin-bottom: 1.5rem;
  border: 1px solid #e9ecef;
}

.add-phone-form h4 {
  margin: 0 0 1rem 0;
  color: #333;
  font-size: 1.1rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: #333;
}

.form-control {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 0.9rem;
  transition: border-color 0.2s;
}

.form-control:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.form-actions {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}

/* Phone List */
.phones-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.phone-item {
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: box-shadow 0.2s;
}

.phone-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.phone-info {
  flex: 1;
}

.phone-number {
  font-size: 1.1rem;
  font-weight: 600;
  color: #333;
  margin-bottom: 0.25rem;
}

.phone-type {
  font-size: 0.9rem;
  color: #666;
  text-transform: capitalize;
  background: #f0f0f0;
  padding: 0.2rem 0.5rem;
  border-radius: 12px;
  display: inline-block;
}

.phone-actions {
  display: flex;
  gap: 0.5rem;
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
.icon-plus::before { content: '+'; }
.icon-edit::before { content: '✏️'; }
.icon-delete::before { content: '🗑️'; }

/* Button Styles */
.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  gap: 0.25rem;
  transition: all 0.2s;
  font-size: 0.9rem;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-primary {
  background: #007bff;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #0056b3;
}

.btn-success {
  background: #28a745;
  color: white;
}

.btn-success:hover:not(:disabled) {
  background: #1e7e34;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background: #545b62;
}

.btn-danger {
  background: #dc3545;
  color: white;
}

.btn-danger:hover {
  background: #c82333;
}

.btn-sm {
  padding: 0.375rem 0.75rem;
  font-size: 0.8rem;
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
