<template>
  <nav class="navbar">
    <div class="nav-container">
      <!-- Brand/Logo Section -->
      <div class="nav-brand">
        <router-link to="/dashboard" class="brand-link">
          <i class="fas fa-tachometer-alt brand-icon"></i>
          <span class="brand-text">User Management</span>
        </router-link>
      </div>

      <!-- Mobile Menu Toggle -->
      <button class="mobile-toggle" @click="toggleMobileMenu" :class="{ active: isMobileMenuOpen }">
        <span class="hamburger-line"></span>
        <span class="hamburger-line"></span>
        <span class="hamburger-line"></span>
      </button>

      <!-- Navigation Menu -->
      <div class="nav-menu" :class="{ active: isMobileMenuOpen }">
        <div class="nav-links">
          <router-link to="/dashboard" class="nav-link" @click="closeMobileMenu">
            <i class="fas fa-tachometer-alt"></i>
            <span>Dashboard</span>
          </router-link>
          
          <router-link to="/users" class="nav-link" @click="closeMobileMenu">
            <i class="fas fa-user-plus"></i>
            <span>Add Users</span>
          </router-link>
          
          <router-link to="/userlist" class="nav-link" @click="closeMobileMenu">
            <i class="fas fa-list"></i>
            <span>User List</span>
          </router-link>
          


          <!-- Dropdown Menu -->
          <div class="dropdown" ref="dropdown">
            <button class="dropdown-toggle" @click="toggleDropdown" :class="{ active: isDropdownOpen }">
              <i class="fas fa-ellipsis-h"></i>
              <span>More</span>
              <i class="fas fa-chevron-down dropdown-arrow"></i>
            </button>
            <div class="dropdown-menu" :class="{ show: isDropdownOpen }">
              <a href="#" class="dropdown-item" @click="closeMobileMenu">
                <i class="fas fa-chart-bar"></i>
                <span>Analytics</span>
              </a>
              <a href="#" class="dropdown-item" @click="closeMobileMenu">
                <i class="fas fa-cog"></i>
                <span>Settings</span>
              </a>
              <a href="#" class="dropdown-item" @click="closeMobileMenu">
                <i class="fas fa-question-circle"></i>
                <span>Help & Support</span>
              </a>
              <div class="dropdown-divider"></div>
              <a href="#" class="dropdown-item" @click="closeMobileMenu">
                <i class="fas fa-info-circle"></i>
                <span>About</span>
              </a>
            </div>
          </div>
        </div>

        <!-- User Profile Section -->
        <div class="nav-user">
          <div class="user-info">
            <div class="user-avatar">
              <i class="fas fa-user-circle"></i>
            </div>
            <div class="user-details">
              <span class="user-name">{{ userName }}</span>
              <span class="user-role">Administrator</span>
            </div>
          </div>
          
          <button @click="handleLogout" class="logout-btn">
            <i class="fas fa-sign-out-alt"></i>
            <span>Logout</span>
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const emit = defineEmits(['logout'])

const isMobileMenuOpen = ref(false)
const isDropdownOpen = ref(false)
const userName = ref('John Doe')
const dropdown = ref(null)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
  if (isMobileMenuOpen.value) {
    isDropdownOpen.value = false
  }
}

const closeMobileMenu = () => {
  isMobileMenuOpen.value = false
  isDropdownOpen.value = false
}

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
}

const handleLogout = () => {
  closeMobileMenu()
  emit('logout')
}

// Close dropdown when clicking outside
const handleClickOutside = (event) => {
  if (dropdown.value && !dropdown.value.contains(event.target)) {
    isDropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  // You can fetch the actual user name here
  // userName.value = authService.getCurrentUser()?.name || 'User'
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.navbar {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
}

/* Brand Section */
.nav-brand .brand-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  text-decoration: none;
  color: white;
  font-size: 1.5rem;
  font-weight: 700;
  transition: all 0.3s ease;
}

.nav-brand .brand-link:hover {
  transform: translateY(-1px);
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.brand-icon {
  font-size: 1.75rem;
  color: #fbbf24;
}

.brand-text {
  background: linear-gradient(45deg, #ffffff, #f3f4f6);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

/* Mobile Toggle */
.mobile-toggle {
  display: none;
  flex-direction: column;
  justify-content: center;
  width: 30px;
  height: 30px;
  background: none;
  border: none;
  cursor: pointer;
  position: relative;
}

.hamburger-line {
  width: 25px;
  height: 3px;
  background-color: white;
  margin: 3px 0;
  transition: all 0.3s ease;
  border-radius: 2px;
}

.mobile-toggle.active .hamburger-line:nth-child(1) {
  transform: rotate(45deg) translate(6px, 6px);
}

.mobile-toggle.active .hamburger-line:nth-child(2) {
  opacity: 0;
}

.mobile-toggle.active .hamburger-line:nth-child(3) {
  transform: rotate(-45deg) translate(6px, -6px);
}

/* Navigation Menu */
.nav-menu {
  display: flex;
  align-items: center;
  gap: 2rem;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: rgba(255, 255, 255, 0.9);
  text-decoration: none;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  font-weight: 500;
  position: relative;
  overflow: hidden;
}

.nav-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.nav-link:hover::before {
  left: 100%;
}

.nav-link:hover,
.nav-link.router-link-active {
  color: white;
  background-color: rgba(255, 255, 255, 0.15);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.nav-link i {
  font-size: 1rem;
}

/* Dropdown */
.dropdown {
  position: relative;
}

.dropdown-toggle {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: rgba(255, 255, 255, 0.9);
  background: none;
  border: none;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: 500;
}

.dropdown-toggle:hover,
.dropdown-toggle.active {
  color: white;
  background-color: rgba(255, 255, 255, 0.15);
}

.dropdown-arrow {
  transition: transform 0.3s ease;
  font-size: 0.8rem;
}

.dropdown-toggle.active .dropdown-arrow {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  border-radius: 12px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
  padding: 0.5rem 0;
  min-width: 200px;
  opacity: 0;
  visibility: hidden;
  transform: translateY(-10px);
  transition: all 0.3s ease;
  border: 1px solid rgba(0, 0, 0, 0.1);
}

.dropdown-menu.show {
  opacity: 1;
  visibility: visible;
  transform: translateY(0);
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  color: #374151;
  text-decoration: none;
  transition: all 0.2s ease;
}

.dropdown-item:hover {
  background-color: #f3f4f6;
  color: #1f2937;
}

.dropdown-item i {
  color: #6b7280;
  width: 16px;
}

.dropdown-divider {
  height: 1px;
  background-color: #e5e7eb;
  margin: 0.5rem 0;
}

/* User Section */
.nav-user {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: white;
}

.user-avatar {
  font-size: 2rem;
  color: #fbbf24;
}

.user-details {
  display: flex;
  flex-direction: column;
  text-align: right;
}

.user-name {
  font-weight: 600;
  font-size: 0.9rem;
}

.user-role {
  font-size: 0.8rem;
  color: rgba(255, 255, 255, 0.8);
}

.logout-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: linear-gradient(45deg, #ef4444, #dc2626);
  color: white;
  border: none;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(239, 68, 68, 0.3);
}

.logout-btn:hover {
  background: linear-gradient(45deg, #dc2626, #b91c1c);
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(239, 68, 68, 0.4);
}

/* Responsive Design */
@media (max-width: 1024px) {
  .nav-links {
    gap: 1rem;
  }
  
  .nav-link span,
  .dropdown-toggle span,
  .logout-btn span {
    display: none;
  }
  
  .user-details {
    display: none;
  }
}

@media (max-width: 768px) {
  .mobile-toggle {
    display: flex;
  }
  
  .nav-menu {
    position: fixed;
    top: 70px;
    left: 0;
    width: 100%;
    height: calc(100vh - 70px);
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    flex-direction: column;
    justify-content: flex-start;
    padding: 2rem 1rem;
    gap: 1rem;
    transform: translateX(-100%);
    transition: transform 0.3s ease;
  }
  
  .nav-menu.active {
    transform: translateX(0);
  }
  
  .nav-links {
    flex-direction: column;
    width: 100%;
    gap: 0.5rem;
  }
  
  .nav-link,
  .dropdown-toggle {
    width: 100%;
    justify-content: flex-start;
    padding: 1rem;
    border-radius: 8px;
    background-color: rgba(255, 255, 255, 0.1);
  }
  
  .nav-link span,
  .dropdown-toggle span,
  .logout-btn span {
    display: inline;
  }
  
  .dropdown-menu {
    position: static;
    opacity: 1;
    visibility: visible;
    transform: none;
    background: rgba(255, 255, 255, 0.95);
    box-shadow: inset 0 2px 8px rgba(0, 0, 0, 0.1);
  }
  
  .user-info {
    order: -1;
    padding: 1rem;
    border-radius: 8px;
    background-color: rgba(255, 255, 255, 0.1);
    width: 100%;
    justify-content: center;
  }
  
  .user-details {
    display: flex;
    text-align: center;
  }
  
  .nav-user {
    flex-direction: column;
    width: 100%;
    gap: 1rem;
  }
}
</style>
