<script setup>
import { computed } from 'vue'
import PodcastCardBase from '@/components/global/podcast-card/PodcastCardBase.vue'
import SinglePodcastActions from '@/components/global/podcast-card/SinglePodcastActions.vue'

const props = defineProps({
  podcast: {
    type: Object,
    required: true,
  },
})

defineEmits(['mark-all-played', 'delete-podcast', 'show-info-dialog'])

const getRouterLink = computed(() => {
  return {
    name: 'single-podcast',
    params: { id: props.podcast.ID },
  }
})
</script>

<template>
  <PodcastCardBase :podcast="props.podcast" :router-link="getRouterLink">
    <template #actions>
      <SinglePodcastActions
        :podcast="props.podcast"
        @mark-all-played="$emit('mark-all-played')"
        @delete-podcast="$emit('delete-podcast')"
        @show-info-dialog="$emit('show-info-dialog', $event)"
      />
    </template>
  </PodcastCardBase>
</template>
