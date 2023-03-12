<template>
  <podcast-card-base :podcast="props.podcastItem" :router-link="getRouterLink">
    <template #actions>
      <single-episode-actions
        :podcast-item="props.podcastItem"
        @change-played-status="$emit('change-played-status', $event)"
        @change-fav-status="$emit('change-fav-status', $event)"
      />
    </template>
  </podcast-card-base>
</template>

<script setup>
import { computed } from "vue";
import PodcastCardBase from "@/components/global/podcast-card/PodcastCardBase.vue";
import SingleEpisodeActions from "@/components/global/podcast-card/SingleEpisodeActions.vue";

defineEmits(["change-played-status", "change-fav-status"]);

const props = defineProps({
  podcastId: { type: String, required: true },
  podcastItem: {
    type: Object,
    required: true,
  },
});

const getRouterLink = computed(() => {
  return {
    name: "single-item",
    params: { id: props.podcastId, itemId: props.podcastItem.ID },
  };
});
</script>
