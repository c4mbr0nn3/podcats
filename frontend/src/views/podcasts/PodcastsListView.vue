<template>
  <v-row align="center" justify="center">
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
          ><v-card
            v-for="(podcast, index) in podcastData"
            :key="index"
            class="mt-3"
            :to="{ name: 'single-podcast', params: { id: podcast.ID } }"
          >
            <div class="d-flex flex-no-wrap justify-space-between">
              <v-col cols="10">
                <v-card-title class="text-truncate">{{
                  podcast.Title
                }}</v-card-title>
                <v-card-text
                  ><Markdown
                    class="overflow-hidden fade"
                    :source="podcast.Summary"
                  />
                  <v-chip label class="mt-2" color="primary" variant="outlined">
                    <v-icon start icon="mdi-counter"></v-icon>
                    {{ `${podcast.PlayedCount}/${podcast.EpisodesCount}` }}
                  </v-chip>
                </v-card-text>
              </v-col>
              <v-col cols="2" class="d-flex justify-end ml-2">
                <v-avatar class="ma-3" size="125" rounded="0">
                  <v-img :src="podcast.Image"></v-img>
                </v-avatar>
              </v-col></div></v-card
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

<style scoped>
/* https://css-tricks.com/line-clampin/ */
.fade {
  position: relative;
  height: 6em; /* exactly four lines with line height = 1.5*/
}
.fade:after {
  content: "";
  text-align: right;
  position: absolute;
  bottom: 0;
  right: 0;
  width: 70%;
  height: 1.5em;
  background: linear-gradient(
    to right,
    rgba(var(--v-theme-surface), 0),
    rgba(var(--v-theme-surface), 1) 50%
  );
}
</style>
