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
  <v-chip label color="primary" variant="outlined" class="ml-3">
    <v-icon start icon="mdi-calendar"></v-icon>
    {{ formatDateToIso(podcast.LatestEpisode) }}
  </v-chip>
  <v-tooltip text="Show podcast info" location="bottom">
    <template #activator="{ props }">
      <v-icon
        v-bind="props"
        class="ml-3"
        color="primary"
        @click="$emit('show-info-dialog', podcast)"
        >mdi-information-variant
      </v-icon>
    </template>
  </v-tooltip>
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
import { PodcastService } from "@/api";
import { formatDateToIso } from "@/utils/date";

export default {
  props: {
    podcast: { type: Object, default: () => {} },
  },
  emits: ["delete-podcast", "mark-all-played", "show-info-dialog"],
  data() {
    return {
      formatDateToIso,
    };
  },
  methods: {
    async deletePodcast(podcast) {
      await PodcastService.deletePodcastById(podcast.ID).then(() =>
        this.$emit("delete-podcast")
      );
    },
    async markAllPlayed(podcast) {
      await PodcastService.setPodcastPlayed(podcast.ID).then(() =>
        this.$emit("mark-all-played")
      );
    },
  },
};
</script>
