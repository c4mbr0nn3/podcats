<script setup>
import { computed } from 'vue'
import PodcastCardBase from '@/components/global/podcast-card/PodcastCardBase.vue'
import SingleEpisodeActions from '@/components/global/podcast-card/SingleEpisodeActions.vue'

const props = defineProps({
  podcastId: { type: [String, Number], required: true },
  podcastItem: {
    type: Object,
    required: true,
  },
})

defineEmits(['change-played-status', 'change-fav-status'])

const getRouterLink = computed(() => {
  return {
    name: 'single-item',
    params: { id: props.podcastItem.PodcastID, itemId: props.podcastItem.ID },
  }
})
</script>

<template>
  <PodcastCardBase :podcast="props.podcastItem" :router-link="getRouterLink">
    <template #actions>
      <SingleEpisodeActions
        :podcast-item="props.podcastItem"
        @change-played-status="$emit('change-played-status', $event)"
        @change-fav-status="$emit('change-fav-status', $event)"
      />
    </template>
  </PodcastCardBase>
</template>
