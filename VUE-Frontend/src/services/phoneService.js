import axios from 'axios'
const API_URL = 'http://localhost:8080/api'

export const getPhonesByUser = (userId) => axios.get(`${API_URL}/users/${userId}/phones`)
export const createPhone = (userId, data) => axios.post(`${API_URL}/users/${userId}/phones`, data)
export const updatePhone = (userId, phoneId, data) => axios.put(`${API_URL}/users/${userId}/phones/${phoneId}`, data)
export const deletePhone = (userId, phoneId) => axios.delete(`${API_URL}/users/${userId}/phones/${phoneId}`)
