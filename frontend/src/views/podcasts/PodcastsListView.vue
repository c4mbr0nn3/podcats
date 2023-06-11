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
              <v-icon color="primary" @click="triggerImport">
                mdi-plus
              </v-icon></template
            ></v-text-field
          ><single-podcast-card
            v-for="(podcast, index) in podcastsStore.getPodcasts"
            :key="index + podcast.UpdatedAt"
            :podcast="podcast"
            @mark-all-played="refreshPodcastList"
            @delete-podcast="refreshPodcastList"
            @show-info-dialog="showInfoDialog"
          />
        </v-card-text>
      </v-card>
      <podcast-info-dialog
        v-if="infoDialogData"
        v-model:model-value="infoDialog"
        :info-dialog-data="infoDialogData"
      />
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from "vue";
import { PodcastService } from "@/api";
import SinglePodcastCard from "@/components/SinglePodcastCard.vue";
import PodcastInfoDialog from "@/components/PodcastInfoDialog.vue";
import { usePodcastsStore } from "@/stores/podcasts";

const podcastUrl = ref("");
const infoDialog = ref(false);
const infoDialogData = ref(null);

const podcastsStore = usePodcastsStore();
const { fetchPodcasts } = podcastsStore;

function showInfoDialog(event) {
  infoDialogData.value = event;
  infoDialog.value = true;
}

function triggerImport() {
  let payload = { podcastUrl: podcastUrl.value };
  PodcastService.importPodcast(payload).then(() => {
    podcastUrl.value = null;
    refreshPodcastList();
  });
}

function refreshPodcastList() {
  fetchPodcasts();
}
</script>
