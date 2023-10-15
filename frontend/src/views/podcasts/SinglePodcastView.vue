<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card>
        <v-card-title>
          {{ cardTitle }}
        </v-card-title>
        <v-card-text>
          <SinglePodcastItemCard
            v-for="(podcastItem, index) in podcastItems"
            :key="index"
            class="mt-3"
            :podcast-id="$route.params.id"
            :podcast-item="podcastItem"
            @change-fav-status="changeFavStatus"
            @change-played-status="changePlayedStatus"
          />
        </v-card-text>
      </v-card>
      <v-card v-intersect="onIntersect"></v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRoute } from "vue-router";
import { PodcastService } from "@/services";
import SinglePodcastItemCard from "@/components/SinglePodcastItemCard.vue";
import { useTitle } from "@vueuse/core";

const podcastData = ref(null);
const podcastItemsData = ref([]);
const currentPage = ref(1);
const nextPage = ref(null);
const pageCount = ref(null);

const route = useRoute();

const cardTitle = computed(() => {
  return podcastData.value ? podcastData.value.Title : "";
});

const podcastItems = computed(() => {
  return podcastItemsData.value ? podcastItemsData.value : [];
});

const title = computed(() => {
  return podcastData.value ? `PodCats | ${cardTitle.value}` : "PodCats";
});

useTitle(title);

onMounted(async () => {
  await fetchData();
});

const fetchData = async () => {
  podcastData.value = await PodcastService.getById(route.params.id);
  await fetchPodcastItems(1);
};

const fetchPodcastItems = async (pageId) => {
  const data = await PodcastService.getItemsById(route.params.id, pageId);
  if (data.podcastItems.length > 0) {
    data.podcastItems.forEach((item) => podcastItemsData.value.push(item));
  }
  currentPage.value = data.thisPage;
  pageCount.value = data.pageCount;
  nextPage.value = data.nextPage;
};

const onIntersect = async () => {
  if (currentPage.value == pageCount.value || !nextPage.value) return;
  await fetchPodcastItems(nextPage.value);
};

const changePlayedStatus = (event) => {
  let item = podcastItemsData.value.find((el) => el.ID == event);
  item.IsPlayed = !item.IsPlayed;
};

const changeFavStatus = (event) => {
  let item = podcastItemsData.value.find((el) => el.ID == event.id);
  item.BookmarkDate = event.bookmarkDate;
};
</script>
