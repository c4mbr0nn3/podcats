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
            :podcast-id="podcastItem.ID"
            :podcast-item="podcastItem"
            @change-fav-status="changeFavStatus"
            @change-played-status="changePlayedStatus"
          /> </v-card-text></v-card
      ><v-card v-intersect="onIntersect"></v-card>
    </v-col>
  </v-row>
</template>

<script>
import { PodcastItemsService } from "@/api";
import SinglePodcastItemCard from "@/components/SinglePodcastItemCard.vue";

export default {
  components: {
    SinglePodcastItemCard,
  },
  data: () => ({
    podcastItemsData: [],
    currentPage: 1,
    nextPage: null,
    pageCount: null,
  }),
  async mounted() {
    await this.fetchData(1);
  },
  methods: {
    async fetchData(pageId) {
      await PodcastItemsService.getLatestPodcastItems(pageId).then((response) => {
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
      if (this.currentPage == this.pageCount || !this.nextPage) return;
      await this.fetchData(this.nextPage);
    },
    changePlayedStatus(event) {
      let item = this.podcastItemsData.find((el) => el.ID == event);
      item.IsPlayed = !item.IsPlayed;
    },
    changeFavStatus(event) {
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
