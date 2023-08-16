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

const route = useRoute();

const getMainContainerClass = computed(() => {
  return {
    "fill-height": route.meta.requiresFillHeight,
  };
});
</script>
