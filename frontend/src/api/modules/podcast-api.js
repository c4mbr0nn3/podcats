import axiosClient from '../axios'

export function getAll() {
  return axiosClient.get('/podcasts/')
}

export function getById(podcastId) {
  return axiosClient.get(`/podcasts/${podcastId}`)
}

export function getItemsById(podcastId, pageId) {
  return axiosClient.get(`/podcasts/${podcastId}/items?page=${pageId}`)
}

export function deleteById(podcastId) {
  return axiosClient.delete(`/podcasts/${podcastId}/remove`)
}

export function setPlayed(podcastId) {
  return axiosClient.post(`/podcasts/${podcastId}/set-all-played`)
}

export function importPodcast(payload) {
  return axiosClient.post('/podcasts/import', payload)
}
