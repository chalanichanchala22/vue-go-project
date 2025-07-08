import axios from 'axios'
const API_URL = 'http://localhost:8080/api/users'

export const getUsers = () => axios.get(API_URL)
export const getUser = (id) => axios.get(`${API_URL}/${id}`)
export const createUser = (data) => axios.post(API_URL, data)
export const createUserWithPhoto = (formData) => {
  return axios.post(API_URL, formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
export const updateUser = (id, data) => axios.put(`${API_URL}/${id}`, data)
export const updateUserWithPhoto = (id, formData) => {
  return axios.put(`${API_URL}/${id}`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
export const deleteUser = (id) => axios.delete(`${API_URL}/${id}`)
