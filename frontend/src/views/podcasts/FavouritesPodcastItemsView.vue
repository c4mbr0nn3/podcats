<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card
        ><v-card-title>Favorites Episodes</v-card-title>
        <v-card-text v-if="podcastItemsData.length > 0">
          <SinglePodcastItemCard
            v-for="(podcastItem, index) in getFilteredFavs"
            :key="index"
            class="mt-3"
            :podcast-id="podcastItem.PodcastID"
            :podcast-item="podcastItem"
            @change-fav-status="changeFavStatus"
            @change-played-status="changePlayedStatus" />
          <v-card v-intersect="onIntersect"></v-card
        ></v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { PodcastItemService } from "@/services";
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
      const data = await PodcastItemService.getFavourites(pageId);

      if (data.podcastItems.length > 0) {
        data.podcastItems.forEach((item) => this.podcastItemsData.push(item));
      }
      this.currentPage = data.thisPage;
      this.pageCount = data.pageCount;
      this.nextPage = data.nextPage;
    },
    // TODO: fix doppia chiamata non voluta...
    async onIntersect() {
      if (this.currentPage == this.pageCount) return;
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
