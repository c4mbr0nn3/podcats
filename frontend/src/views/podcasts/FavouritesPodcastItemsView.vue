<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card
        ><v-card-title>Favourites Episodes</v-card-title>
        <v-card-text v-if="podcastItemsData.length > 0"
          ><v-card
            v-for="(podcastItem, index) in getFilteredFavs"
            :key="index"
            class="mt-3"
          >
            <div class="d-flex flex-no-wrap justify-space-between">
              <v-col cols="10">
                <v-card-title class="text-truncate">{{
                  podcastItem.Title
                }}</v-card-title>
                <v-card-text
                  ><p-markdown
                    class="overflow-hidden fade"
                    :markdown="podcastItem.Summary"
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
                    params: {
                      id: podcastItem.PodcastID,
                      itemId: podcastItem.ID,
                    },
                  }"
                />
              </v-col></div></v-card
          ><v-card v-intersect="onIntersect"></v-card
        ></v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import {
  getFavouritesPodcastItems,
  switchPodcastItemPlayedStatus,
  switchPodcastItemFavouriteStatus,
} from "@/api";
import { formatDate } from "@/utils/date";
import PMarkdown from "@/components/PMarkdown.vue";
import PodcastAvatar from "@/components/global/PodcastAvatar.vue";

export default {
  components: {
    PMarkdown,
    PodcastAvatar,
  },
  data: () => ({
    podcastItemsData: [],
    currentPage: 1,
    nextPage: null,
    pageCount: null,

    formatDate,
  }),
  computed: {
    getFilteredFavs() {
      return this.podcastItemsData.filter((el) => el.BookmarkDate);
    },
  },
  async mounted() {
    await this.fetchData(1);
  },
  methods: {
    async fetchData(pageId) {
      await getFavouritesPodcastItems(pageId).then((response) => {
        if (response.data.podcastItems.length > 0) {
          response.data.podcastItems.forEach((item) =>
            this.podcastItemsData.push(item)
          );
        }
        this.currentPage = response.data.thisPage;
        this.pageCount = response.data.pageCount;
        this.nextPage = response.data.nextPage;
      });
    },
    // TODO: fix doppia chiamata non voluta...
    async onIntersect() {
      if (this.currentPage == this.pageCount) return;
      await this.fetchData(this.nextPage);
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
