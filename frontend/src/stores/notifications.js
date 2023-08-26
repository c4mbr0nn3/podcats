import { defineStore } from "pinia";
import { NotificationService } from "@/services";
import { ref, unref, computed } from "vue";

export const useNotificationsStore = defineStore("notifications", () => {
  const notifications = ref([]);

  const anyUnread = computed(() => {
    return notifications.value.length > 0;
  });

  async function fetchNotifications() {
    notifications.value = await NotificationService.getAll();
  }

  async function setAllRead() {
    await NotificationService.setAllRead();
    fetchNotifications();
  }

  async function setRead(id) {
    await NotificationService.setRead(unref(id));
    fetchNotifications();
  }

  return {
    notifications,
    anyUnread,
    fetchNotifications,
    setAllRead,
    setRead,
  };
});
