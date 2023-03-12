<template>
  <v-row justify="center">
    <v-col cols="8">
      <podcast-card-base :podcast="props.podcastData"
        ><template #other> <HowlerPlayer :data="props.podcastData" /></template
      ></podcast-card-base>
    </v-col>
  </v-row>
</template>

<script setup>
import HowlerPlayer from "@/components/global/podcast-card/HowlerPlayer.vue";
import PodcastCardBase from "@/components/global/podcast-card/PodcastCardBase.vue";

const props = defineProps({
  podcastData: {
    type: Object,
    required: true,
  },
});
</script>
