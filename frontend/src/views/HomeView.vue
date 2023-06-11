<template>
  <v-responsive class="text-center">
    <PodcatsLogoVue :width="256" :height="256" />
    <div class="text-body-2 font-weight-light mb-n1">Welcome to</div>
    <h1 class="text-h2 font-weight-bold">Podcats</h1>
  </v-responsive>
</template>

<script setup>
import PodcatsLogoVue from "@/components/svg/PodcatsLogo.vue";
</script>
