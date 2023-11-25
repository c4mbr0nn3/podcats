import { defineStore } from 'pinia'
import { computed, ref, unref } from 'vue'
import { NotificationService } from '@/services'

export const useNotificationsStore = defineStore('notifications', () => {
  const notifications = ref([])

  const anyUnread = computed(() => {
    return notifications.value.length > 0
  })
  const unreadCount = computed(() => {
    return notifications.value.length
  })

  async function fetchNotifications() {
    notifications.value = await NotificationService.getAll()
  }

  async function setAllRead() {
    await NotificationService.setAllRead()
    fetchNotifications()
  }

  async function setRead(id) {
    await NotificationService.setRead(unref(id))
    fetchNotifications()
  }

  function $reset() {
    notifications.value = []
  }

  return {
    notifications,
    anyUnread,
    unreadCount,
    fetchNotifications,
    setAllRead,
    setRead,
    $reset,
  }
})
