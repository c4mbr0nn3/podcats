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
                  <div class="d-flex align-center mt-2">
                    <v-chip label color="primary" variant="outlined">
                      <v-icon start icon="mdi-counter"></v-icon>
                      {{ `${podcast.PlayedCount}/${podcast.EpisodesCount}` }}
                    </v-chip>
                    <v-tooltip text="Remove podcast" location="bottom">
                      <template #activator="{ props }">
                        <v-icon
                          v-bind="props"
                          class="mx-3"
                          color="red"
                          @click="deletePodcast(podcast)"
                          >mdi-delete-outline
                        </v-icon>
                      </template>
                    </v-tooltip>
                  </div>
                </v-card-text>
              </v-col>
              <v-col cols="2" class="d-flex justify-end ml-2">
                <v-hover v-slot="{ isHovering, props }"
                  ><router-link
                    :to="{ name: 'single-podcast', params: { id: podcast.ID } }"
                  >
                    <v-avatar
                      class="ma-3"
                      size="125"
                      rounded="0"
                      :class="{ 'on-hover': isHovering }"
                      v-bind="props"
                    >
                      <v-img
                        :src="podcast.Image ? podcast.Image : missingImage"
                      >
                        <template #placeholder>
                          <div
                            class="d-flex align-center justify-center fill-height"
                          >
                            <v-progress-circular
                              indeterminate
                              color="primary"
                            ></v-progress-circular></div
                        ></template>
                        <div
                          class="fill-height d-flex justify-center align-center"
                        >
                          <v-icon
                            color="transparent"
                            size="x-large"
                            :class="{ 'show-btns': isHovering }"
                            >mdi-play-circle</v-icon
                          >
                        </div></v-img
                      >
                    </v-avatar></router-link
                  >
                </v-hover>
              </v-col>
            </div></v-card
          ></v-card-text
        >
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import missingImage from "@/assets/missing_image.png";
import { importPodcast, getAllPodcasts, deletePodcastById } from "@/api";
import Markdown from "vue3-markdown-it";
export default {
  components: {
    Markdown,
  },
  data: () => ({
    podcastData: null,
    podcastUrl: "",
    missingImage: missingImage,
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
    async deletePodcast(podcast) {
      await deletePodcastById(podcast.ID).then(() => this.fetchData());
    },
  },
};
</script>

<style scoped>
.v-avatar {
  transition: opacity 0.4s ease-in-out;
}
.v-avatar:not(.on-hover) {
  opacity: 0.6;
}

.show-btns {
  color: rgba(var(--v-theme-primary), 1) !important;
  opacity: 1 !important;
}

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
