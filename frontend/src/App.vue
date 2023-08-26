<template>
  <v-app>
    <TheAppBar v-if="route.meta.requiresNavbar" />
    <v-main>
      <v-container :class="getMainContainerClass">
        <router-view></router-view>
      </v-container>
    </v-main>
    <TheSnackbar />
  </v-app>
</template>

<script setup>
import TheAppBar from "@/components/global/TheAppBar.vue";
import TheSnackbar from "@/components/global/TheSnackbar.vue";
import { computed } from "vue";
import { useRoute } from "vue-router";
import { useTitle } from "@vueuse/core";
import { useNotificationsStore } from "./stores/notifications";

const route = useRoute();
const notificationsStore = useNotificationsStore();

const getMainContainerClass = computed(() => {
  return {
    "fill-height": route.meta.requiresFillHeight,
  };
});

const title = computed(() => {
  if (notificationsStore.anyUnread) {
    return `(${notificationsStore.unreadCount}) PodCats`;
  }

  return "PodCats";
});

useTitle(title);
</script>
