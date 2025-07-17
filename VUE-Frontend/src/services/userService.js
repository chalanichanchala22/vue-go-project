import axios from 'axios'
import authService from './authService'

const API_URL = 'http://localhost:8080/api/users'

// Configure axios to include auth headers
const getAuthHeaders = () => {
  const token = authService.getToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}

export const getUsers = () => axios.get(API_URL, { headers: getAuthHeaders() })
export const getUsersWithPhones = () => axios.get(`${API_URL}/with-phones`, { headers: getAuthHeaders() })
export const getUser = (id) => axios.get(`${API_URL}/${id}`, { headers: getAuthHeaders() })
export const createUser = (data) => axios.post(API_URL, data, { headers: getAuthHeaders() })
export const createUserWithPhoto = (formData) => {
  return axios.post(API_URL, formData, {
    headers: {
      ...getAuthHeaders(),
      'Content-Type': 'multipart/form-data'
    }
  })
}
export const updateUser = (id, data) => axios.put(`${API_URL}/${id}`, data, { headers: getAuthHeaders() })
export const updateUserWithPhoto = (id, formData) => {
  return axios.put(`${API_URL}/${id}`, formData, {
    headers: {
      ...getAuthHeaders(),
      'Content-Type': 'multipart/form-data'
    }
  })
}
export const deleteUser = (id) => axios.delete(`${API_URL}/${id}`, { headers: getAuthHeaders() })
