<template>
  <v-card class="mt-3">
    <div class="d-flex flex-no-wrap justify-space-between">
      <v-col cols="10">
        <v-card-title class="text-truncate">{{ podcast.Title }}</v-card-title>
        <v-card-text
          ><Markdown class="overflow-hidden fade" :source="podcast.Summary" />
          <div class="d-flex align-center mt-2">
            <v-chip
              label
              :color="
                podcast.PlayedCount === podcast.EpisodesCount
                  ? 'success'
                  : 'primary'
              "
              variant="outlined"
            >
              <v-icon start icon="mdi-counter"></v-icon>
              {{ `${podcast.PlayedCount}/${podcast.EpisodesCount}` }}
            </v-chip>
            <v-tooltip
              text="Mark all as played"
              location="bottom"
              :disabled="podcast.PlayedCount === podcast.EpisodesCount"
            >
              <template #activator="{ props }">
                <v-icon
                  v-if="podcast.PlayedCount !== podcast.EpisodesCount"
                  v-bind="props"
                  class="ml-3"
                  color="primary"
                  @click="markAllPlayed(podcast)"
                  >mdi-check-all
                </v-icon>
              </template>
            </v-tooltip>
            <v-tooltip text="Remove podcast" location="bottom">
              <template #activator="{ props }">
                <v-icon
                  v-bind="props"
                  class="ml-3"
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
        <PodcastAvatar
          :image="podcast.Image"
          :router-link="{ name: 'single-podcast', params: { id: podcast.ID } }"
        />
      </v-col></div
  ></v-card>
</template>

<script>
import Markdown from "vue3-markdown-it";
import PodcastAvatar from "@/components/global/PodcastAvatar.vue";
import { deletePodcastById, setPodcastPlayed } from "@/api";

export default {
  components: {
    Markdown,
    PodcastAvatar,
  },
  props: {
    data: { type: Object, default: () => {} },
  },
  emits: ["delete-podcast", "mark-all-played"],
  data() {
    return {
      podcast: this.data,
    };
  },
  updated() {
    this.podcast = this.data;
  },
  methods: {
    async deletePodcast(podcast) {
      await deletePodcastById(podcast.ID).then(() =>
        this.$emit("delete-podcast")
      );
    },
    async markAllPlayed(podcast) {
      await setPodcastPlayed(podcast.ID).then(() =>
        this.$emit("mark-all-played")
      );
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
