import axiosClient from "../axios";

export function getAllPodcasts() {
  return axiosClient.get("/podcasts");
}

export function getLatestPodcastItems(pageId) {
  return axiosClient.get(`/podcasts/latest-items?page=${pageId}`);
}

export function getPodcastById(podcastId) {
  return axiosClient.get(`/podcasts/${podcastId}`);
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

export function importPodcast(payload) {
  return axiosClient.post("/podcasts/import", payload);
}
