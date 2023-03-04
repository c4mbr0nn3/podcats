<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card
        ><v-card-title>My Podcasts</v-card-title>
        <v-card-text
          ><v-text-field
            v-model="podcastUrl"
            label="Import Podcast"
            color="primary"
            variant="underlined"
          >
            <template #append-inner>
              <v-icon color="primary" @click="importPodcast">
                mdi-plus
              </v-icon></template
            ></v-text-field
          ><PodcastCard
            v-for="(podcast, index) in podcastData"
            :key="index + podcast.UpdatedAt"
            :podcast="podcast"
            :router-link="{
              name: 'single-podcast',
              params: { id: podcast.ID },
            }"
            ><template #actions>
              <SinglePodcastActions
                :podcast="podcast"
                @mark-all-played="triggerFetch()"
                @delete-podcast="triggerFetch()"
              /> </template
          ></PodcastCard>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { importPodcast, getAllPodcasts } from "@/api";
import PodcastCard from "@/components/PodcastCard.vue";
import SinglePodcastActions from "@/components/global/SinglePodcastActions.vue";

export default {
  components: {
    PodcastCard,
    SinglePodcastActions,
  },
  data: () => ({
    podcastData: null,
    podcastUrl: "",
  }),
  created() {
    this.triggerFetch();
  },
  methods: {
    async triggerFetch() {
      await this.fetchData();
    },
    async fetchData() {
      await getAllPodcasts().then(
        (response) => (this.podcastData = response.data)
      );
    },
    async importPodcast() {
      let payload = { podcastUrl: this.podcastUrl };
      await importPodcast(payload).then(() => {
        this.podcastUrl = null;
        this.fetchData();
      });
    },
  },
};
</script>
