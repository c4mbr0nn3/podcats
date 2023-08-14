import axios from "axios";
import { useAuthStore } from "@/stores/auth";

const baseURL = import.meta.env.DEV
  ? "http://localhost:8000/api/v1"
  : "/api/v1";

const axiosOptions = {
  baseURL: baseURL,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
};

const axiosClient = axios.create(axiosOptions);

axiosClient.interceptors.request.use((config) => {
  const authStore = useAuthStore();
  if (authStore.isAuthenticated)
    config.headers.Authorization = `Bearer ${authStore.getToken}`;
  return config;
});

axiosClient.interceptors.response.use(
  (res) => res,
  (err) => {
    console.error("Api Error:", err);
    return Promise.reject(err);
  }
);

export default axiosClient;
