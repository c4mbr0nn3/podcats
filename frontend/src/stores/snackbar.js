import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSnackbarStore = defineStore('snackbar', () => {
  const show = ref(false)
  const message = ref('')
  const color = ref('')
  const timeout = ref(0)

  function showSnackbar(msg, clr, tmo) {
    show.value = true
    message.value = msg
    color.value = clr
    timeout.value = tmo
  }

  function showSuccessSnackbar(message, timeout) {
    showSnackbar(message, 'success', timeout)
  }

  function showErrorSnackbar(message, timeout) {
    showSnackbar(message, 'error', timeout)
  }

  return {
    show,
    message,
    color,
    timeout,
    showSuccessSnackbar,
    showErrorSnackbar,
  }
})
