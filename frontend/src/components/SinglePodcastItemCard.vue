<template>
  <podcast-card-base :podcast="props.podcastItem" :router-link="getRouterLink">
    <template #actions>
      <single-episode-actions
        :podcast-item="props.podcastItem"
        @change-status="$emit('change-status', $event)"
      />
    </template>
  </podcast-card-base>
</template>

<script setup>
import { computed } from "vue";
import PodcastCardBase from "@/components/global/podcast-card/PodcastCardBase.vue";
import SingleEpisodeActions from "@/components/global/podcast-card/SingleEpisodeActions.vue";

defineEmits(["change-status"]);

const props = defineProps({
  podcastId: { type: [String, Number], required: true },
  podcastItem: {
    type: Object,
    required: true,
  },
});

const getRouterLink = computed(() => {
  return {
    name: "single-item",
    params: { id: props.podcastItem.PodcastID, itemId: props.podcastItem.ID },
  };
});
</script>
