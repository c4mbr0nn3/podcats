import axiosClient from "../axios";

export function getAllPodcasts() {
  return axiosClient.get("/podcasts");
}

export function getPodcastById(podcastId) {
  return axiosClient.get(`/podcasts/${podcastId}`);
}

export function getPodcastItemsByPodcastId(podcastId, pageId) {
  return axiosClient.get(`/podcasts/${podcastId}/items?page=${pageId}`);
}

export function deletePodcastById(podcastId) {
  return axiosClient.delete(`/podcasts/${podcastId}/remove`);
}

export function setPodcastPlayed(podcastId) {
  return axiosClient.post(`/podcasts/${podcastId}/set-all-played`);
}

export function importPodcast(payload) {
  return axiosClient.post("/podcasts/import", payload);
}
