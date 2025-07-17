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
        <label for="password">Password *</label>
        <input 
          id="password"
          v-model="user.password" 
          type="password"
          placeholder="Enter password" 
          required 
          minlength="6"
        />
      </div>

      <div class="form-group">
        <label for="confirmPassword">Confirm Password *</label>
        <input 
          id="confirmPassword"
          v-model="user.confirmPassword" 
          type="password"
          placeholder="Re-enter password" 
          required 
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
  name: '', email: '', password: '', confirmPassword: '', nic: '', address: '', birthday: '', gender: ''
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
    if (!user.value.name || !user.value.email || !user.value.password || 
        !user.value.confirmPassword || !user.value.nic || 
        !user.value.address || !user.value.birthday || !user.value.gender) {
      errorMessage.value = 'Please fill in all required fields'
      return
    }
    
    // Validate password match
    if (user.value.password !== user.value.confirmPassword) {
      errorMessage.value = 'Passwords do not match'
      return
    }
    
    // Validate password length
    if (user.value.password.length < 6) {
      errorMessage.value = 'Password must be at least 6 characters long'
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
    let response
    
    // Always use FormData to ensure consistent behavior
    const formData = new FormData()
    formData.append('name', user.value.name.trim())
    formData.append('email', user.value.email.trim())
    formData.append('password', user.value.password)
    formData.append('confirmPassword', user.value.confirmPassword)
    formData.append('nic', user.value.nic.trim())
    formData.append('address', user.value.address.trim())
    formData.append('birthday', user.value.birthday) // Send date in YYYY-MM-DD format
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
    user.value = { name: '', email: '', password: '', confirmPassword: '', nic: '', address: '', birthday: '', gender: '' }
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

.user-form::before {
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

.user-form > * {
  position: relative;
  z-index: 1;
}

.user-form h2 {
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

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.2rem;
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

.form-group textarea {
  resize: vertical;
  min-height: 90px;
  font-family: inherit;
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

.error-message {
  color: #991b1b;
  background: linear-gradient(135deg, #fef2f2 0%, #fee2e2 100%);
  border: 2px solid #fca5a5;
  padding: 1rem 1.5rem;
  border-radius: 8px;
  font-size: 1rem;
  width: 100%;
  text-align: center;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(220, 38, 38, 0.2);
  letter-spacing: 0.2px;
}

.btn-primary {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%);
  color: #1f2937;
  padding: 1rem 2.5rem;
  border: 2px solid rgba(255, 255, 255, 0.8);
  border-radius: 10px;
  font-size: 1.1rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s ease;
  min-width: 160px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
  position: relative;
  overflow: hidden;
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

.btn-primary:active:not(:disabled) {
  transform: translateY(0);
}

.button-group {
  display: flex;
  gap: 1.2rem;
  flex-wrap: wrap;
  justify-content: center;
}

.btn-secondary {
  background: linear-gradient(135deg, #6b7280 0%, #4b5563 100%);
  color: white;
  padding: 1rem 2.5rem;
  border: none;
  border-radius: 10px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  min-width: 160px;
  box-shadow: 0 4px 15px rgba(107, 114, 128, 0.3);
}

.btn-secondary:hover {
  background: linear-gradient(135deg, #4b5563 0%, #374151 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(107, 114, 128, 0.4);
}

.btn-secondary:active {
  transform: translateY(0);
}

/* Phone Numbers Section */
.phone-section {
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 12px;
  padding: 1.8rem;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.15) 0%, rgba(59, 130, 246, 0.15) 100%);
  backdrop-filter: blur(10px);
  box-shadow: inset 0 2px 4px rgba(255, 255, 255, 0.1);
  position: relative;
  overflow: hidden;
}

.phone-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, rgba(16, 185, 129, 0.05) 0%, rgba(59, 130, 246, 0.05) 50%, rgba(147, 51, 234, 0.05) 100%);
  pointer-events: none;
}

.phone-section > * {
  position: relative;
  z-index: 1;
}

.phone-section h3 {
  margin: 0 0 1.5rem 0;
  color: #ffffff;
  font-size: 1.4rem;
  font-weight: 800;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.5);
  letter-spacing: 0.3px;
  background: linear-gradient(135deg, #10b981 0%, #3b82f6 100%);
  padding: 1rem 1.5rem;
  border-radius: 8px;
  margin: -1.8rem -1.8rem 1.5rem -1.8rem;
  box-shadow: 0 4px 15px rgba(16, 185, 129, 0.3);
}

.phone-section h3::before {
  content: "📱";
  font-size: 1.2rem;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.3));
}

.phone-entry {
  margin-bottom: 1.2rem;
  padding: 1.2rem;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(248, 250, 252, 0.95) 100%);
  border: 2px solid transparent;
  background-clip: padding-box;
  border-radius: 10px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  backdrop-filter: blur(5px);
  position: relative;
  overflow: hidden;
}

.phone-entry::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, #10b981 0%, #3b82f6 50%, #8b5cf6 100%);
  border-radius: 10px;
  padding: 2px;
  z-index: -1;
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask-composite: exclude;
}

.phone-entry:hover {
  box-shadow: 0 6px 25px rgba(16, 185, 129, 0.2);
  transform: translateY(-2px);
}

.phone-entry:hover::before {
  background: linear-gradient(135deg, #059669 0%, #2563eb 50%, #7c3aed 100%);
}

.phone-entry:last-child {
  margin-bottom: 0;
}

/* Specific styling for phone number and type fields */
.phone-entry .form-group:first-child label {
  color: #10b981;
  font-weight: 800;
  text-shadow: 0 1px 3px rgba(16, 185, 129, 0.3);
}

.phone-entry .form-group:first-child input {
  border: 2px solid rgba(16, 185, 129, 0.4);
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.05) 0%, rgba(255, 255, 255, 0.95) 100%);
  color: #065f46;
  font-weight: 600;
}

.phone-entry .form-group:first-child input:focus {
  border-color: #10b981;
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.3);
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.1) 0%, #ffffff 100%);
}

.phone-entry .form-group:first-child input:hover {
  border-color: rgba(16, 185, 129, 0.6);
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.08) 0%, #ffffff 100%);
}

.phone-entry .form-group:nth-child(2) label {
  color: #3b82f6;
  font-weight: 800;
  text-shadow: 0 1px 3px rgba(59, 130, 246, 0.3);
}

.phone-entry .form-group:nth-child(2) select {
  border: 2px solid rgba(59, 130, 246, 0.4);
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.05) 0%, rgba(255, 255, 255, 0.95) 100%);
  color: #1e40af;
  font-weight: 600;
}

.phone-entry .form-group:nth-child(2) select:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.3);
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.1) 0%, #ffffff 100%);
}

.phone-entry .form-group:nth-child(2) select:hover {
  border-color: rgba(59, 130, 246, 0.6);
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.08) 0%, #ffffff 100%);
}

.btn-add, .btn-remove {
  padding: 0.7rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.btn-add {
  background: linear-gradient(135deg, #10b981 0%, #3b82f6 100%);
  color: white;
  margin-top: 1.2rem;
  position: relative;
  overflow: hidden;
}

.btn-add::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.btn-add:hover {
  background: linear-gradient(135deg, #059669 0%, #2563eb 100%);
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(16, 185, 129, 0.4);
}

.btn-add:hover::before {
  left: 100%;
}

.btn-remove {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  color: white;
  height: fit-content;
  margin-top: auto;
}

.btn-remove:hover:not(:disabled) {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}

.btn-remove:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #6b7280 100%);
  cursor: not-allowed;
  transform: none;
  box-shadow: 0 1px 4px rgba(156, 163, 175, 0.2);
}

/* Photo Upload Section */
.photo-upload-section {
  margin-bottom: 1.8rem;
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
  color: #ffffff;
  padding: 1rem;
  text-shadow: 0 1px 4px rgba(0, 0, 0, 0.4);
}

.upload-icon {
  font-size: 2.5rem;
  margin-bottom: 0.8rem;
  filter: brightness(1.2);
}

.upload-placeholder p {
  margin: 0.3rem 0;
  font-size: 1rem;
  font-weight: 600;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);
}

.upload-hint {
  font-size: 0.85rem !important;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 500 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
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

.remove-photo-btn span {
  margin-top: -2px;
}

/* Responsive design */
@media (max-width: 768px) {
  .user-form {
    padding: 2rem 1.5rem;
    margin: 1rem;
    border-radius: 12px;
  }
  
  .user-form h2 {
    font-size: 1.7rem;
    margin-bottom: 2rem;
  }
  
  .form-row {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
  
  .button-group {
    flex-direction: column;
    align-items: center;
  }
  
  .btn-primary, .btn-secondary {
    width: 100%;
    max-width: 300px;
  }
  
  .phone-section {
    padding: 1.5rem;
  }
  
  .photo-upload-container {
    width: 150px;
    height: 150px;
  }
}

@media (max-width: 480px) {
  .user-form {
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
  
  .btn-primary, .btn-secondary {
    padding: 0.8rem 2rem;
    font-size: 1rem;
  }
}
</style>
