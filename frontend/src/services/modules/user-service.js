import { UserApi } from "@/api";

export async function getAll() {
  const { data } = await UserApi.getAll();
  return data;
}

export async function create(user) {
  const { data } = await UserApi.create(user);
  return data;
}

export async function update(id, user) {
  const { data } = await UserApi.update(id, user);
  return data;
}

export async function resetPassword(id) {
  const { data } = await UserApi.resetPassword(id);
  return data;
}

export async function setPassword(id, payload) {
  const { data } = await UserApi.setPassword(id, payload);
  return data;
}

export async function remove(id) {
  const { data } = await UserApi.remove(id);
  return data;
}
