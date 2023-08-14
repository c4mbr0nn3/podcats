import axiosClient from "../axios";

export function login(payload) {
  return axiosClient.post("/auth/login", payload);
}
