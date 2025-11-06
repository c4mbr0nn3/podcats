<script setup>
import { ref } from 'vue'
import PodcastInfoDialog from '@/components/PodcastInfoDialog.vue'
import SinglePodcastCard from '@/components/SinglePodcastCard.vue'
import { usePodcastsStore } from '@/stores/podcasts'

const infoDialog = ref(false)
const infoDialogData = ref(null)

const podcastsStore = usePodcastsStore()
const { fetchPodcasts } = podcastsStore

function showInfoDialog(event) {
  infoDialogData.value = event
  infoDialog.value = true
}

function refreshPodcastList() {
  fetchPodcasts()
}
</script>

<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card>
        <v-card-title>My Podcasts</v-card-title>
        <v-card-text>
          <SinglePodcastCard
            v-for="(podcast, index) in podcastsStore.podcasts"
            :key="index + podcast.UpdatedAt"
            :podcast="podcast"
            @mark-all-played="refreshPodcastList"
            @delete-podcast="refreshPodcastList"
            @show-info-dialog="showInfoDialog"
          />
        </v-card-text>
      </v-card>
      <PodcastInfoDialog
        v-if="infoDialogData"
        v-model:model-value="infoDialog"
        :info-dialog-data="infoDialogData"
      />
    </v-col>
  </v-row>
</template>
