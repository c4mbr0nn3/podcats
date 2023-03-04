<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card>
        <v-card-title>
          {{ getCardTitle }}
        </v-card-title>
        <v-card-text>
          <PodcastCard
            v-for="(podcastItem, index) in getPodcastItemsList"
            :key="index"
            class="mt-3"
            :podcast="podcastItem"
            :router-link="{
              name: 'single-item',
              params: { id: $route.params.id, itemId: podcastItem.ID },
            }"
          >
            <template #actions>
              <SingleEpisodeActions
                :podcast-item="podcastItem"
                @change-played-status="changeStatus($event)"
                @change-fav-status="changeFavStatus($event)"
              />
            </template>
          </PodcastCard>
        </v-card-text>
      </v-card>
      <v-card v-intersect="onIntersect"></v-card>
    </v-col>
  </v-row>
</template>

<script>
import { getPodcastItemsByPodcastId, getPodcastById } from "@/api";
import { formatDate } from "@/utils/date";
import PodcastCard from "@/components/PodcastCard.vue";
import SingleEpisodeActions from "@/components/global/SingleEpisodeActions.vue";

export default {
  components: {
    PodcastCard,
    SingleEpisodeActions,
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
    getCardTitle() {
      return this.podcastData ? this.podcastData.Title : "";
    },
    getPodcastItemsList() {
      return this.podcastItemsData ? this.podcastItemsData : [];
    },
  },
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      await getPodcastById(this.$route.params.id).then(
        (response) => (this.podcastData = response.data)
      );
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
    changeStatus(event) {
      let item = this.podcastItemsData.find((el) => el.ID == event);
      item.IsPlayed = !item.IsPlayed;
    },
    async changeFavStatus(event) {
      let item = this.podcastItemsData.find((el) => el.ID == event.id);
      item.BookmarkDate = event.bookmarkDate;
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
