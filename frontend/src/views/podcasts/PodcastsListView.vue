<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card
        ><v-card-title>My Podcasts</v-card-title>
        <v-card-text
          ><v-text-field
            v-model="podcastUrl"
            label="Import Podcast"
            color="primary"
            variant="underlined"
          >
            <template #append-inner>
              <v-icon color="primary" @click="importPodcast">
                mdi-plus
              </v-icon></template
            ></v-text-field
          ><single-podcast-card
            v-for="(podcast, index) in podcastData"
            :key="index + podcast.UpdatedAt"
            :podcast="podcast"
            @mark-all-played="triggerFetch"
            @delete-podcast="triggerFetch"
            @show-info-dialog="showInfoDialog"
          />
        </v-card-text>
      </v-card>
      <podcast-info-dialog
        v-if="infoDialogData"
        v-model="infoDialog"
        :info-dialog-data="infoDialogData"
        @close-info-dialog="infoDialog = false"
      />
    </v-col>
  </v-row>
</template>

<script>
import { importPodcast, getAllPodcasts } from "@/api";
import SinglePodcastCard from "@/components/SinglePodcastCard.vue";
import PodcastInfoDialog from "@/components/PodcastInfoDialog.vue";

export default {
  components: {
    SinglePodcastCard,
    PodcastInfoDialog,
  },
  data: () => ({
    podcastData: null,
    podcastUrl: "",
    infoDialog: false,
    infoDialogData: null,
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
    showInfoDialog(event) {
      this.infoDialogData = event;
      this.infoDialog = true;
    },
  },
};
</script>
