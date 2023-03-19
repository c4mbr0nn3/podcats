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
      <v-col cols="3"> </v-col>
    </v-row>
  </v-app-bar>
</template>

<script setup>
import GlobalSearch from "./GlobalSearch.vue";

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
</script>
