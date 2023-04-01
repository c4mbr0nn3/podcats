import { defineStore } from "pinia";
import { AuthService } from "@/api";
import { useLocalStorage } from "@vueuse/core";
import { computed } from "vue";

export const useAuthStore = defineStore("auth", () => {
  const user = useLocalStorage("vueUseUser", {});
  const token = useLocalStorage("vueUseToken", "");

  const isAuthenticated = computed(() => !!token.value);
  const getUser = computed(() => user.value);
  const getToken = computed(() => token.value);

  async function login(payload) {
    await AuthService.login(payload).then((response) => {
      this.user = response.data.user;
      this.token = response.data.token;
      this.$router.push({ name: "home" });
    });
  }

  function logout() {
    this.user = null;
    this.token = null;
    this.$router.push({ name: "login" });
  }

  return {
    user,
    token,
    isAuthenticated,
    getUser,
    getToken,
    login,
    logout,
  };
});
