import axios from 'axios'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { showErrorSnackbar } from '@/utils/snackbar'

const baseURL = import.meta.env.DEV
  ? 'http://localhost:8000/api/v1'
  : '/api/v1'

const axiosOptions = {
  baseURL,
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
  },
}

const axiosClient = axios.create(axiosOptions)

axiosClient.interceptors.request.use((config) => {
  const authStore = useAuthStore()
  if (authStore.isAuthenticated)
    config.headers.Authorization = `Bearer ${authStore.token}`
  return config
})

axiosClient.interceptors.response.use(
  res => res,
  (err) => {
    const authStore = useAuthStore()
    const router = useRouter()
    const messager = err.response.data?.error || getGenericMessage(err)
    showErrorSnackbar(messager)
    if (err.response.status === 401)
      authStore.logout()
    if (err.response.status === 403)
      router.push({ name: 'home' })
    return Promise.reject(err)
  },
)

function getGenericMessage(err) {
  const status = err.response.status
  switch (status) {
    case 401:
      return 'Unauthorized, please login'
    case 403:
      return 'Forbidden, you don\'t have permission to access this resource'
    case 404:
      return 'Not Found, resource not found'
    case 500:
      return 'Internal Server Error, something went wrong'
    default:
      return 'Something went wrong, please try again later'
  }
}

export default axiosClient
