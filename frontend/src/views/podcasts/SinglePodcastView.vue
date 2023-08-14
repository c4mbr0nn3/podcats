<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card>
        <v-card-title>
          {{ getCardTitle }}
        </v-card-title>
        <v-card-text>
          <SinglePodcastItemCard
            v-for="(podcastItem, index) in getPodcastItemsList"
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

<script>
import { PodcastService } from "@/services";
import SinglePodcastItemCard from "@/components/SinglePodcastItemCard.vue";

export default {
  components: {
    SinglePodcastItemCard,
  },
  data: () => ({
    podcastData: null,
    podcastItemsData: [],
    currentPage: 1,
    nextPage: null,
    pageCount: null,
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
      this.podcastData = await PodcastService.getById(this.$route.params.id);
      await this.fetchPodcastItems(1);
    },
    async fetchPodcastItems(pageId) {
      const data = await PodcastService.getItemsById(
        this.$route.params.id,
        pageId
      );
      if (data.podcastItems.length > 0) {
        data.podcastItems.forEach((item) => this.podcastItemsData.push(item));
      }
      this.currentPage = data.thisPage;
      this.pageCount = data.pageCount;
      this.nextPage = data.nextPage;
    },
    // TODO: fix doppia chiamata non voluta...
    async onIntersect() {
      if (this.currentPage == this.pageCount || !this.nextPage) return;
      await this.fetchPodcastItems(this.nextPage);
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
