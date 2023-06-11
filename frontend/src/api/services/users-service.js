import axiosClient from "../axios";

export function getAll() {
  return axiosClient.get("/users/");
}

export function create(user) {
  return axiosClient.post("/users/create", user);
}

export function update(id, user) {
  return axiosClient.post(`/users/${id}/update`, user);
}

export function resetPassword(id) {
  return axiosClient.get(`/users/${id}/reset-password`);
}

export function setPassword(id, payload) {
  return axiosClient.post(`/users/${id}/set-password`, payload);
}

export function remove(id) {
  return axiosClient.delete(`/users/${id}/delete`);
}
