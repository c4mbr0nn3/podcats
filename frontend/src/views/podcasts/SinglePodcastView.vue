<template>
  <v-row align="center" justify="center">
    <v-col cols="9">
      <v-card
        v-for="(podcast, index) in getPodcastItemsList"
        :key="index"
        class="mt-3"
        :to="{
          name: 'single-item',
          params: { id: $route.params.id, itemId: podcast.ID },
        }"
      >
        <div class="d-flex flex-no-wrap justify-space-between">
          <v-col cols="10">
            <v-card-title class="text-truncate">{{
              podcast.Title
            }}</v-card-title>
            <v-card-text
              ><Markdown
                class="overflow-hidden fade"
                :source="podcast.Summary"
              />
              <v-chip label class="mt-2" color="primary" variant="outlined">
                <v-icon start icon="mdi-calendar"></v-icon>
                {{ formatDate(podcast.PublicationDate) }}
              </v-chip>
            </v-card-text>
          </v-col>
          <v-col cols="2" class="d-flex justify-end ml-2">
            <v-avatar class="ma-3" size="125" rounded="0">
              <v-img :src="podcast.Image ? podcast.Image : missingImage">
                <template #placeholder>
                  <div class="d-flex align-center justify-center fill-height">
                    <v-progress-circular
                      indeterminate
                      color="primary"
                    ></v-progress-circular></div></template
              ></v-img>
            </v-avatar>
          </v-col>
        </div>
      </v-card>
      <v-card v-intersect="onIntersect"></v-card>
    </v-col>
  </v-row>
</template>

<script>
import missingImage from "@/assets/missing_image.png";
import { getPodcastById, getPodcastItemsByPodcastId } from "@/api";
import { formatDate } from "@/utils/date";
import Markdown from "vue3-markdown-it";

export default {
  components: {
    Markdown,
  },
  data: () => ({
    podcastData: null,
    podcastItemsData: [],
    missingImage: missingImage,

    formatDate,
  }),
  computed: {
    getPodcastItemsList() {
      return this.podcastItemsData ? this.podcastItemsData : [];
    },
  },
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      await getPodcastById(this.$route.params.id).then((response) => {
        this.podcastData = response.data;
      });
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
      if (this.currentPage == this.pageCount) return;
      await this.fetchPodcastItems(this.nextPage);
    },
  },
};
</script>

<style scoped>
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
