<template>
  <div class="dashboard">
    <div class="unified-dashboard-card">
      <!-- Dashboard Header -->
      <div class="dashboard-header">
        <h1>Dashboard</h1>
        <p>Welcome to your User Management Dashboard</p>
      </div>

      <!-- Stats Section -->
      <div class="stats-section">
        <div class="stat-card users-card">
          <div class="stat-icon">
            <i class="fas fa-users"></i>
          </div>
          <div class="stat-content">
            <h3>{{ totalUsers }}</h3>
            <p>Total Users</p>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="quick-actions">
        <h2>Quick Actions</h2>
        <div class="actions-grid">
          <router-link to="/users" class="action-card add-user">
            <div class="action-icon">
              <i class="fas fa-user-plus"></i>
            </div>
            <div class="action-content">
              <h3>Add New User</h3>
              <p>Create a new user account</p>
            </div>
          </router-link>

          <router-link to="/userlist" class="action-card view-users">
            <div class="action-icon">
              <i class="fas fa-list"></i>
            </div>
            <div class="action-content">
              <h3>View All Users</h3>
              <p>Browse and manage users</p>
            </div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getUsers } from '../services/userService'

// Reactive data
const totalUsers = ref(0)
const newUsersThisMonth = ref(0)

// Methods
const loadDashboardData = async () => {
  try {
    const response = await getUsers()
    const users = response.data
    
    totalUsers.value = users.length
    
    // Calculate new users this month
    const thisMonth = new Date()
    thisMonth.setDate(1)
    
    newUsersThisMonth.value = users.filter(user => {
      const createdAt = new Date(user.created_at)
      return createdAt >= thisMonth
    }).length
    
  } catch (error) {
    console.error('Error loading dashboard data:', error)
    // Set default values if API fails
    totalUsers.value = 25
    newUsersThisMonth.value = 12
  }
}

onMounted(() => {
  loadDashboardData()
})
</script>

<style scoped>
.dashboard {
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
  min-height: calc(100vh - 200px);
}

.unified-dashboard-card {
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(15px);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  padding: 2rem;
}

.dashboard-header {
  text-align: center;
  margin-bottom: 2rem;
}

.dashboard-header h1 {
  color: #1f2937;
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 0.5rem;
  text-shadow: 0 3px 10px rgba(255, 255, 255, 0.3);
  letter-spacing: 0.5px;
}

.dashboard-header p {
  color: #374151;
  font-size: 1.2rem;
  font-weight: 600;
  text-shadow: 0 2px 6px rgba(255, 255, 255, 0.2);
  opacity: 0.95;
}

/* Stats Section */
.stats-section {
  margin-bottom: 2rem;
}

.stat-card {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.15) 0%, rgba(29, 78, 216, 0.1) 100%);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(59, 130, 246, 0.3);
  padding: 1.5rem;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  transition: all 0.3s ease;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.1);
  max-width: 400px;
  margin: 0 auto;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(59, 130, 246, 0.2);
  border-color: rgba(59, 130, 246, 0.4);
}

.stat-icon {
  width: 80px;
  height: 80px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  color: white;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.users-card .stat-icon {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%);
}

.stat-content h3 {
  font-size: 2.5rem;
  font-weight: 800;
  color: #111827;
  margin: 0 0 0.5rem 0;
  text-shadow: 0 3px 8px rgba(255, 255, 255, 0.4);
}

.stat-content p {
  font-size: 1.1rem;
  color: #1f2937;
  margin: 0 0 0.8rem 0;
  font-weight: 700;
  text-shadow: 0 2px 4px rgba(255, 255, 255, 0.3);
  opacity: 0.95;
}

.stat-trend {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  font-weight: 700;
  text-shadow: 0 2px 4px rgba(255, 255, 255, 0.3);
}

.stat-trend.positive {
  color: #059669;
  text-shadow: 0 2px 4px rgba(5, 150, 105, 0.4);
}

.stat-trend.neutral {
  color: #d97706;
  opacity: 0.9;
}

/* Quick Actions */
.quick-actions {
  margin-top: 2rem;
}

.quick-actions h2 {
  color: #111827;
  font-size: 1.8rem;
  font-weight: 800;
  margin-bottom: 1.5rem;
  text-shadow: 0 3px 8px rgba(255, 255, 255, 0.4);
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1.5rem;
}

.action-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  padding: 1.5rem;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  text-decoration: none;
  transition: all 0.3s ease;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
}

.add-user.action-card {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.12) 0%, rgba(29, 78, 216, 0.08) 100%);
  border: 1px solid rgba(59, 130, 246, 0.25);
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.08);
}

.add-user.action-card:hover {
  box-shadow: 0 12px 40px rgba(59, 130, 246, 0.15);
  border-color: rgba(59, 130, 246, 0.35);
}

.view-users.action-card {
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.12) 0%, rgba(4, 120, 87, 0.08) 100%);
  border: 1px solid rgba(16, 185, 129, 0.25);
  box-shadow: 0 4px 16px rgba(16, 185, 129, 0.08);
}

.view-users.action-card:hover {
  box-shadow: 0 12px 40px rgba(16, 185, 129, 0.15);
  border-color: rgba(16, 185, 129, 0.35);
}

.action-card:hover {
  transform: translateY(-4px);
}

.action-icon {
  width: 60px;
  height: 60px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  color: white;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.add-user .action-icon {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%);
}

.view-users .action-icon {
  background: linear-gradient(135deg, #10b981 0%, #047857 100%);
}

.action-content h3 {
  color: #1f2937;
  font-size: 1.2rem;
  font-weight: 800;
  margin: 0 0 0.5rem 0;
  text-shadow: 0 2px 6px rgba(255, 255, 255, 0.4);
}

.action-content p {
  color: #374151;
  font-size: 0.95rem;
  margin: 0;
  text-shadow: 0 2px 4px rgba(255, 255, 255, 0.3);
  font-weight: 600;
  opacity: 0.95;
}

/* Responsive Design */
@media (max-width: 768px) {
  .dashboard {
    padding: 1rem;
  }
  
  .unified-dashboard-card {
    padding: 1.5rem;
  }
  
  .dashboard-header h1 {
    font-size: 2rem;
  }
  
  .actions-grid {
    grid-template-columns: 1fr;
  }
  
  .stat-card, .action-card {
    padding: 1.5rem;
  }
}

@media (max-width: 480px) {
  .dashboard {
    padding: 0.5rem;
  }
  
  .unified-dashboard-card {
    padding: 1rem;
  }
  
  .dashboard-header h1 {
    font-size: 1.7rem;
  }
  
  .stat-card, .action-card {
    padding: 1rem;
    gap: 1rem;
  }
  
  .stat-icon, .action-icon {
    width: 50px;
    height: 50px;
    font-size: 1.2rem;
  }
  
  .stat-content h3 {
    font-size: 2rem;
  }
}
</style>
