<template>
  <v-app
    ><TheAppBarVue />
    <v-main>
      <v-container :class="getMainContainerClass"
        ><router-view></router-view
      ></v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import TheAppBarVue from "./components/global/TheAppBar.vue";
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const getMainContainerClass = computed(() => {
  return {
    "fill-height": route.name == "home",
  };
});
</script>
