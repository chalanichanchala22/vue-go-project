import axios from 'axios'
const API_URL = 'http://localhost:8080/api'

export const getPhonesByUser = (userId) => axios.get(`${API_URL}/users/${userId}/phones`)
export const createPhone = (data) => axios.post(`${API_URL}/users/${data.user_id}/phones`, data)
export const deletePhone = (userId, phoneId) => axios.delete(`${API_URL}/users/${userId}/phones/${phoneId}`)
