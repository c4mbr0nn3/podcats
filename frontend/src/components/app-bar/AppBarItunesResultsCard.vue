<script setup>
import { ref } from 'vue'
import { usePodcastsStore } from '@/stores/podcasts'

const props = defineProps({
  podcast: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['imported'])

const isImporting = ref(false)

const podcastsStore = usePodcastsStore()
const { importPodcast } = podcastsStore

async function onClick() {
  isImporting.value = true
  await importPodcast(props.podcast.feedUrl)
  isImporting.value = false
  emit('imported')
}
</script>

<template>
  <v-card variant="tonal" :loading="isImporting" @click="onClick">
    <template #loader="{ isActive }">
      <v-progress-linear
        :active="isActive"
        color="primary"
        height="4"
        indeterminate
      />
    </template>
    <v-row>
      <v-col cols="10">
        <v-card-title class="pb-0 text-truncate">
          {{ podcast.collectionName }}
        </v-card-title>
        <v-card-subtitle class="text-subtitle-2 pb-2">
          {{ podcast.artistName }}
        </v-card-subtitle>
      </v-col>
      <v-col cols="2" class="d-flex align-center">
        <v-avatar rounded>
          <v-img cover :src="podcast.artworkUrl100" />
        </v-avatar>
      </v-col>
    </v-row>
  </v-card>
</template>
