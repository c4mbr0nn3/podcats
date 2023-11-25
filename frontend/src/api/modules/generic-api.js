import axios from 'axios'

export function get(url, query) {
  return axios.get(url, query)
}

export function post(url, data) {
  return axios.post(url, data)
}
