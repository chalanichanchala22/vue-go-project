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

       <!-- Profile Photo Section -->
      <div class="form-group photo-upload-section">
        <label for="photo">Profile Photo</label>
        <div class="photo-upload-container">
          <div v-if="photoPreview" class="photo-preview">
            <img :src="photoPreview" alt="Profile Preview" class="preview-image" />
            <button type="button" @click="removePhoto" class="remove-photo-btn">
              <span>×</span>
            </button>
          </div>
          <div v-else class="upload-placeholder">
            <div class="upload-icon">📷</div>
            <p>Click to upload photo</p>
            <p class="upload-hint">JPG, PNG, GIF up to 5MB</p>
          </div>
          <input 
            id="photo"
            ref="fileInput"
            type="file"
            accept="image/*"
            @change="handlePhotoUpload"
            class="photo-input"
          />
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
                pattern="[0-9+\s()-]*"
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
import { createUser, createUserWithPhoto } from '../services/userService'
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
const selectedPhoto = ref(null)
const photoPreview = ref(null)
const fileInput = ref(null)

const addPhone = () => {
  phoneNumbers.value.push({ number: '', type: 'mobile' })
}

const removePhone = (index) => {
  if (phoneNumbers.value.length > 1) {
    phoneNumbers.value.splice(index, 1)
  }
}

const handlePhotoUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    // Validate file size (5MB limit)
    if (file.size > 5 * 1024 * 1024) {
      errorMessage.value = 'Photo size must be less than 5MB'
      return
    }
    
    // Validate file type
    if (!file.type.startsWith('image/')) {
      errorMessage.value = 'Please select a valid image file'
      return
    }
    
    selectedPhoto.value = file
    
    // Create preview
    const reader = new FileReader()
    reader.onload = (e) => {
      photoPreview.value = e.target.result
    }
    reader.readAsDataURL(file)
    
    errorMessage.value = ''
  }
}

const removePhoto = () => {
  selectedPhoto.value = null
  photoPreview.value = null
  if (fileInput.value) {
    fileInput.value.value = ''
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
    
    // Format the user data
    const birthdayDate = new Date(user.value.birthday)
    birthdayDate.setUTCHours(0, 0, 0, 0)
    
    let response
    
    // Always use FormData to ensure consistent behavior
    const formData = new FormData()
    formData.append('name', user.value.name.trim())
    formData.append('email', user.value.email.trim())
    formData.append('nic', user.value.nic.trim())
    formData.append('address', user.value.address.trim())
    formData.append('birthday', birthdayDate.toISOString())
    formData.append('gender', user.value.gender)
    
    if (selectedPhoto.value) {
      formData.append('photo', selectedPhoto.value)
      console.log('Submitting user data with photo')
      console.log('Photo file:', selectedPhoto.value)
    } else {
      console.log('Submitting user data without photo')
    }
    
    console.log('FormData entries:')
    for (let [key, value] of formData.entries()) {
      console.log(key, value)
    }
    
    response = await createUser(formData)
    
    console.log('User created successfully:', response.data)
    
    // Create phone numbers if any are provided
    if (validPhones.length > 0) {
      for (const phone of validPhones) {
        const phoneData = {
          number: phone.number.trim(),
          type: phone.type
        }
        try {
          await createPhone(response.data.id, phoneData)
          console.log('Phone number created:', phoneData)
        } catch (phoneError) {
          console.error('Error creating phone number:', phoneError)
        }
      }
    }
    
    // Reset form
    user.value = { name: '', email: '', nic: '', address: '', birthday: '', gender: '' }
    phoneNumbers.value = [{ number: '', type: 'mobile' }]
    removePhoto()
    
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

/* Photo Upload Section */
.photo-upload-section {
  margin-bottom: 1.5rem;
}

.photo-upload-container {
  position: relative;
  width: 150px;
  height: 150px;
  border: 2px dashed #ddd;
  border-radius: 8px;
  cursor: pointer;
  transition: border-color 0.3s ease;
  overflow: hidden;
}

.photo-upload-container:hover {
  border-color: #4CAF50;
}

.photo-input {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  cursor: pointer;
}

.upload-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  text-align: center;
  color: #666;
  padding: 1rem;
}

.upload-icon {
  font-size: 2rem;
  margin-bottom: 0.5rem;
}

.upload-placeholder p {
  margin: 0.25rem 0;
  font-size: 0.9rem;
}

.upload-hint {
  font-size: 0.8rem !important;
  color: #999;
}

.photo-preview {
  position: relative;
  width: 100%;
  height: 100%;
}

.preview-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 6px;
}

.remove-photo-btn {
  position: absolute;
  top: -8px;
  right: -8px;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;
  line-height: 1;
  transition: background-color 0.3s ease;
}

.remove-photo-btn:hover {
  background: #c82333;
}

.remove-photo-btn span {
  margin-top: -2px;
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
