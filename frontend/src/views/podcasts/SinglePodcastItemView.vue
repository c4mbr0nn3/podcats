<template>
  <v-row align="center" justify="center">
    <v-col cols="8">
      <v-card>
        <div class="d-flex flex-no-wrap justify-space-between">
          <v-col cols="10">
            <v-card-title class="text-truncate">{{
              getPodcastItemTitle
            }}</v-card-title>
            <v-card-text
              ><Markdown :source="podcastData ? podcastData.Summary : ''" />
            </v-card-text>
          </v-col>
          <v-col cols="2" class="d-flex justify-end ml-2">
            <v-avatar class="ma-3" size="125" rounded="0">
              <v-img :src="podcastData ? podcastData.Image : ''"></v-img>
            </v-avatar>
          </v-col>
        </div>
        <div>
          <v-col cols="12">
            <!--qui mettiamo tutti i controlli del player da spostare in component-->
            <v-row no-gutters
              >{{ timePlayed }} <v-spacer></v-spacer> {{ totalTime }}</v-row
            >
            <v-row no-gutters class="mt-1">
              <v-progress-linear
                v-model="getRelativeProgress"
                color="primary"
              ></v-progress-linear
            ></v-row>
            <v-row no-gutters justify="center" class="mt-2">
              <v-card-actions>
                <v-btn
                  :icon="isPlaying ? 'mdi-pause' : 'mdi-play'"
                  :color="isPlaying ? 'info' : 'primary'"
                  @click="switchTrackStatus"
                ></v-btn
                ><v-btn
                  icon="mdi-stop"
                  color="secondary"
                  class="mx-1"
                  @click="stopTrack"
                ></v-btn
                ><v-btn
                  :icon="isMuted ? 'mdi-volume-mute' : 'mdi-volume-high'"
                  @click="muteTrack"
                ></v-btn>
              </v-card-actions>
              <v-slider
                v-model="podcastVolume"
                :min="0"
                :max="1"
                :step="0.05"
                color="primary"
                thumb-label
                @update:model-value="setTrackVolume()"
              ></v-slider>
            </v-row>
          </v-col>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { getPodcastItemById } from "@/api";
import { Howl, Howler } from "howler";
import { intervalToDuration, formatDuration } from "date-fns";
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
    podcastProgress: 0,
    podcastVolume: 0.5,
    timeTrackerId: null,
  }),
  computed: {
    getPodcastItemTitle() {
      return this.podcastData ? this.podcastData.Title : "";
    },
    getRelativeProgress() {
      if (!this.podcastData) return 0;
      return (this.podcastProgress / this.podcastData.Duration) * 100;
    },
    timePlayed() {
      if (!this.howlerData) return "00:00";
      const duration = intervalToDuration({
        start: 0,
        end: this.podcastProgress * 1000,
      });
      const zeroPad = (num) => String(num).padStart(2, "0");

      return formatDuration(duration, {
        format: ["minutes", "seconds"],
        // format: ["hours", "minutes", "seconds"],
        zero: true,
        delimiter: ":",
        locale: {
          formatDistance: (_token, count) => zeroPad(count),
        },
      });
    },
    totalTime() {
      if (!this.podcastData) return 0;

      const duration = intervalToDuration({
        start: 0,
        end: this.podcastData.Duration * 1000,
      });

      const zeroPad = (num) => String(num).padStart(2, "0");

      return formatDuration(duration, {
        format: ["minutes", "seconds"],
        // format: ["hours", "minutes", "seconds"],
        zero: true,
        delimiter: ":",
        locale: {
          formatDistance: (_token, count) => zeroPad(count),
        },
      });
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
          volume: this.podcastVolume,
          loop: false,
          html5: true,
          preload: true,
        };
        this.howlerData = new Howl(howlerConfig);
        this.timeTrackerId = setInterval(() => {
          this.podcastProgress = this.howlerData.seek();
        }, 1000);
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
    setTrackVolume() {
      this.howlerData.volume(this.podcastVolume);
    },
  },
};
</script>
