import axios from 'axios'

// The HTTP request method should be POST with request header:
// Content-Type: multipart/form-data

const apiClient = axios.create({
  baseURL: 'http://localhost:8080',
})

export const embedRequest = async (formData) => {
  return apiClient.post('/api/embed', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  })
}
