<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card
        v-for="(podcastItem, index) in getPodcastItemsList"
        :key="index"
        class="mt-3"
      >
        <div class="d-flex flex-no-wrap justify-space-between">
          <v-col cols="10">
            <v-card-title class="text-truncate">{{
              podcastItem.Title
            }}</v-card-title>
            <v-card-text>
              <Markdown
                class="overflow-hidden fade"
                :source="podcastItem.Summary"
              />
              <div class="d-flex align-center mt-2">
                <v-chip label color="primary" variant="outlined">
                  <v-icon start icon="mdi-calendar"></v-icon>
                  {{ formatDate(podcastItem.PublicationDate) }}
                </v-chip>
                <v-tooltip
                  :text="
                    podcastItem.IsPlayed
                      ? 'Mark as not-played'
                      : 'Mark as played'
                  "
                  location="bottom"
                >
                  <template #activator="{ props }">
                    <v-icon
                      v-bind="props"
                      class="mx-3"
                      :color="podcastItem.IsPlayed ? 'success' : 'primary'"
                      @click="changeStatus(podcastItem)"
                      >{{
                        podcastItem.IsPlayed
                          ? "mdi-checkbox-marked-circle-outline"
                          : "mdi-checkbox-blank-circle-outline"
                      }}</v-icon
                    >
                  </template>
                </v-tooltip>
                <v-tooltip
                  :text="
                    podcastItem.BookmarkDate
                      ? 'Remove from favs'
                      : 'Add to favs'
                  "
                  location="bottom"
                >
                  <template #activator="{ props }">
                    <v-icon
                      v-bind="props"
                      class="mr-3"
                      :color="podcastItem.BookmarkDate ? 'red' : 'primary'"
                      @click="changeFavStatus(podcastItem)"
                      >{{
                        podcastItem.BookmarkDate
                          ? "mdi-heart"
                          : "mdi-heart-outline"
                      }}</v-icon
                    >
                  </template>
                </v-tooltip>
              </div>
            </v-card-text>
          </v-col>
          <v-col cols="2" class="d-flex justify-end ml-2">
            <PodcastAvatar
              :image="podcastItem.Image"
              :router-link="{
                name: 'single-item',
                params: { id: $route.params.id, itemId: podcastItem.ID },
              }"
            />
          </v-col>
        </div>
      </v-card>
      <v-card v-intersect="onIntersect"></v-card>
    </v-col>
  </v-row>
</template>

<script>
import {
  getPodcastItemsByPodcastId,
  switchPodcastItemPlayedStatus,
  switchPodcastItemFavouriteStatus,
} from "@/api";
import { formatDate } from "@/utils/date";
import Markdown from "vue3-markdown-it";
import PodcastAvatar from "@/components/global/PodcastAvatar.vue";

export default {
  components: {
    Markdown,
    PodcastAvatar,
  },
  data: () => ({
    podcastData: null,
    podcastItemsData: [],
    currentPage: 1,
    nextPage: null,
    pageCount: null,

    formatDate,
  }),
  computed: {
    getPodcastItemsList() {
      return this.podcastItemsData ? this.podcastItemsData : [];
    },
  },
  // TODO: perchè diavolo parte una chiamata con page=null?
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      await this.fetchPodcastItems(1);
    },
    async fetchPodcastItems(pageId) {
      await getPodcastItemsByPodcastId(this.$route.params.id, pageId).then(
        (response) => {
          if (response.data.podcastItems.length > 0) {
            response.data.podcastItems.forEach((item) =>
              this.podcastItemsData.push(item)
            );
          }
          this.currentPage = response.data.thisPage;
          this.pageCount = response.data.pageCount;
          this.nextPage = response.data.nextPage;
        }
      );
    },
    // TODO: fix doppia chiamata non voluta...
    async onIntersect() {
      if (this.currentPage == this.pageCount || !this.nextPage) return;
      await this.fetchPodcastItems(this.nextPage);
    },
    async changeStatus(podcastItem) {
      await switchPodcastItemPlayedStatus(
        this.$route.params.id,
        podcastItem.ID
      ).then(() => (podcastItem.IsPlayed = !podcastItem.IsPlayed));
    },
    async changeFavStatus(podcastItem) {
      await switchPodcastItemFavouriteStatus(
        this.$route.params.id,
        podcastItem.ID
      ).then((response) => (podcastItem.BookmarkDate = response.data));
    },
  },
};
</script>

<style scoped>
.v-avatar {
  transition: opacity 0.4s ease-in-out;
}
.v-avatar:not(.on-hover) {
  opacity: 0.6;
}

.show-btns {
  color: rgba(var(--v-theme-primary), 1) !important;
  opacity: 1 !important;
}

/* https://css-tricks.com/line-clampin/ */
.fade {
  position: relative;
  height: 6em; /* exactly four lines with line height = 1.5*/
}
.fade:after {
  content: "";
  text-align: right;
  position: absolute;
  bottom: 0;
  right: 0;
  width: 70%;
  height: 1.5em;
  background: linear-gradient(
    to right,
    rgba(var(--v-theme-surface), 0),
    rgba(var(--v-theme-surface), 1) 50%
  );
}
</style>
