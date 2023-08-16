import { PodcastApi } from "@/api";
import { showSuccessSnackbar, showErrorSnackbar } from "@/utils/snackbar";

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
  let data = null;
  try {
    data = await PodcastApi.deleteById(podcastId);
    showSuccessSnackbar("Podcast deleted");
  } catch {
    showErrorSnackbar("Error deleting podcast");
  }
  return data;
}

export async function setPlayed(podcastId) {
  let data = null;
  try {
    const { data: responseData } = await PodcastApi.setPlayed(podcastId);
    data = responseData;
    showSuccessSnackbar("Podcast marked as played");
  } catch {
    showErrorSnackbar("Error marking podcast as played");
  }
  return data;
}

export async function importPodcast(payload) {
  let data = null;
  try {
    const { data: responseData } = await PodcastApi.importPodcast(payload);
    data = responseData;
    showSuccessSnackbar("Podcast imported");
  } catch {
    showErrorSnackbar("Error importing podcast");
  }
  return data;
}
