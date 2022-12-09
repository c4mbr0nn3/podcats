<template>
  <v-app
    ><the-app-bar-vue></the-app-bar-vue>
    <v-main>
      <v-container class="fill-height"><router-view></router-view></v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import TheAppBarVue from "./components/global/TheAppBar.vue";
</script>
