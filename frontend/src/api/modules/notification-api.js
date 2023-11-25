import axiosClient from '../axios'

export function getAll() {
  return axiosClient.get('/notifications/')
}

export function setAllRead() {
  return axiosClient.post('/notifications/set-all-read')
}

export function setRead(id) {
  return axiosClient.post(`/notifications/${id}/set-read`)
}
