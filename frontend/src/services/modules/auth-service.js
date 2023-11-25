import { AuthApi } from '@/api'

export async function login(payload) {
  const { data } = await AuthApi.login(payload)
  return data
}
