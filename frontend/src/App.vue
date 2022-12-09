<template>
  <v-app
    ><the-app-bar-vue></the-app-bar-vue>
    <v-main class="bg-grey-lighten-3">
      <v-container class="fill-height"><router-view></router-view></v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import TheAppBarVue from "./components/global/TheAppBar.vue";
</script>
