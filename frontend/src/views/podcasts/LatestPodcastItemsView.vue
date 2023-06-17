<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card
        ><v-card-title>Latest Episodes</v-card-title>
        <v-card-text v-if="podcastItemsData.length > 0">
          <SinglePodcastItemCard
            v-for="(podcastItem, index) in podcastItemsData"
            :key="index"
            class="mt-3"
            :podcast-id="podcastItem.PodcastID"
            :podcast-item="podcastItem"
            @change-fav-status="changeFavStatus"
            @change-played-status="changePlayedStatus"
          /> </v-card-text></v-card
      ><v-card v-intersect="onIntersect"></v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { PodcastItemsService } from "@/api";
import { useInfiniteQuery } from "@tanstack/vue-query";
import { ref } from "vue";
import SinglePodcastItemCard from "@/components/SinglePodcastItemCard.vue";

const podcastItemsData = ref([]);
const currentPage = ref(1);
const nextPage = ref(null);
const pageCount = ref(null);

async function fetchData(pageId) {
  await PodcastItemsService.getLatestPodcastItems(pageId).then((response) => {
    if (response.data.podcastItems.length > 0) {
      response.data.podcastItems.forEach((item) =>
        podcastItemsData.value.push(item)
      );
    }
    currentPage.value = response.data.thisPage;
    pageCount.value = response.data.pageCount;
    nextPage.value = response.data.nextPage;
  });
}

// TODO: fix doppia chiamata non voluta...
async function onIntersect() {
  if (currentPage.value == pageCount.value || !nextPage.value) return;
  await fetchData(nextPage.value);
}

function changePlayedStatus(event) {
  let item = podcastItemsData.value.find((el) => el.ID == event);
  item.IsPlayed = !item.IsPlayed;
}

function changeFavStatus(event) {
  let item = podcastItemsData.value.find((el) => el.ID == event.id);
  item.BookmarkDate = event.bookmarkDate;
}
</script>
