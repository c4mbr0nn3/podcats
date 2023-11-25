import { useSnackbarStore } from '@/stores/snackbar'

const defaultTimeout = 3000

function useSnackbar() {
  const snackbarStore = useSnackbarStore()
  const { showSuccessSnackbar, showErrorSnackbar } = snackbarStore
  return {
    showSuccessSnackbar,
    showErrorSnackbar,
  }
}

function showSuccessSnackbar(message, timeout = defaultTimeout) {
  const { showSuccessSnackbar: show } = useSnackbar()
  show(message, timeout)
}

function showErrorSnackbar(message, timeout = defaultTimeout) {
  const { showErrorSnackbar: show } = useSnackbar()
  show(message, timeout)
}

export { showSuccessSnackbar, showErrorSnackbar }
