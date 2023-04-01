import { createRouter, createWebHashHistory } from "vue-router";
import routes from "./routes";
import { useAuthStore } from "@/stores/auth";

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});

// redirect to login page if not logged in and trying to access a restricted page
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();
  const isAuthenticated = authStore.isAuthenticated;
  const authRequired = to.meta.requiresAuth;

  if (authRequired && !isAuthenticated) {
    return next({ name: "login" });
  }

  next();
});

export default router;
