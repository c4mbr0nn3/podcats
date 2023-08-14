import { PodcastItemApi } from "@/api";

export async function getLatest(pageId) {
  const { data } = await PodcastItemApi.getLatest(pageId);
  return data;
}

export async function getFavourites(pageId) {
  const { data } = await PodcastItemApi.getFavourites(pageId);
  return data;
}

export async function getById(itemId) {
  const { data } = await PodcastItemApi.getById(itemId);
  return data;
}

export async function updatePodcastItemPlayedTime(itemId, timePlayed) {
  const { data } = await PodcastItemApi.updatePodcastItemPlayedTime(
    itemId,
    timePlayed
  );
  return data;
}

export async function setPodcastItemCompleted(itemId) {
  const { data } = await PodcastItemApi.setPodcastItemCompleted(itemId);
  return data;
}

export async function switchStatus(itemId) {
  const { data } = await PodcastItemApi.switchStatus(itemId);
  return data;
}

export async function switchFavourite(itemId) {
  const { data } = await PodcastItemApi.switchFavourite(itemId);
  return data;
}
