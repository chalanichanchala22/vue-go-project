import axios from 'axios'
const API_URL = 'http://localhost:8080/api/users'

export const getUsers = () => axios.get(API_URL)
export const createUser = (data) => axios.post(API_URL, data)
export const deleteUser = (id) => axios.delete(`${API_URL}/${id}`)
