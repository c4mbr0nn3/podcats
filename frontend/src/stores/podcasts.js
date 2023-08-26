import { defineStore } from "pinia";
import { PodcastService } from "@/services";
import { ref, unref } from "vue";

export const usePodcastsStore = defineStore("podcasts", () => {
  const podcasts = ref([]);

  async function fetchPodcasts() {
    podcasts.value = await PodcastService.getAll();
  }

  async function importPodcast(podcastUrl) {
    let payload = { podcastUrl: unref(podcastUrl) };
    await PodcastService.importPodcast(payload);
    fetchPodcasts();
  }

  function $reset() {
    podcasts.value = [];
  }

  return {
    podcasts,
    fetchPodcasts,
    importPodcast,
    $reset,
  };
});
