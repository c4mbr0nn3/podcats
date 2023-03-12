<template>
  <v-row justify="center">
    <v-col cols="8">
      <podcast-card-base :podcast="props.podcastData" :apply-fade="false"
        ><template #other>
          <div>
            <v-col cols="12">
              <v-card-text>
                <HowlerPlayer :data="props.podcastData"
              /></v-card-text>
            </v-col>
          </div> </template
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
