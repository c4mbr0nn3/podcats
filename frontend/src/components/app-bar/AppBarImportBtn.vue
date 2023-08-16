<template>
  <v-menu v-model="menu" :close-on-content-click="false">
    <template #activator="{ props }">
      <v-btn color="primary" icon="mdi-plus" size="small" v-bind="props" />
    </template>
    <v-card width="450">
      <v-card-title> Import Podcast </v-card-title>
      <v-card-text>
        <v-text-field
          v-model="podcastSearch"
          label="Start typing to search in Apple Podcasts"
          append-inner-icon="mdi-magnify"
          color="primary"
          hide-details
          @input="onSearch"
        />
        <v-responsive
          v-if="podcastSearch"
          height="450"
          class="overflow-y-auto mt-2"
        >
          <div
            v-if="isSearching && !isDelayElapsed"
            class="d-flex align-center justify-center h-100"
          >
            <Icon
              icon="svg-spinners:3-dots-fade"
              color="secondary"
              width="50"
            />
          </div>
          <div v-else-if="searchResults.length > 0">
            <ItunesResultsCard
              v-for="(podcast, i) in searchResults"
              :key="podcast.collectionId"
              :podcast="podcast"
              :class="{ 'mb-2': i !== searchResults.length - 1 }"
            />
          </div>
          <div
            v-else
            class="d-flex flex-column align-center justify-center h-100 text-disabled"
          >
            <Icon icon="mdi:emoticon-dead-outline" width="100" />
            <div class="mt-1">No results found</div>
          </div>
        </v-responsive>
        <v-card-title
          v-if="!podcastSearch"
          class="d-flex align-center text-overline"
        >
          or <v-divider inset color="primary" />
        </v-card-title>
        <v-text-field
          v-if="!podcastSearch"
          v-model="podcastUrl"
          label="Paste podcast feed URL to import"
          color="primary"
          hide-details
        >
          <template #append-inner>
            <v-btn
              :disabled="!podcastUrl"
              color="primary"
              icon="mdi-microphone-plus"
              density="compact"
              variant="text"
              :loading="isImporting"
              @click="onImport"
            />
          </template>
        </v-text-field>
      </v-card-text>
    </v-card>
  </v-menu>
</template>

<script setup>
import { ref, computed } from "vue";
import { usePodcastsStore } from "@/stores/podcasts";
import { ItunesService } from "@/services";
import { useDebounceFn, whenever } from "@vueuse/core";
import ItunesResultsCard from "./AppBarItunesResultsCard.vue";
import { Icon } from "@iconify/vue";

const podcastSearch = ref("");
const podcastUrl = ref("");
const isImporting = ref(false);
const isSearching = ref(false);
const isDelayElapsed = ref(true);
const searchResults = ref([]);
const menu = ref(false);

const podcastsStore = usePodcastsStore();
const { importPodcast } = podcastsStore;

const isMenuClosed = computed(() => !menu.value);

whenever(isMenuClosed, () => {
  podcastSearch.value = "";
  podcastUrl.value = "";
  searchResults.value = [];
});

const onImport = async () => {
  isImporting.value = true;
  importPodcast(podcastUrl);
  podcastUrl.value = "";
  isImporting.value = false;
};

const onSearch = async () => {
  isSearching.value = true;
  isDelayElapsed.value = false;
  searchResults.value = [];
  if (!podcastSearch.value) return;
  debouncedSearch();
};

const debouncedSearch = useDebounceFn(async () => {
  searchResults.value = await ItunesService.search(podcastSearch);
  isSearching.value = false;
  setTimeout(() => {
    isDelayElapsed.value = true;
  }, 200);
}, 1000);
</script>
