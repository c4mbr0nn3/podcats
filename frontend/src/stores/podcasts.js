import { defineStore } from "pinia";
import { getAllPodcasts } from "../api/services/podcastsService";

export const usePodcastsStore = defineStore("podcasts", {
  state: () => ({
    podcasts: [],
  }),
  getters: {
    getPodcasts: (state) => state.podcasts,
  },
  actions: {
    async fetchPodcasts() {
      await getAllPodcasts().then((response) => {
        this.setPodcasts(response.data);
      });
    },
    setPodcasts(podcasts) {
      this.podcasts = podcasts;
    },
  },
});
