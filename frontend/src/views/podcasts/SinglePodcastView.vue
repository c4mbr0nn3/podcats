<template>
  <v-row align="center" justify="center">
    <v-col cols="8">
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
                :source="podcastData ? podcastData.Summary : ''"
              />
            </v-card-text>
          </v-col>
          <v-col cols="2" class="d-flex justify-end ml-2">
            <v-avatar class="ma-3" size="125" rounded="0">
              <v-img :src="podcast.Image"></v-img>
            </v-avatar>
          </v-col>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { getPodcastById } from "@/api";
import Markdown from "vue3-markdown-it";

export default {
  components: {
    Markdown,
  },
  data: () => ({
    podcastData: null,
  }),
  computed: {
    getPodcastItemsList() {
      return this.podcastData ? this.podcastData.PodcastItems : [];
    },
  },
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      await getPodcastById(this.$route.params.id).then((response) => {
        this.podcastData = response.data.data;
      });
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
