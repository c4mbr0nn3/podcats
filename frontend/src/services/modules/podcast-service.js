import { PodcastApi } from "@/api";

export async function getAll() {
  const { data } = await PodcastApi.getAll();
  return data;
}

export async function getById(podcastId) {
  const { data } = await PodcastApi.getById(podcastId);
  return data;
}

export async function getItemsById(podcastId, pageId) {
  const { data } = await PodcastApi.getItemsById(podcastId, pageId);
  return data;
}

export async function deleteById(podcastId) {
  const { data } = await PodcastApi.deleteById(podcastId);
  return data;
}

export async function setPlayed(podcastId) {
  const { data } = await PodcastApi.setPlayed(podcastId);
  return data;
}

export async function importPodcast(payload) {
  const { data } = await PodcastApi.importPodcast(payload);
  return data;
}
