<template>
  <v-chip
    label
    :color="
      podcast.PlayedCount === podcast.EpisodesCount ? 'success' : 'primary'
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
</template>

<script>
import { deletePodcastById, setPodcastPlayed } from "@/api";

export default {
  props: {
    podcast: { type: Object, default: () => {} },
  },
  emits: ["delete-podcast", "mark-all-played"],
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