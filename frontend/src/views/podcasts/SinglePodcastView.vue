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
          <div>
            <v-card-title>{{ podcast.Title }}</v-card-title>
          </div>
          <v-avatar class="ma-3" size="125" rounded="0">
            <v-img :src="podcast.Image"></v-img>
          </v-avatar>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { getPodcastById } from "@/api";
export default {
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
