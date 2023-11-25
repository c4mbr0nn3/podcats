<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const { logout } = authStore

const router = useRouter()

const dropdownMenu = [
  {
    title: 'Settings',
    icon: 'mdi-cog',
    onClick: () => {
      router.push({ name: 'settings-admin' })
    },
    onlyAdmin: true,
  },
  {
    title: 'Logout',
    icon: 'mdi-logout',
    onClick: logout,
  },
]

const dropdown = computed(() => {
  if (authStore.isAdmin)
    return dropdownMenu
  return dropdownMenu.filter(item => !item.onlyAdmin)
})

const username = computed(() => `@${authStore.user?.Username}`)
</script>

<template>
  <v-menu>
    <template #activator="{ props }">
      <v-avatar image="@/assets/nyan-cat.jpg" v-bind="props" class="mr-3" />
    </template>
    <v-card min-width="250">
      <v-card-title class="pb-0">
        {{ authStore.fullName }}
      </v-card-title>
      <v-card-subtitle>{{ username }}</v-card-subtitle>
      <v-list>
        <v-list-item
          v-for="(item, i) in dropdown"
          :key="i"
          @click="item.onClick"
        >
          <template #prepend>
            <v-icon :icon="item.icon" color="primary" />
          </template>
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-card>
  </v-menu>
</template>
