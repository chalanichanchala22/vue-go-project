import axios from 'axios'
import authService from './authService'

const API_URL = 'http://localhost:8080/api'

// Configure axios to include auth headers
const getAuthHeaders = () => {
  const token = authService.getToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}

export const getPhonesByUser = (userId) => axios.get(`${API_URL}/users/${userId}/phones`, { headers: getAuthHeaders() })
export const createPhone = (userId, data) => axios.post(`${API_URL}/users/${userId}/phones`, data, { headers: getAuthHeaders() })
export const updatePhone = (userId, phoneId, data) => axios.put(`${API_URL}/users/${userId}/phones/${phoneId}`, data, { headers: getAuthHeaders() })
export const deletePhone = (userId, phoneId) => axios.delete(`${API_URL}/users/${userId}/phones/${phoneId}`, { headers: getAuthHeaders() })
