import { defineStore } from "pinia";
import { PodcastService } from "@/services";
import { ref } from "vue";

export const usePodcastsStore = defineStore("podcasts", () => {
  const podcasts = ref([]);

  async function fetchPodcasts() {
    podcasts.value = await PodcastService.getAll();
  }

  return {
    podcasts,
    fetchPodcasts,
  };
});
