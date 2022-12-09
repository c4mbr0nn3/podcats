import axiosClient from "../axios";

export function getAllPodcasts() {
  return axiosClient.get("/podcasts");
}

export function getPodcastById(podcastId) {
  return axiosClient.get(`/podcasts/${podcastId}`);
}

export function getPodcastItemById(podcastId, itemId) {
  return axiosClient.get(`/podcasts/${podcastId}/item/${itemId}`);
}
export function importPodcast(payload) {
  return axiosClient.post("/podcasts/import", payload);
}
