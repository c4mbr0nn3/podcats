<template>
  <v-app-bar flat>
    <v-row>
      <v-col cols="3">
        <global-search class="ml-6" />
      </v-col>
      <v-col cols="6" class="d-flex justify-center">
        <v-btn
          v-for="(item, index) in menuList"
          :key="index"
          :to="item.route"
          variant="text"
          color="primary"
          :prepend-icon="item.icon ? item.icon : null"
          class="mx-1"
          >{{ item.label }}
        </v-btn>
      </v-col>
      <v-col cols="3" class="d-flex justify-end">
        <v-btn
          v-if="authStore.isRoot"
          size="small"
          icon="mdi-cog"
          color="primary"
          class="mx-3"
          :to="{ name: 'admin' }"
        >
        </v-btn>
        <v-divider v-if="authStore.isRoot" vertical></v-divider>
        <v-btn
          color="primary"
          icon="mdi-logout"
          size="small"
          class="mx-3"
          @click="logout"
        ></v-btn>
      </v-col>
    </v-row>
  </v-app-bar>
</template>

<script setup>
import { onMounted } from "vue";
import GlobalSearch from "./GlobalSearch.vue";
import { useAuthStore } from "@/stores/auth";
import { usePodcastsStore } from "@/stores/podcasts";

const menuList = [
  { label: "Home", route: { name: "home" } },
  { label: "Podcasts", route: { name: "podcasts" } },
  { label: "Latest", route: { name: "latest-podcasts" } },
  {
    label: "Favs",
    icon: "mdi-heart",
    route: { name: "favourites-podcasts" },
  },
];
const authStore = useAuthStore();
const { logout } = authStore;

const podcastsStore = usePodcastsStore();
const { fetchPodcasts } = podcastsStore;

onMounted(() => {
  fetchPodcasts();
});
</script>
