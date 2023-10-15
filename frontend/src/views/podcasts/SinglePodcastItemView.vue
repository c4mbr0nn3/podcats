<template>
  <podcast-player-card v-if="item" :podcast-data="item" />
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute } from "vue-router";
import { useTitle } from "@vueuse/core";
import { PodcastItemService } from "@/services";
import PodcastPlayerCard from "@/components/PodcastPlayerCard.vue";

const item = ref(null);

const route = useRoute();

const title = computed(() => {
  return item.value ? `PodCats | ${item.value.Title}` : "PodCats";
});

useTitle(title);

onMounted(async () => {
  item.value = await PodcastItemService.getById(route.params.itemId);
});
</script>
