<template>
  <div class="user-edit-form">
    <h2>Edit User</h2>
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
            readonly
            disabled
          />
        </div>

        <div class="form-group">
          <label for="gender">Gender *</label>
          <select id="gender" v-model="user.gender" required>
            <option value="">Select Gender</option>
            <option value="Male">Male</option>
            <option value="Female">Female</option>
            <option value="Other">Other</option>
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
              <i class="icon-trash"></i>
            </button>
          </div>
          
          <div v-else class="photo-upload-area" @click="$refs.photoInput.click()">
            <i class="icon-camera"></i>
            <p>Click to upload photo</p>
            <small>JPG, PNG, GIF up to 5MB</small>
          </div>
          
          <input
            ref="photoInput"
            id="photo"
            type="file"
            accept="image/*"
            @change="handlePhotoUpload"
            style="display: none"
          />
          
          <div v-if="!photoPreview" class="photo-upload-buttons">
            <button type="button" @click="$refs.photoInput.click()" class="btn btn-outline">
              <i class="icon-upload"></i> Choose Photo
            </button>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="form-actions">
        <button type="submit" class="btn btn-primary" :disabled="loading">
          <span v-if="loading">Updating...</span>
          <span v-else>Update User</span>
        </button>
        <button type="button" @click="cancel" class="btn btn-secondary">
          Cancel
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue'
import { updateUser, updateUserWithPhoto } from '../services/userService'

const props = defineProps({
  userData: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['updated', 'cancelled'])

const loading = ref(false)
const photoFile = ref(null)
const photoPreview = ref('')

// Initialize user data
const user = reactive({
  name: '',
  email: '',
  nic: '',
  address: '',
  birthday: '',
  gender: ''
})

// Helper function to format date for input
const formatDateForInput = (dateString) => {
  if (!dateString) return ''
  try {
    const date = new Date(dateString)
    // Check if date is valid
    if (isNaN(date.getTime())) return ''
    return date.toISOString().split('T')[0]
  } catch (error) {
    console.error('Error formatting date:', error)
    return ''
  }
}

// Watch for changes in userData prop and update local user object
watch(() => props.userData, (newUserData) => {
  if (newUserData) {
    user.name = newUserData.name || ''
    user.email = newUserData.email || ''
    user.nic = newUserData.nic || ''
    user.address = newUserData.address || ''
    const formattedBirthday = formatDateForInput(newUserData.birthday)
    console.log('Formatting birthday:', newUserData.birthday, '->', formattedBirthday)
    user.birthday = formattedBirthday || ''
    user.gender = newUserData.gender || ''
    
    // Set existing photo preview
    if (newUserData.photo) {
      photoPreview.value = `http://localhost:8080${newUserData.photo}`
    } else {
      photoPreview.value = ''
    }
  }
}, { immediate: true })

// Also initialize on mount to ensure data is loaded
onMounted(() => {
  if (props.userData) {
    user.name = props.userData.name || ''
    user.email = props.userData.email || ''
    user.nic = props.userData.nic || ''
    user.address = props.userData.address || ''
    const formattedBirthday = formatDateForInput(props.userData.birthday)
    console.log('OnMounted - Formatting birthday:', props.userData.birthday, '->', formattedBirthday)
    user.birthday = formattedBirthday || ''
    user.gender = props.userData.gender || ''
    
    if (props.userData.photo) {
      photoPreview.value = `http://localhost:8080${props.userData.photo}`
    }
  }
})

// Watch birthday field specifically for debugging
watch(() => user.birthday, (newValue, oldValue) => {
  console.log('Birthday field changed from:', oldValue, 'to:', newValue)
})

const getMaxDate = () => {
  const today = new Date()
  return today.toISOString().split('T')[0]
}

const handlePhotoUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    // Validate file type
    if (!file.type.startsWith('image/')) {
      alert('Please select an image file')
      return
    }
    
    // Validate file size (5MB)
    if (file.size > 5 * 1024 * 1024) {
      alert('File size must be less than 5MB')
      return
    }
    
    photoFile.value = file
    
    // Create preview
    const reader = new FileReader()
    reader.onload = (e) => {
      photoPreview.value = e.target.result
    }
    reader.readAsDataURL(file)
  }
}

const removePhoto = () => {
  photoFile.value = null
  photoPreview.value = ''
  if (props.userData.photo) {
    // Mark for removal if there was an existing photo
    photoFile.value = 'REMOVE'
  }
}

const submit = async () => {
  try {
    loading.value = true
    
    console.log('Submitting user data:', {
      name: user.name,
      email: user.email,
      nic: user.nic,
      address: user.address,
      birthday: user.birthday,
      gender: user.gender
    })
    
    let updatedUser
    
    if (photoFile.value && photoFile.value !== 'REMOVE') {
      // Update with new photo
      const formData = new FormData()
      formData.append('name', user.name)
      formData.append('email', user.email)
      formData.append('nic', user.nic)
      formData.append('address', user.address)
      formData.append('birthday', user.birthday)
      formData.append('gender', user.gender)
      formData.append('photo', photoFile.value)
      
      console.log('Submitting with photo, birthday value:', user.birthday)
      const response = await updateUserWithPhoto(props.userData.id, formData)
      updatedUser = response.data
    } else {
      // Update without photo or remove photo
      const userData = {
        name: user.name,
        email: user.email,
        nic: user.nic,
        address: user.address,
        birthday: user.birthday,
        gender: user.gender
      }
      
      console.log('Submitting without photo, userData:', userData)
      const response = await updateUser(props.userData.id, userData)
      updatedUser = response.data
    }
    
    console.log('Update response:', updatedUser)
    emit('updated', updatedUser)
  } catch (error) {
    console.error('Error updating user:', error)
    alert('Error updating user: ' + (error.response?.data?.error || error.message))
  } finally {
    loading.value = false
  }
}

const cancel = () => {
  emit('cancelled')
}
</script>

<style scoped>
.user-edit-form {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.user-edit-form h2 {
  margin-bottom: 2rem;
  color: #333;
  text-align: center;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: #555;
}

.form-group input,
.form-group select,
.form-group textarea {
  padding: 0.75rem;
  border: 2px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #007bff;
}

.form-group input:disabled,
.form-group input[readonly] {
  background-color: #f8f9fa;
  color: #6c757d;
  cursor: not-allowed;
  border-color: #dee2e6;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.photo-upload-section {
  margin: 1rem 0;
}

.photo-upload-container {
  border: 2px dashed #ddd;
  border-radius: 8px;
  padding: 1rem;
  text-align: center;
  transition: border-color 0.3s;
}

.photo-upload-container:hover {
  border-color: #007bff;
}

.photo-preview {
  position: relative;
  display: inline-block;
}

.preview-image {
  max-width: 200px;
  max-height: 200px;
  border-radius: 8px;
  object-fit: cover;
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
}

.photo-upload-area {
  padding: 2rem;
  cursor: pointer;
  color: #666;
}

.photo-upload-area i {
  font-size: 2rem;
  margin-bottom: 0.5rem;
  display: block;
}

.photo-upload-buttons {
  margin-top: 1rem;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 2rem;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background-color: #0056b3;
}

.btn-primary:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background-color: #545b62;
}

.btn-outline {
  background-color: transparent;
  color: #007bff;
  border: 2px solid #007bff;
}

.btn-outline:hover {
  background-color: #007bff;
  color: white;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .user-edit-form {
    padding: 1rem;
    margin: 1rem;
  }
  
  .form-actions {
    flex-direction: column;
  }
}
</style>
