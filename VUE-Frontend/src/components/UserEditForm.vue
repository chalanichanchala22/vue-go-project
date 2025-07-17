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
            readonly
            disabled
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
  max-width: 700px;
  margin: 0 auto;
  padding: 2.5rem;
  background: linear-gradient(145deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  border-radius: 16px;
  box-shadow: 0 10px 40px rgba(102, 126, 234, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
}

.user-edit-form::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0.05) 100%);
  backdrop-filter: blur(10px);
  z-index: 0;
}

.user-edit-form > * {
  position: relative;
  z-index: 1;
}

.user-edit-form h2 {
  text-align: center;
  color: #ffffff;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.5);
  margin-bottom: 2.5rem;
  font-size: 2rem;
  font-weight: 700;
  letter-spacing: 0.5px;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 1.8rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.form-group label {
  font-weight: 700;
  color: #ffffff;
  font-size: 1rem;
  margin-bottom: 0.4rem;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);
  letter-spacing: 0.3px;
}

.form-group input,
.form-group select,
.form-group textarea {
  padding: 0.9rem 1rem;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.3s ease;
  background-color: rgba(255, 255, 255, 0.95);
  color: #1f2937;
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #ffffff;
  box-shadow: 0 0 0 3px rgba(255, 255, 255, 0.3);
  background-color: #ffffff;
  transform: translateY(-1px);
}

.form-group input:hover,
.form-group select:hover,
.form-group textarea:hover {
  border-color: rgba(255, 255, 255, 0.6);
  background-color: #ffffff;
}

.form-group input:disabled,
.form-group input[readonly] {
  background-color: rgba(248, 249, 250, 0.8);
  color: #6c757d;
  cursor: not-allowed;
  border-color: rgba(222, 226, 230, 0.8);
  text-shadow: none;
}

.form-group textarea {
  resize: vertical;
  min-height: 90px;
  font-family: inherit;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.2rem;
}

.photo-upload-section {
  margin: 1rem 0;
}

.photo-upload-container {
  position: relative;
  width: 180px;
  height: 180px;
  border: 3px dashed rgba(255, 255, 255, 0.5);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
}

.photo-upload-container:hover {
  border-color: rgba(255, 255, 255, 0.8);
  background: rgba(255, 255, 255, 0.2);
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(255, 255, 255, 0.2);
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
  border-radius: 8px;
}

.remove-photo-btn {
  position: absolute;
  top: -10px;
  right: -10px;
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  color: white;
  border: none;
  border-radius: 50%;
  width: 28px;
  height: 28px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.3rem;
  line-height: 1;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(239, 68, 68, 0.3);
}

.remove-photo-btn:hover {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.4);
}

.photo-upload-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  text-align: center;
  color: #ffffff;
  padding: 1rem;
  text-shadow: 0 1px 4px rgba(0, 0, 0, 0.4);
}

.photo-upload-area i {
  font-size: 2.5rem;
  margin-bottom: 0.8rem;
  filter: brightness(1.2);
}

.photo-upload-area p {
  margin: 0.3rem 0;
  font-size: 1rem;
  font-weight: 600;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);
}

.photo-upload-area small {
  font-size: 0.85rem;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 500;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

.photo-upload-buttons {
  margin-top: 1rem;
}

.form-actions {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
  margin-top: 2rem;
  padding-top: 1.5rem;
  border-top: 1px solid rgba(255, 255, 255, 0.3);
}

.btn {
  padding: 1rem 2.5rem;
  border: none;
  border-radius: 10px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  min-width: 160px;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  position: relative;
  overflow: hidden;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
}

.btn-primary {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%);
  color: #1f2937;
  border: 2px solid rgba(255, 255, 255, 0.8);
  font-weight: 700;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  letter-spacing: 0.3px;
}

.btn-primary::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.btn-primary:hover:not(:disabled) {
  background: linear-gradient(135deg, #f8fafc 0%, #e5e7eb 100%);
  color: #111827;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
  border-color: #ffffff;
}

.btn-primary:hover:not(:disabled)::before {
  left: 100%;
}

.btn-primary:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #6b7280 100%);
  color: #ffffff;
  cursor: not-allowed;
  transform: none;
  box-shadow: 0 2px 8px rgba(156, 163, 175, 0.2);
  border-color: #9ca3af;
}

.btn-secondary {
  background: linear-gradient(135deg, #6b7280 0%, #4b5563 100%);
  color: white;
  box-shadow: 0 4px 15px rgba(107, 114, 128, 0.3);
}

.btn-secondary:hover {
  background: linear-gradient(135deg, #4b5563 0%, #374151 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(107, 114, 128, 0.4);
}

.btn-outline {
  background: rgba(255, 255, 255, 0.1);
  color: #ffffff;
  border: 2px solid rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(5px);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

.btn-outline:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: #ffffff;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(255, 255, 255, 0.2);
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .user-edit-form {
    padding: 2rem 1.5rem;
    margin: 1rem;
    border-radius: 12px;
  }
  
  .user-edit-form h2 {
    font-size: 1.7rem;
    margin-bottom: 2rem;
  }
  
  .form-actions {
    flex-direction: column;
    align-items: center;
  }
  
  .btn {
    width: 100%;
    max-width: 300px;
  }
  
  .photo-upload-container {
    width: 150px;
    height: 150px;
  }
}

@media (max-width: 480px) {
  .user-edit-form {
    padding: 1.5rem 1rem;
    margin: 0.5rem;
  }
  
  .form {
    gap: 1.5rem;
  }
  
  .form-group input,
  .form-group select,
  .form-group textarea {
    padding: 0.8rem;
  }
  
  .btn {
    padding: 0.8rem 2rem;
    font-size: 1rem;
  }
}
</style>
