<template>
  <podcast-player-card v-if="item" :podcast-data="item" />
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { PodcastItemService } from "@/services";
import PodcastPlayerCard from "@/components/PodcastPlayerCard.vue";

const item = ref(null);

const route = useRoute();

onMounted(async () => {
  item.value = await PodcastItemService.getById(route.params.itemId);
});
</script>
