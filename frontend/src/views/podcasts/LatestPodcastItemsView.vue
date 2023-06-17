<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card>
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
                @change-fav-status="changeFavStatus"
                @change-played-status="changePlayedStatus"
              />
            </v-lazy>
          </div>
        </v-card-text>
      </v-card>
      <v-card
        v-if="!isLoading || !isFetchingNextPage || !hasNextPage"
        ref="sentinel"
      />
      <v-row v-if="isLoading || isFetchingNextPage" justify="center">
        <v-col class="d-flex justify-center"> Loading... </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script setup>
import { PodcastItemsService } from "@/api";
import { useInfiniteQuery } from "@tanstack/vue-query";
import SinglePodcastItemCard from "@/components/SinglePodcastItemCard.vue";
import { computed, ref } from "vue";
import { useIntersectionObserver } from "@vueuse/core";
import { watch } from "vue";

const sentinel = ref(null);
const sentinelVisible = ref(false);

const { data, fetchNextPage, hasNextPage, isLoading, isFetchingNextPage } =
  useInfiniteQuery({
    queryKey: ["latest-podcast-items"],
    queryFn: fetchData,
    getNextPageParam: (lastPage) => {
      return lastPage.nextCursor;
    },
  });

watch(
  () => sentinelVisible.value,
  (value) => {
    if (value) {
      onIntersect();
    }
  }
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
  console.log("onIntersect");
  if (!hasNextPage.value) return;
  fetchNextPage();
}

function changePlayedStatus(event) {
  console.log("changePlayedStatus", event);
  // let item = podcastItemsData.value.find((el) => el.ID == event);
  // item.IsPlayed = !item.IsPlayed;
}

function changeFavStatus(event) {
  console.log("changeFavStatus", event);

  // let item = podcastItemsData.value.find((el) => el.ID == event.id);
  // item.BookmarkDate = event.bookmarkDate;
}

useIntersectionObserver(sentinel, ([{ isIntersecting }]) => {
  sentinelVisible.value = isIntersecting;
});
</script>
