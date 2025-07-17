const API_BASE_URL = 'http://localhost:8080/api'

class AuthService {
  constructor() {
    this.token = localStorage.getItem('token')
  }

  async login(email, password) {
    try {
      console.log('Attempting login with:', { email, password })
      
      const response = await fetch(`${API_BASE_URL}/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password })
      })

      console.log('Response status:', response.status)
      console.log('Response headers:', response.headers)

      if (!response.ok) {
        const errorText = await response.text()
        console.log('Error response:', errorText)
        
        let error
        try {
          error = JSON.parse(errorText)
        } catch {
          error = { message: errorText || 'Login failed' }
        }
        throw new Error(error.error || error.message || 'Login failed')
      }

      const data = await response.json()
      console.log('Login successful, data:', data)
      
      if (data.token) {
        this.token = data.token
        localStorage.setItem('token', data.token)
        localStorage.setItem('user', JSON.stringify(data.user || { email }))
        return data
      } else {
        throw new Error('No token received from server')
      }
    } catch (error) {
      console.error('Login error:', error)
      throw error
    }
  }

  logout() {
    this.token = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  isAuthenticated() {
    return !!this.token
  }

  getToken() {
    return this.token
  }

  getUser() {
    const user = localStorage.getItem('user')
    return user ? JSON.parse(user) : null
  }

  // Add authorization header to requests
  getAuthHeaders() {
    return {
      'Authorization': `Bearer ${this.token}`,
      'Content-Type': 'application/json'
    }
  }
}

export default new AuthService()
