<template>
  <podcast-card-base
    :podcast="props.podcast"
    :avatar-size="80"
    :to="getRouterLink"
  ></podcast-card-base>
</template>

<script setup>
import { computed } from "vue";
import PodcastCardBase from "@/components/global/podcast-card/PodcastCardBase.vue";

const props = defineProps({
  podcast: {
    type: Object,
    required: true,
  },
});

const getRouterLink = computed(() => {
  return {
    name: "single-podcast",
    params: { id: props.podcast.ID },
  };
});
</script>
