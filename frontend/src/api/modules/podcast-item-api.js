import axiosClient from "../axios";

export function getLatest(pageId) {
  return axiosClient.get(`/podcast-items/latest?page=${pageId}`);
}

export function getFavourites(pageId) {
  return axiosClient.get(`/podcast-items/favourites?page=${pageId}`);
}

export function getById(itemId) {
  return axiosClient.get(`/podcast-items/${itemId}`);
}

export function updatePlayedTime(itemId, timePlayed) {
  return axiosClient.post(
    `/podcast-items/${itemId}/update-played-time?time=${timePlayed}`
  );
}

export function setCompleted(itemId) {
  return axiosClient.post(`/podcast-items/${itemId}/set-completed`);
}

export function switchStatus(itemId) {
  return axiosClient.post(`/podcast-items/${itemId}/switch-played-status`);
}

export function switchFavourite(itemId) {
  return axiosClient.post(`/podcast-items/${itemId}/switch-fav-status`);
}
