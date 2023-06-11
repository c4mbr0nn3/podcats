import { defineStore } from "pinia";
import { AuthService, UsersService } from "@/api";
import { useLocalStorage } from "@vueuse/core";
import { computed } from "vue";

export const useAuthStore = defineStore("auth", () => {
  const user = useLocalStorage("vueUseUser", {});
  const token = useLocalStorage("vueUseToken", "");

  const isAuthenticated = computed(() => !!token.value);
  const needPasswordChange = computed(() => user.value.NeedPasswordChange);
  const isRoot = computed(() => user.value.ID === 1);
  const getUser = computed(() => user.value);
  const getToken = computed(() => token.value);

  async function login(payload) {
    await AuthService.login(payload).then((response) => {
      user.value = response.data.user;
      token.value = response.data.token;
      if (needPasswordChange.value)
        this.$router.push({ name: "change-password" });
      else this.$router.push({ name: "home" });
    });
  }

  async function setPassword(payload) {
    await UsersService.setPassword(user.value.ID, payload).then(() => {
      user.value.NeedPasswordChange = false;
      this.$router.push({ name: "home" });
    });
  }

  async function logout() {
    await this.$router.push({ name: "login" });
    user.value = null;
    token.value = null;
  }

  return {
    user,
    token,
    isAuthenticated,
    needPasswordChange,
    isRoot,
    getUser,
    getToken,
    login,
    logout,
    setPassword,
  };
});
