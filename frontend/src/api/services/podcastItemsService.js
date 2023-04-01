import axiosClient from "../axios";

export function getLatestPodcastItems(pageId) {
  return axiosClient.get(`/podcast-items/latest?page=${pageId}`);
}

export function getFavouritesPodcastItems(pageId) {
  return axiosClient.get(`/podcast-items/favourites?page=${pageId}`);
}

export function getPodcastItemById(itemId) {
  return axiosClient.get(`/podcast-items/${itemId}`);
}

export function updatePodcastItemPlayedTime(itemId, timePlayed) {
  return axiosClient.post(
    `/podcast-items/${itemId}/update-played-time?time=${timePlayed}`
  );
}

export function setPodcastItemCompleted(itemId) {
  return axiosClient.post(`/podcast-items/${itemId}/set-completed`);
}

export function switchPodcastItemPlayedStatus(itemId) {
  return axiosClient.post(`/podcast-items/${itemId}/switch-played-status`);
}

export function switchPodcastItemFavouriteStatus(itemId) {
  return axiosClient.post(`/podcast-items/${itemId}/switch-fav-status`);
}
