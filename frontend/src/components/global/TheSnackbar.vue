<script setup>
import { storeToRefs } from 'pinia'
import { useSnackbarStore } from '@/stores/snackbar'

const snackbarStore = useSnackbarStore()
const { show, message, color, timeout } = storeToRefs(snackbarStore)
</script>

<template>
  <v-snackbar v-model="show" :timeout="timeout" :color="color" variant="tonal">
    {{ message }}
    <template #actions>
      <v-btn
        :color="color"
        variant="text"
        icon="mdi-close"
        size="small"
        @click="show = false"
      />
    </template>
  </v-snackbar>
</template>
