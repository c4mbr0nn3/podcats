<template>
  <v-row align="center" justify="center">
    <v-col cols="8">
      <v-card
        ><v-card-title>I miei podcasts</v-card-title>
        <v-card-text
          ><v-text-field
            v-model="podcastUrl"
            label="Import Podcast"
            append-icon="mdi-plus"
            @click:append="importPodcast"
          ></v-text-field
          ><v-card
            v-for="(podcast, index) in podcastData"
            :key="index"
            class="mt-3"
          >
            <v-card-title> {{ podcast.Title }} </v-card-title>
            <v-card-text>
              <Markdown :source="podcast.Summary" />
            </v-card-text> </v-card
        ></v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import axios from "axios";
import Markdown from "vue3-markdown-it";
export default {
  components: {
    Markdown,
  },
  data: () => ({
    podcastData: null,
    podcastUrl: "",
  }),
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      await axios
        .get("api/v1/podcasts/")
        .then((response) => (this.podcastData = response.data.data));
    },
    async importPodcast() {
      let payload = { podcastUrl: this.podcastUrl };
      await axios.post("api/v1/podcasts/import", payload).then(() => {
        this.podcastUrl = null;
        alert("bellali importato");
        this.fetchData();
      });
    },
  },
};
</script>
