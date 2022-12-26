<template>
  <v-row align="center" justify="center">
    <v-col cols="9">
      <v-card
        ><v-card-title>Ultimi episodi</v-card-title>
        <v-card-text v-if="podcastItemsData.length > 0"
          ><v-card
            v-for="(podcastItem, index) in podcastItemsData"
            :key="index"
            class="mt-3"
            :to="{
              name: 'single-item',
              params: { id: podcastItem.PodcastID, itemId: podcastItem.ID },
            }"
          >
            <div class="d-flex flex-no-wrap justify-space-between">
              <v-col cols="10">
                <v-card-title class="text-truncate">{{
                  podcastItem.Title
                }}</v-card-title>
                <v-card-text
                  ><Markdown
                    class="overflow-hidden fade"
                    :source="podcastItem.Summary"
                  />
                  <v-chip label class="mt-2" color="primary" variant="outlined">
                    <v-icon start icon="mdi-calendar"></v-icon>
                    {{ formatDate(podcastItem.PublicationDate) }}
                  </v-chip>
                </v-card-text>
              </v-col>
              <v-col cols="2" class="d-flex justify-end ml-2">
                <v-avatar class="ma-3" size="125" rounded="0">
                  <v-img :src="podcastItem.Image"></v-img>
                </v-avatar>
              </v-col></div></v-card
          ><v-card v-intersect="onIntersect"></v-card
        ></v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { getLatestPodcastItems } from "@/api";
import { formatDate } from "@/utils/date";
import Markdown from "vue3-markdown-it";
export default {
  components: {
    Markdown,
  },
  data: () => ({
    podcastItemsData: [],
    currentPage: 1,
    nextPage: null,
    pageCount: null,

    formatDate,
  }),
  async mounted() {
    await this.fetchData(1);
  },
  methods: {
    async fetchData(pageId) {
      await getLatestPodcastItems(pageId).then((response) => {
        if (response.data.data.length > 0) {
          response.data.data.forEach((item) =>
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
  },
};
</script>

<style>
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
