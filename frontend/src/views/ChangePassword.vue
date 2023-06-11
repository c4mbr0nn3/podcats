<template>
  <v-responsive>
    <v-card width="500px" class="mx-auto">
      <v-card-title>Please change your password 🔑</v-card-title>
      <v-card-subtitle>
        You need to change your password before you can continue
      </v-card-subtitle>
      <v-form
        ref="change_password_form"
        v-model="formValidation"
        @submit.prevent="submit"
      >
        <v-card-text>
          <v-text-field
            v-model="password"
            label="Password"
            variant="underlined"
            :rules="[isRequiredRule, validPasswordRule]"
            :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            :type="showPassword ? 'text' : 'password'"
            @click:append-inner="() => (showPassword = !showPassword)"
            @click.right.prevent
            @copy.prevent
            @paste.prevent
          />
          <v-text-field
            v-model="confirmPassword"
            label="Confirm Password"
            variant="underlined"
            :rules="[
              isRequiredRule,
              confirmPasswordRule(password, confirmPassword),
            ]"
            :append-inner-icon="showConfirmPassword ? 'mdi-eye-off' : 'mdi-eye'"
            :type="showConfirmPassword ? 'text' : 'password'"
            @click:append-inner="
              () => (showConfirmPassword = !showConfirmPassword)
            "
            @click.right.prevent
            @copy.prevent
            @paste.prevent
          />
        </v-card-text>
        <v-card-actions>
          <v-btn
            block
            type="submit"
            color="primary"
            :loading="isLoading"
            :disabled="!formValidation"
          >
            Continue
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </v-responsive>
</template>

<script setup>
import { ref } from "vue";
import {
  confirmPasswordRule,
  isRequiredRule,
  validPasswordRule,
} from "@/utils/validation";
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router";
import { onMounted } from "vue";

const formValidation = ref(false);
const showPassword = ref(false);
const showConfirmPassword = ref(false);
const password = ref("");
const confirmPassword = ref("");
const isLoading = ref(false);

const authStore = useAuthStore();
const { setPassword } = authStore;
const router = useRouter();

onMounted(() => {
  if (!authStore.needPasswordChange) {
    router.push({ name: "home" });
  }
});

async function submit() {
  isLoading.value = true;

  const payload = {
    password: password.value,
  };

  await setPassword(payload).finally(() => {
    isLoading.value = false;
  });
}
</script>
