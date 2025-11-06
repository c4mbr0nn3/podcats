<script>
import { PodcastItemService } from '@/services'
import { usePodcastsStore } from '@/stores/podcasts'
import { formatDate } from '@/utils/date'

export default {
  props: {
    podcastItem: { type: Object, default: () => {} },
  },
  emits: ['change-played-status', 'change-fav-status'],
  setup() {
    const podcastsStore = usePodcastsStore()
    const { fetchPodcasts } = podcastsStore
    return {
      formatDate,
      fetchPodcasts,
    }
  },
  methods: {
    async changePlayedStatus(podcastItem) {
      await PodcastItemService.switchStatus(
        podcastItem.ID,
        podcastItem.IsPlayed,
      )
      this.$emit('change-played-status', podcastItem.ID)
      this.fetchPodcasts()
    },
    async changeFavStatus(podcastItem) {
      const data = await PodcastItemService.switchFavourite(
        podcastItem.ID,
        !!podcastItem.BookmarkDate,
      )
      this.$emit('change-fav-status', {
        id: podcastItem.ID,
        bookmarkDate: data,
      })
      this.fetchPodcasts()
    },
  },
}
</script>

<template>
  <div class="d-flex align-center mt-2">
    <v-chip label color="primary" variant="outlined">
      <v-icon start icon="mdi-calendar" />
      {{ formatDate(podcastItem.PublicationDate) }}
    </v-chip>
    <v-tooltip
      :text="podcastItem.IsPlayed ? 'Mark as not-played' : 'Mark as played'"
      location="bottom"
    >
      <template #activator="{ props }">
        <v-icon
          v-bind="props"
          class="mx-3"
          :color="podcastItem.IsPlayed ? 'success' : 'primary'"
          @click="changePlayedStatus(podcastItem)"
        >
          {{
            podcastItem.IsPlayed
              ? "mdi-checkbox-marked-circle-outline"
              : "mdi-checkbox-blank-circle-outline"
          }}
        </v-icon>
      </template>
    </v-tooltip>
    <v-tooltip
      :text="podcastItem.BookmarkDate ? 'Remove from favs' : 'Add to favs'"
      location="bottom"
    >
      <template #activator="{ props }">
        <v-icon
          v-bind="props"
          class="mr-3"
          :color="podcastItem.BookmarkDate ? 'red' : 'primary'"
          @click="changeFavStatus(podcastItem)"
        >
          {{
            podcastItem.BookmarkDate ? "mdi-heart" : "mdi-heart-outline"
          }}
        </v-icon>
      </template>
    </v-tooltip>
  </div>
</template>
