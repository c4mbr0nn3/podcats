import axios from "axios";
import { useAuthStore } from "@/stores/auth";

const axiosOptions = {
  baseURL: `api/v1`,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
};

const axiosClient = axios.create(axiosOptions);

axiosClient.interceptors.request.use((config) => {
  const authStore = useAuthStore();
  if (authStore.isAuthenticated) config.headers.Authorization = `Bearer ${authStore.getToken}`;
  return config;
});


export default axiosClient;
