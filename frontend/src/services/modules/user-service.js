import { UserApi } from '@/api'
import { showErrorSnackbar, showSuccessSnackbar } from '@/utils/snackbar'

export async function getAll() {
  const { data } = await UserApi.getAll()
  return data
}

export async function create(user) {
  let data = null
  try {
    const { data: responseData } = await UserApi.create(user)
    data = responseData
    showSuccessSnackbar('User created')
  }
  catch {
    showErrorSnackbar('Error creating user')
  }
  return data
}

export async function update(id, user) {
  let data = null
  try {
    const { data: responseData } = await UserApi.update(id, user)
    data = responseData
    showSuccessSnackbar('User updated')
  }
  catch {
    showErrorSnackbar('Error updating user')
  }
  return data
}

export async function resetPassword(id) {
  let data = null
  try {
    const { data: responseData } = await UserApi.resetPassword(id)
    data = responseData
    showSuccessSnackbar('Password reset')
  }
  catch {
    showErrorSnackbar('Error resetting password')
  }
  return data
}

export async function setPassword(id, payload) {
  let data = null
  try {
    const { data: responseData } = await UserApi.setPassword(id, payload)
    data = responseData
    showSuccessSnackbar('Password set')
  }
  catch {
    showErrorSnackbar('Error setting password')
  }
  return data
}

export async function remove(id) {
  let data = null
  try {
    const { data: responseData } = await UserApi.remove(id)
    data = responseData
    showSuccessSnackbar('User deleted')
  }
  catch {
    showErrorSnackbar('Error deleting user')
  }
  return data
}
