import { defineStore } from 'pinia'
import { ref, unref } from 'vue'
import { PodcastService } from '@/services'

export const usePodcastsStore = defineStore('podcasts', () => {
  const podcasts = ref([])

  async function fetchPodcasts() {
    podcasts.value = await PodcastService.getAll()
  }

  async function importPodcast(podcastUrl) {
    const payload = { podcastUrl: unref(podcastUrl) }
    await PodcastService.importPodcast(payload)
    fetchPodcasts()
  }

  function $reset() {
    podcasts.value = []
  }

  return {
    podcasts,
    fetchPodcasts,
    importPodcast,
    $reset,
  }
})
