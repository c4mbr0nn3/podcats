<template>
  <v-row align="center" justify="center">
    <v-col cols="8">
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
          ><v-card
            v-for="(podcast, index) in podcastData"
            :key="index"
            class="mt-3"
            :to="{ name: 'single-podcast', params: { id: podcast.ID } }"
          >
            <v-card-title> {{ podcast.Title }} </v-card-title>
            <v-card-text>
              <Markdown :source="podcast.Summary" />
            </v-card-text> </v-card
        ></v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { importPodcast, getAllPodcasts } from "@/api";
import Markdown from "vue3-markdown-it";
export default {
  components: {
    Markdown,
  },
  data: () => ({
    podcastData: null,
    podcastUrl: "",
  }),
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      await getAllPodcasts().then(
        (response) => (this.podcastData = response.data.data)
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
