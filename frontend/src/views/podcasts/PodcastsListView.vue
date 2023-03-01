<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card
        ><v-card-title>I miei podcasts</v-card-title>
        <v-card-text
          ><v-text-field
            v-model="podcastUrl"
            label="Import Podcast"
            append-inner-icon="mdi-plus"
            color="primary"
            @click:append-inner="importPodcast"
          ></v-text-field
          ><PodcastCard
            v-for="(podcast, index) in podcastData"
            :key="index + podcast.UpdatedAt"
            :data="podcast"
            @mark-all-played="triggerFetch()"
            @delete-podcast="triggerFetch()"
        /></v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { importPodcast, getAllPodcasts } from "@/api";
import PodcastCard from "@/components/PodcastCard.vue";
export default {
  components: {
    PodcastCard,
  },
  data: () => ({
    podcastData: null,
    podcastUrl: "",
  }),
  created() {
    this.triggerFetch();
  },
  methods: {
    async triggerFetch() {
      await this.fetchData();
    },
    async fetchData() {
      await getAllPodcasts().then(
        (response) => (this.podcastData = response.data)
      );
    },
    async importPodcast() {
      let payload = { podcastUrl: this.podcastUrl };
      await importPodcast(payload).then(() => {
        this.podcastUrl = null;
        this.fetchData();
      });
    },
  },
};
</script>
