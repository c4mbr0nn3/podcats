<template>
  <podcast-player-card v-if="podcastData" :podcast-data="podcastData" />
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { getPodcastItemById } from "@/api";
import PodcastPlayerCard from "@/components/PodcastPlayerCard.vue";

const route = useRoute();

let podcastData = ref(null);

onMounted(async () => {
  await getPodcastItemById(route.params.id, route.params.itemId).then(
    (response) => {
      podcastData.value = response.data;
    }
  );
});
</script>
