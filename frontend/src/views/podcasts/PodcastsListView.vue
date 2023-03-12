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
      <v-dialog v-model="infoDialog" width="700">
        <v-card>
          <v-card-title>Podcast Info</v-card-title>
          <v-card-text>
            <div class="text-h6 font-italic text-primary">Title</div>
            <p-markdown :markdown="infoDialogData.Title" />
            <v-divider class="my-1"></v-divider>
            <div class="text-h6 font-italic text-primary">Author</div>
            <p-markdown :markdown="infoDialogData.Author" />
            <v-divider class="my-1"></v-divider>
            <div class="text-h6 font-italic text-primary">Summary</div>
            <p-markdown :markdown="infoDialogData.Summary" />
            <v-divider class="my-1"></v-divider>
            <div class="text-h6 font-italic text-primary">Show URL</div>
            <p-markdown :markdown="infoDialogData.ShowURL" />
            <v-divider class="my-1"></v-divider>
            <div class="text-h6 font-italic text-primary">Imported At</div>
            <p-markdown :markdown="formatDate(infoDialogData.CreatedAt)" />
            <v-divider class="my-1"></v-divider>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" block @click="infoDialog = false"
              >Close</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-col>
  </v-row>
</template>

<script>
import { importPodcast, getAllPodcasts } from "@/api";
import SinglePodcastCard from "@/components/SinglePodcastCard.vue";
import PMarkdown from "@/components/global/PMarkdown.vue";
import { formatDate } from "@/utils/date";

export default {
  components: {
    SinglePodcastCard,
    PMarkdown,
  },
  data: () => ({
    podcastData: null,
    podcastUrl: "",
    infoDialog: false,
    infoDialogData: null,
    formatDate,
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
    closeInfoDialog() {
      this.infoDialogData = null;
      this.infoDialog = false;
    },
  },
};
</script>
