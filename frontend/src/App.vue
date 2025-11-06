<script setup>
import { useTitle } from '@vueuse/core'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import TheAppBar from '@/components/global/TheAppBar.vue'
import TheSnackbar from '@/components/global/TheSnackbar.vue'
import { useNotificationsStore } from './stores/notifications'

const route = useRoute()
const notificationsStore = useNotificationsStore()

const getMainContainerClass = computed(() => {
  return {
    'fill-height': route.meta.requiresFillHeight,
  }
})

const title = computed(() => {
  const title = route.meta.name ? `PodCats | ${route.meta.name}` : 'PodCats'

  if (notificationsStore.anyUnread)
    return `(${notificationsStore.unreadCount}) ${title}`

  return title
})

useTitle(title)
</script>

<template>
  <v-app>
    <TheAppBar v-if="route.meta.requiresNavbar" />
    <v-main>
      <v-container :class="getMainContainerClass">
        <router-view />
      </v-container>
    </v-main>
    <TheSnackbar />
  </v-app>
</template>
