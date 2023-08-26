<template>
  <v-menu :close-on-content-click="false">
    <template #activator="{ props }">
      <v-btn
        color="primary"
        size="small"
        icon="mdi-bell-outline"
        v-bind="props"
        class="mx-3"
      >
        <v-badge
          :model-value="!!notificationsStore.notifications.length"
          color="red"
          dot
        >
          <v-icon color="primary" />
        </v-badge>
      </v-btn>
    </template>
    <v-card width="450" max-height="600">
      <v-card-title class="d-flex">
        Notifications <v-spacer />
        <v-btn
          v-if="notificationsStore.anyUnread"
          variant="text"
          density="comfortable"
          icon="mdi-notification-clear-all"
          @click="setAllRead"
        >
        </v-btn>
      </v-card-title>
      <v-card-text>
        <div v-if="notificationsStore.anyUnread">
          <NotificationCard
            v-for="(item, i) in notificationsStore.notifications"
            :key="i"
            :notification="item"
          />
        </div>
        <div v-else>No new notifications</div>
      </v-card-text>
    </v-card>
  </v-menu>
</template>

<script setup>
import { useNotificationsStore } from "@/stores/notifications";
import NotificationCard from "@/components/app-bar/AppBarNotificationCard.vue";
import { onMounted } from "vue";
import { useIntervalFn } from "@vueuse/core";

const notificationsStore = useNotificationsStore();
const { fetchNotifications, setAllRead } = notificationsStore;

onMounted(() => {
  fetchNotifications();
  useIntervalFn(fetchNotifications, 1000 * 60);
});
</script>
