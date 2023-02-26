import axiosClient from "../axios";

export function getAllPodcasts() {
  return axiosClient.get("/podcasts");
}

export function getPodcastById(podcastId) {
  return axiosClient.get(`/podcasts/${podcastId}`);
}

export function deletePodcastById(podcastId) {
  return axiosClient.delete(`/podcasts/${podcastId}/remove`);
}

export function setPodcastPlayed(podcastId) {
  return axiosClient.post(`/podcasts/${podcastId}/set-all-played`);
}

export function getLatestPodcastItems(pageId) {
  return axiosClient.get(`/podcasts/latest-items?page=${pageId}`);
}

export function getFavouritesPodcastItems(pageId) {
  return axiosClient.get(`/podcasts/favourites-items?page=${pageId}`);
}

export function getPodcastItemsByPodcastId(podcastId, pageId) {
  return axiosClient.get(`/podcasts/${podcastId}/items?page=${pageId}`);
}

export function getPodcastItemById(podcastId, itemId) {
  return axiosClient.get(`/podcasts/${podcastId}/items/${itemId}`);
}

export function updatePodcastItemPlayedTime(podcastId, itemId, timePlayed) {
  return axiosClient.post(
    `/podcasts/${podcastId}/items/${itemId}/update-played-time?time=${timePlayed}`
  );
}

export function setPodcastItemCompleted(podcastId, itemId) {
  return axiosClient.post(
    `/podcasts/${podcastId}/items/${itemId}/set-completed`
  );
}

export function switchPodcastItemPlayedStatus(podcastId, itemId) {
  return axiosClient.post(
    `/podcasts/${podcastId}/items/${itemId}/switch-status`
  );
}

export function switchPodcastItemFavouriteStatus(podcastId, itemId) {
  return axiosClient.post(
    `/podcasts/${podcastId}/items/${itemId}/switch-fav-status`
  );
}

export function importPodcast(payload) {
  return axiosClient.post("/podcasts/import", payload);
}
