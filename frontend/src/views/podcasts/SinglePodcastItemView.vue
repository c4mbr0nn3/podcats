<template>
  <v-row align="center" justify="center">
    <v-col cols="8">
      <v-card
        ><v-card-title>{{ getPodcastItemTitle }}</v-card-title>
        <div class="d-flex flex-no-wrap justify-space-between">
          <v-col cols="10">
            <v-card-text
              ><Markdown :source="podcastData ? podcastData.Summary : ''" />
            </v-card-text>
            <v-card-actions
              ><v-btn
                :icon="isPlaying ? 'mdi-pause' : 'mdi-play'"
                :color="isPlaying ? 'accent' : 'primary'"
                @click="switchTrackStatus"
              ></v-btn
              ><v-btn icon="mdi-stop" color="amber" @click="stopTrack"></v-btn
              ><v-btn
                :icon="isMuted ? 'mdi-volume-mute' : 'mdi-volume-high'"
                @click="muteTrack"
              ></v-btn
            ></v-card-actions>
          </v-col>
          <v-col cols="2" class="d-flex justify-end">
            <v-avatar class="ma-3" size="125" rounded="0">
              <v-img :src="podcastData ? podcastData.Image : ''"></v-img>
            </v-avatar>
          </v-col>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { getPodcastItemById } from "@/api";
import { Howl, Howler } from "howler";
import Markdown from "vue3-markdown-it";
export default {
  components: {
    Markdown,
  },
  data: () => ({
    podcastData: null,
    howlerData: null,
    isPlaying: false,
    isMuted: false,
  }),
  computed: {
    getPodcastItemTitle() {
      return this.podcastData ? this.podcastData.Title : "";
    },
  },
  async created() {
    await this.fetchData();
  },
  unmounted() {
    Howler.unload();
  },
  methods: {
    async fetchData() {
      await getPodcastItemById(
        this.$route.params.id,
        this.$route.params.itemId
      ).then((response) => {
        this.podcastData = response.data.data;
        let howlerConfig = {
          src: [this.podcastData.FileURL],
          volume: 0.5,
          loop: false,
          html5: true,
          preload: "metadata",
        };
        this.howlerData = new Howl(howlerConfig);
      });
    },
    switchTrackStatus() {
      this.isPlaying ? this.pauseTrack() : this.playTrack();
    },
    playTrack() {
      if (this.isPlaying) return;
      this.howlerData.play();
      this.isPlaying = true;
    },
    pauseTrack() {
      if (!this.isPlaying) return;
      this.howlerData.pause();
      this.isPlaying = false;
    },
    stopTrack() {
      if (!this.isPlaying) return;
      this.howlerData.stop();
      this.isPlaying = false;
    },
    muteTrack() {
      this.isMuted ? Howler.mute(false) : Howler.mute(true);
      this.isMuted = !this.isMuted;
    },
  },
};
</script>
