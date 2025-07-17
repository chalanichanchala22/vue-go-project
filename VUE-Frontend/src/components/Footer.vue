<template>
  <footer class="footer">
    <div class="footer-container">
      <div class="footer-content">
        <div class="footer-section">
          <h3 class="footer-title">User Management System</h3>
          <p class="footer-description">
            A comprehensive solution for managing users and their information efficiently.
          </p>
        </div>
        
        <div class="footer-section">
          <h4 class="footer-subtitle">Quick Links</h4>
          <ul class="footer-links">
            <li><router-link to="/users" class="footer-link">Users</router-link></li>
            <li><router-link to="/userlist" class="footer-link">User List</router-link></li>
            <li><a href="#" class="footer-link">Help & Support</a></li>
            <li><a href="#" class="footer-link">Privacy Policy</a></li>
          </ul>
        </div>
        
        <div class="footer-section">
          <h4 class="footer-subtitle">Contact Info</h4>
          <div class="contact-info">
            <p class="contact-item">
              <i class="fas fa-envelope"></i>
              <span>support@usermanagement.com</span>
            </p>
            <p class="contact-item">
              <i class="fas fa-phone"></i>
              <span>+94 77 6060 100</span>
            </p>
            <p class="contact-item">
              <i class="fas fa-map-marker-alt"></i>
              <span>123 Tech Street,Kegalle</span>
            </p>
          </div>
        </div>
        
        <div class="footer-section">
          <h4 class="footer-subtitle">Follow Us</h4>
          <div class="social-links">
            <a href="https://facebook.com" class="social-link" aria-label="Facebook" target="_blank">
              <i class="fab fa-facebook-f"></i>
            </a>
            <a href="https://linkedin.com" class="social-link" aria-label="LinkedIn" target="_blank">
              <i class="fab fa-linkedin-in"></i>
            </a>
            <a href="https://github.com" class="social-link" aria-label="GitHub" target="_blank">
              <i class="fab fa-github"></i>
            </a>
            <a href="https://instagram.com" class="social-link" aria-label="Instagram" target="_blank">
              <i class="fab fa-instagram"></i>
            </a>
          </div>
        </div>
      </div>
      
      <div class="footer-bottom">
        <div class="footer-bottom-content">
          <p class="copyright">
            &copy; {{ currentYear }} User Management System. All rights reserved.
          </p>
          <div class="footer-bottom-links">
            <a href="#" class="bottom-link">Terms of Service</a>
            <span class="separator">|</span>
            <a href="#" class="bottom-link">Privacy Policy</a>
            <span class="separator">|</span>
            <a href="#" class="bottom-link">Cookie Policy</a>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<style scoped>
.footer {
  background: linear-gradient(135deg, #1f2937 0%, #111827 100%);
  color: #e5e7eb;
  margin-top: auto;
}

.footer-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.footer-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  padding: 3rem 0 2rem 0;
}

.footer-section {
  display: flex;
  flex-direction: column;
}

.footer-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 1rem;
}

.footer-subtitle {
  font-size: 1.125rem;
  font-weight: 600;
  color: #ffffff;
  margin-bottom: 1rem;
}

.footer-description {
  color: #d1d5db;
  line-height: 1.6;
  margin-bottom: 1rem;
}

.footer-links {
  list-style: none;
  padding: 0;
  margin: 0;
}

.footer-links li {
  margin-bottom: 0.5rem;
}

.footer-link {
  color: #d1d5db;
  text-decoration: none;
  transition: color 0.15s ease;
}

.footer-link:hover {
  color: #60a5fa;
}

.contact-info {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #d1d5db;
}

.contact-item i {
  width: 16px;
  height: 16px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #60a5fa;
}

.social-links {
  display: flex;
  gap: 1rem;
}

.social-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background-color: #374151;
  color: #d1d5db;
  text-decoration: none;
  border-radius: 50%;
  transition: all 0.15s ease;
}

.social-link:hover {
  background-color: #60a5fa;
  color: #ffffff;
  transform: translateY(-2px);
}

/* Individual social media colors on hover */
.social-link[aria-label="Facebook"]:hover {
  background-color: #1877f2;
}

.social-link[aria-label="LinkedIn"]:hover {
  background-color: #0077b5;
}

.social-link[aria-label="GitHub"]:hover {
  background-color: #333;
}

.social-link[aria-label="Instagram"]:hover {
  background: linear-gradient(45deg, #f09433 0%,#e6683c 25%,#dc2743 50%,#cc2366 75%,#bc1888 100%);
}

.footer-bottom {
  border-top: 1px solid #374151;
  padding: 1.5rem 0;
}

.footer-bottom-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 1rem;
}

.copyright {
  color: #9ca3af;
  font-size: 0.875rem;
}

.footer-bottom-links {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.bottom-link {
  color: #9ca3af;
  text-decoration: none;
  font-size: 0.875rem;
  transition: color 0.15s ease;
}

.bottom-link:hover {
  color: #d1d5db;
}

.separator {
  color: #6b7280;
}

/* Responsive Design */
@media (max-width: 768px) {
  .footer-content {
    grid-template-columns: 1fr;
    gap: 1.5rem;
    padding: 2rem 0 1.5rem 0;
  }
  
  .footer-bottom-content {
    flex-direction: column;
    text-align: center;
  }
  
  .footer-bottom-links {
    flex-wrap: wrap;
    justify-content: center;
  }
}

@media (max-width: 480px) {
  .footer-container {
    padding: 0 0.5rem;
  }
  
  .footer-content {
    padding: 1.5rem 0 1rem 0;
  }
  
  .social-links {
    justify-content: center;
  }
}
</style>
