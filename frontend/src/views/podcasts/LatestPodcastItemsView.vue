<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card class="mb-5">
        <v-card-title>Latest Episodes</v-card-title>
        <v-card-text v-if="podcastItemList.length > 0">
          <div v-for="(podcastItem, id) in podcastItemList" :key="id">
            <v-lazy
              :options="{
                threshold: 0.25,
              }"
              transition="fade-transition"
              min-height="220"
              class="mt-3"
            >
              <SinglePodcastItemCard
                :podcast-id="podcastItem.PodcastID"
                :podcast-item="podcastItem"
                @change-status="changeStatus"
              />
            </v-lazy>
          </div>
        </v-card-text>
      </v-card>
      <div
        v-if="!isLoading || !isFetchingNextPage || !hasNextPage"
        ref="sentinel"
      ></div>
    </v-col>
  </v-row>
</template>

<script setup>
import { PodcastItemsService } from "@/api";
import { useInfiniteQuery } from "@tanstack/vue-query";
import SinglePodcastItemCard from "@/components/SinglePodcastItemCard.vue";
import { computed, ref } from "vue";
import { useIntersectionObserver, watchDebounced } from "@vueuse/core";

const sentinel = ref(null);
const sentinelVisible = ref(false);

const {
  data,
  fetchNextPage,
  hasNextPage,
  isLoading,
  isFetchingNextPage,
  refetch,
} = useInfiniteQuery({
  queryKey: ["latest-podcast-items"],
  queryFn: fetchData,
  getNextPageParam: (lastPage) => {
    return lastPage.nextCursor;
  },
});

useIntersectionObserver(sentinel, ([{ isIntersecting }]) => {
  sentinelVisible.value = isIntersecting;
});

watchDebounced(
  sentinelVisible,
  (value) => {
    if (value) onIntersect();
  },
  { debounce: 500, maxWait: 1000 }
);

const podcastItemList = computed(() => {
  if (!data.value) return [];
  return data.value.pages.reduce((acc, page) => {
    return [...acc, ...page.data];
  }, []);
});

async function fetchData({ pageParam = 1 }) {
  return await PodcastItemsService.getLatestPodcastItems(pageParam).then(
    (response) => {
      return {
        data: response.data.podcastItems,
        nextCursor: response.data.nextPage ? response.data.nextPage : undefined,
      };
    }
  );
}

async function onIntersect() {
  if (!hasNextPage.value) return;
  fetchNextPage();
}

function changeStatus(event) {
  refetch({
    refetchPage: (page, index) => index == getPodcastItemPage(event),
  });
}

function getPodcastItemPage(podcastItemId) {
  return data.value.pages.findIndex(
    (page) =>
      page.data.findIndex((podcastItem) => podcastItem.ID == podcastItemId) > -1
  );
}
</script>
