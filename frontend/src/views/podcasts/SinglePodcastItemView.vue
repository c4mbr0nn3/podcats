<template>
  <v-row justify="center">
    <v-col cols="8">
      <v-card v-if="podcastData">
        <template #loader="{ isActive }">
          <v-progress-linear
            :active="isActive"
            color="primary"
            height="4"
            indeterminate
          ></v-progress-linear>
        </template>
        <div class="d-flex flex-no-wrap justify-space-between">
          <v-col cols="10">
            <v-card-title class="text-truncate">{{
              getPodcastItemTitle
            }}</v-card-title>
            <v-card-text
              ><Markdown :source="podcastData.Summary" />
            </v-card-text>
          </v-col>
          <v-col cols="2" class="d-flex justify-end ml-2">
            <v-avatar class="ma-3" size="125" rounded="lg">
              <v-img :src="podcastData.Image || missingImage">
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
        <div>
          <v-col cols="12">
            <v-card-text>
              <HowlerPlayer v-if="podcastData" :data="podcastData" />
            </v-card-text>
          </v-col>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import missingImage from "@/assets/missing_image.png";
import { getPodcastItemById } from "@/api";
import Markdown from "vue3-markdown-it";
import HowlerPlayer from "@/components/HowlerPlayer.vue";
export default {
  components: {
    Markdown,
    HowlerPlayer,
  },
  data: () => ({
    podcastData: null,
    missingImage: missingImage,
  }),
  computed: {
    getPodcastItemTitle() {
      return this.podcastData ? this.podcastData.Title : "";
    },
  },
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      await getPodcastItemById(
        this.$route.params.id,
        this.$route.params.itemId
      ).then((response) => {
        this.podcastData = response.data;
      });
    },
  },
};
</script>
