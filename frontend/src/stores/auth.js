import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core'
import { computed } from 'vue'
import { usePodcastsStore } from './podcasts'
import { useNotificationsStore } from './notifications'
import { AuthService, UserService } from '@/services'

export const useAuthStore = defineStore('auth', () => {
  const user = useLocalStorage('vueUseUser', {})
  const token = useLocalStorage('vueUseToken', '')

  const isAuthenticated = computed(() => !!token.value)
  const needPasswordChange = computed(() => user.value.NeedPasswordChange)
  const isRoot = computed(() => user.value.ID === 1)
  const isAdmin = computed(() => user.value.IsAdmin)
  const fullName = computed(() => `${user.value.Name} ${user.value.Surname}`)

  async function login(payload) {
    const data = await AuthService.login(payload)
    user.value = data.user
    token.value = data.token
    if (needPasswordChange.value)
    // eslint-disable-next-line no-invalid-this
      this.$router.push({ name: 'change-password' })
    // eslint-disable-next-line no-invalid-this
    else this.$router.push({ name: 'home' })
  }

  async function setPassword(payload) {
    // TODO: check if still working
    await UserService.setPassword(user.value.ID, payload).then(() => {
      user.value.NeedPasswordChange = false
      // eslint-disable-next-line no-invalid-this
      this.$router.push({ name: 'home' })
    })
  }

  async function logout() {
    // eslint-disable-next-line no-invalid-this
    await this.$router.push({ name: 'login' })
    $reset()
    usePodcastsStore().$reset()
    useNotificationsStore().$reset()
  }

  function $reset() {
    user.value = null
    token.value = null
  }

  return {
    user,
    token,
    isAuthenticated,
    needPasswordChange,
    isRoot,
    isAdmin,
    fullName,
    login,
    logout,
    setPassword,
    $reset,
  }
})
