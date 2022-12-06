<template>
  <v-container class="fill-height">
    <v-responsive class="d-flex align-center text-center fill-height">
      <vuetify-logo-vue />
      <div class="text-body-2 font-weight-light mb-n1">Welcome to</div>

      <h1 class="text-h2 font-weight-bold">Vuetify</h1>

      <div class="py-14" />

      <v-row class="d-flex align-center justify-center">
        <v-col cols="auto">
          <v-btn
            color="primary"
            :to="{ name: 'about' }"
            min-width="228"
            size="x-large"
            variant="flat"
          >
            <v-icon icon="mdi-information-outline" size="large" start />
            About Page
          </v-btn>
        </v-col>
      </v-row>
    </v-responsive>
  </v-container>
</template>

<script setup>
import VuetifyLogoVue from "./svg/VuetifyLogo.vue";
</script>
