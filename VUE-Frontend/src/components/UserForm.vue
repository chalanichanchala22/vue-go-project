<template>
  <div class="user-form">
    <h2>Add New User</h2>
    <form @submit.prevent="submit" class="form">
      <div class="form-group">
        <label for="name">Name *</label>
        <input 
          id="name"
          v-model="user.name" 
          type="text"
          placeholder="Enter full name" 
          required 
        />
      </div>

      <div class="form-group">
        <label for="email">Email *</label>
        <input 
          id="email"
          v-model="user.email" 
          type="email"
          placeholder="Enter email address" 
          required 
          pattern="[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$"
        />
      </div>

      <div class="form-group">
        <label for="nic">NIC *</label>
        <input 
          id="nic"
          v-model="user.nic" 
          type="text"
          placeholder="Enter NIC number" 
          required 
          maxlength="12"
        />
      </div>

      <div class="form-group">
        <label for="address">Address *</label>
        <textarea 
          id="address"
          v-model="user.address" 
          placeholder="Enter address"
          rows="3"
          required
        ></textarea>
      </div>

      <div class="form-row">
        <div class="form-group">
          <label for="birthday">Birthday *</label>
          <input 
            id="birthday"
            v-model="user.birthday" 
            type="date" 
            required 
            :max="getMaxDate()"
          />
        </div>

        <div class="form-group">
          <label for="gender">Gender *</label>
          <select id="gender" v-model="user.gender" required>
            <option value="">Select Gender</option>
            <option value="Male">Male</option>
            <option value="Female">Female</option>
          </select>
        </div>
      </div>

      <!-- Phone Numbers Section -->
      <div class="phone-section">
        <h3>Phone Numbers</h3>
        <div v-for="(phone, index) in phoneNumbers" :key="index" class="phone-entry">
          <div class="form-row">
            <div class="form-group">
              <label :for="`phone-${index}`">Phone Number</label>
              <input 
                :id="`phone-${index}`"
                v-model="phone.number" 
                type="tel"
                placeholder="Enter phone number"
                pattern="[0-9+\-\s()]*"
                maxlength="10"
              />
            </div>
            <div class="form-group">
              <label :for="`type-${index}`">Type</label>
              <select :id="`type-${index}`" v-model="phone.type">
                <option value="mobile">Mobile</option>
                <option value="home">Home</option>
                <option value="work">Work</option>
                <option value="other">Other</option>
              </select>
            </div>
            <div class="form-group">
              <label>&nbsp;</label>
              <button 
                type="button" 
                @click="removePhone(index)"
                class="btn-remove"
                :disabled="phoneNumbers.length <= 1"
              >
                Remove
              </button>
            </div>
          </div>
        </div>
        <button type="button" @click="addPhone" class="btn-add">
          Add Phone Number
        </button>
      </div>

      <div class="form-actions">
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
        <div class="button-group">
          <button type="submit" class="btn-primary" :disabled="isLoading">
            <span v-if="isLoading">Adding User...</span>
            <span v-else>Add User</span>
          </button>
          <button type="button" class="btn-secondary" @click="viewUserDetails">
            View User Details
          </button>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { createUser } from '../services/userService'
import { createPhone } from '../services/phoneService'

const emit = defineEmits(['refresh', 'viewUserDetails'])
const router = useRouter()

const user = ref({
  name: '', email: '', nic: '', address: '', birthday: '', gender: ''
})

const phoneNumbers = ref([
  { number: '', type: 'mobile' }
])

const isLoading = ref(false)
const errorMessage = ref('')

const addPhone = () => {
  phoneNumbers.value.push({ number: '', type: 'mobile' })
}

const removePhone = (index) => {
  if (phoneNumbers.value.length > 1) {
    phoneNumbers.value.splice(index, 1)
  }
}

const getMaxDate = () => {
  const today = new Date()
  return today.toISOString().split('T')[0]
}

const viewUserDetails = () => {
  // Navigate to user list view
  router.push('/userlist')
}

const submit = async () => {
  try {
    isLoading.value = true
    errorMessage.value = ''
    
    // Validate required fields
    if (!user.value.name || !user.value.email || !user.value.nic || 
        !user.value.address || !user.value.birthday || !user.value.gender) {
      errorMessage.value = 'Please fill in all required fields'
      return
    }
    
    // Validate phone numbers
    const validPhones = phoneNumbers.value.filter(phone => phone.number.trim() !== '')
    for (const phone of validPhones) {
      if (phone.number.trim().length < 10) {
        errorMessage.value = 'Phone numbers must be at least 10 digits long'
        return
      }
    }
    
    // Validate birthday is not in the future
    const today = new Date()
    const selectedDate = new Date(user.value.birthday)
    if (selectedDate > today) {
      errorMessage.value = 'Birthday cannot be in the future'
      return
    }
    
    // Format the user data to match backend expectations
    const birthdayDate = new Date(user.value.birthday)
    birthdayDate.setUTCHours(0, 0, 0, 0) // Set to midnight UTC
    
    const userData = {
      name: user.value.name.trim(),
      email: user.value.email.trim(),
      nic: user.value.nic.trim(),
      address: user.value.address.trim(),
      birthday: birthdayDate.toISOString(), // Convert to RFC3339 format
      gender: user.value.gender
    }
    
    console.log('Submitting user data:', userData)
    
    const response = await createUser(userData)
    console.log('User created successfully:', response.data)
    
    // Create phone numbers if any are provided
    if (validPhones.length > 0) {
      for (const phone of validPhones) {
        const phoneData = {
          user_id: response.data.id,
          number: phone.number.trim(),
          type: phone.type
        }
        try {
          await createPhone(phoneData)
          console.log('Phone number created:', phoneData)
        } catch (phoneError) {
          console.error('Error creating phone number:', phoneError)
        }
      }
    }
    
    // Reset form
    user.value = { name: '', email: '', nic: '', address: '', birthday: '', gender: '' }
    phoneNumbers.value = [{ number: '', type: 'mobile' }]
    
    // Emit refresh event
    emit('refresh')
    
  } catch (error) {
    console.error('Error creating user:', error)
    console.error('Error response:', error.response?.data)
    
    // More detailed error handling
    if (error.response?.data?.error) {
      errorMessage.value = error.response.data.error
    } else if (error.response?.status === 400) {
      errorMessage.value = 'Invalid data provided. Please check all fields.'
    } else {
      errorMessage.value = 'Failed to create user. Please try again.'
    }
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.user-form {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  background: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.user-form h2 {
  text-align: center;
  color: #333;
  margin-bottom: 2rem;
  font-size: 1.5rem;
  font-weight: 600;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.form-group label {
  font-weight: 500;
  color: #555;
  font-size: 0.9rem;
}

.form-group input,
.form-group select,
.form-group textarea {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #4CAF50;
  box-shadow: 0 0 0 2px rgba(76, 175, 80, 0.2);
}

.form-group textarea {
  resize: vertical;
  min-height: 80px;
}

.form-actions {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  margin-top: 1rem;
}

.error-message {
  color: #dc3545;
  background-color: #f8d7da;
  border: 1px solid #f5c6cb;
  padding: 0.75rem;
  border-radius: 4px;
  font-size: 0.9rem;
  width: 100%;
  text-align: center;
}

.btn-primary {
  background-color: #4CAF50;
  color: white;
  padding: 0.75rem 2rem;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s ease;
  min-width: 120px;
}

.btn-primary:hover:not(:disabled) {
  background-color: #45a049;
}

.btn-primary:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.btn-primary:active:not(:disabled) {
  transform: translateY(1px);
}

.button-group {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
  justify-content: center;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
  padding: 0.75rem 2rem;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s ease;
  min-width: 120px;
}

.btn-secondary:hover {
  background-color: #5a6268;
}

.btn-secondary:active {
  transform: translateY(1px);
}

/* Phone Numbers Section */
.phone-section {
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 1.5rem;
  background-color: #f9f9f9;
}

.phone-section h3 {
  margin: 0 0 1rem 0;
  color: #333;
  font-size: 1.1rem;
  font-weight: 600;
}

.phone-entry {
  margin-bottom: 1rem;
  padding: 1rem;
  background-color: white;
  border: 1px solid #e0e0e0;
  border-radius: 4px;
}

.phone-entry:last-child {
  margin-bottom: 0;
}

.btn-add, .btn-remove {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-add {
  background-color: #28a745;
  color: white;
  margin-top: 1rem;
}

.btn-add:hover {
  background-color: #218838;
}

.btn-remove {
  background-color: #dc3545;
  color: white;
  height: fit-content;
  margin-top: auto;
}

.btn-remove:hover:not(:disabled) {
  background-color: #c82333;
}

.btn-remove:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

/* Responsive design */
@media (max-width: 768px) {
  .user-form {
    padding: 1.5rem;
    margin: 1rem;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .button-group {
    flex-direction: column;
    align-items: center;
  }
  
  .btn-primary, .btn-secondary {
    width: 100%;
    max-width: 300px;
  }
}
</style>
