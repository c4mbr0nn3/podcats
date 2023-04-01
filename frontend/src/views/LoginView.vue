<template>
  <v-responsive>
    <v-card width="500px" class="mx-auto">
      <v-card-title>Hey there! 👋</v-card-title>
      <v-card-subtitle>Please login with your credentials</v-card-subtitle>
      <v-form @submit.prevent="submit">
        <v-card-text>
          <v-text-field
            v-model="username"
            label="Username"
            type="text"
            variant="underlined"
          />
          <v-text-field
            v-model="password"
            label="Password"
            type="password"
            variant="underlined"
          />
        </v-card-text>
        <v-card-actions
          ><v-btn block type="submit" color="primary" :loading="isLoginOngoing"
            >Login</v-btn
          ></v-card-actions
        >
      </v-form>
    </v-card>
  </v-responsive>
</template>

<script setup>
import { ref } from "vue";
import { useAuthStore } from "@/stores/auth";

const username = ref("");
const password = ref("");
const isLoginOngoing = ref(false);

const authStore = useAuthStore();
const { login } = authStore;

async function submit() {
  isLoginOngoing.value = true;
  const payload = {
    username: username.value,
    password: password.value,
  };
  await login(payload);
}
</script>
