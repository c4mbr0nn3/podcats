<template>
  <v-row align="center" justify="center">
    <v-col cols="8">
      <v-card v-if="podcastData" :loading="!isTrackLoaded">
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
            <v-avatar class="ma-3" size="125" rounded="0">
              <v-img
                :src="podcastData.Image ? podcastData.Image : missingImage"
              >
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
        <div v-if="isTrackLoaded">
          <v-col cols="12">
            <!--qui mettiamo tutti i controlli del player da spostare in component-->
            <v-card-text>
              <v-row no-gutters
                >{{ timePlayed }} <v-spacer></v-spacer> {{ totalTime }}</v-row
              >
              <v-row no-gutters class="mt-1">
                <v-slider
                  v-model="podcastProgress"
                  min="0"
                  :max="getPodcastDuration"
                  color="primary"
                  thumb-size="15"
                  hide-details
                  @update:model-value="onPodcastProgressUpdate()"
                  ><template #prepend>
                    <v-icon
                      icon="mdi-timer-outline"
                      color="primary"
                    ></v-icon> </template></v-slider
              ></v-row>
              <v-row no-gutters justify="center" class="mt-2">
                <v-col cols="3" class="d-flex justify-center align-center">
                  <v-menu location="bottom">
                    <template #activator="{ props }">
                      <v-btn
                        variant="text"
                        prepend-icon="mdi-speedometer"
                        v-bind="props"
                        color="primary"
                        >{{ getTrackSpeed }}</v-btn
                      >
                    </template>
                    <v-list>
                      <v-list-item
                        v-for="(item, index) in availableSpeed"
                        :key="index"
                        @click="setTrackSpeed(item)"
                      >
                        <v-list-item-title>{{ item }}</v-list-item-title>
                      </v-list-item>
                    </v-list>
                  </v-menu>
                </v-col>
                <v-col cols="6" class="d-flex justify-center align-center">
                  <v-btn
                    icon="mdi-rewind-10"
                    color="secondary"
                    class="mx-1"
                    variant="text"
                    @click="moveTrackSeek(-10)"
                  ></v-btn>
                  <v-btn
                    :icon="isPlaying ? 'mdi-pause' : 'mdi-play'"
                    :color="isPlaying ? 'info' : 'primary'"
                    size="large"
                    class="mx-1"
                    @click="switchTrackStatus"
                  ></v-btn
                  ><v-btn
                    icon="mdi-stop"
                    color="secondary"
                    class="mx-1"
                    size="large"
                    variant="outlined"
                    @click="stopTrack"
                  ></v-btn>
                  <v-btn
                    icon="mdi-fast-forward-10"
                    color="secondary"
                    class="mx-1"
                    variant="text"
                    @click="moveTrackSeek(10)"
                  ></v-btn>
                </v-col>
                <v-col cols="3" class="d-flex justify-center align-center">
                  <v-slider
                    v-model="podcastVolume"
                    min="0"
                    max="1"
                    :step="0.01"
                    color="primary"
                    thumb-size="15"
                    track-size="2"
                    hide-details
                    @update:model-value="setTrackVolume()"
                  >
                    <template #prepend>
                      <v-btn
                        variant="text"
                        :icon="getVolumeIcon"
                        color="primary"
                        @click="muteTrack"
                      ></v-btn> </template
                  ></v-slider>
                </v-col>
              </v-row>
            </v-card-text>
          </v-col>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import missingImage from "@/assets/missing_image.png";
import {
  getPodcastItemById,
  updatePodcastItemPlayedTime,
  setPodcastItemCompleted,
} from "@/api";
import { Howl, Howler } from "howler";
import { intervalToDuration, formatDuration } from "date-fns";
import Markdown from "vue3-markdown-it";
export default {
  components: {
    Markdown,
  },
  data: () => ({
    isTrackLoaded: false,
    podcastData: null,
    howlerData: null,
    isPlaying: false,
    isMuted: false,
    podcastProgress: 0,
    podcastVolume: 0.5,
    timeTrackerId: null,
    trackSpeed: 1,
    availableSpeed: [0.8, 1, 1.2, 1.8, 2.5, 3.5],
    missingImage: missingImage,
  }),
  computed: {
    getPodcastItemTitle() {
      return this.podcastData ? this.podcastData.Title : "";
    },
    getPodcastDuration() {
      if (!this.podcastData) return 0;
      return this.podcastData.Duration;
    },
    getVolumeIcon() {
      return this.isMuted ? "mdi-volume-mute" : "mdi-volume-high";
    },
    getTrackSpeed() {
      return `${this.trackSpeed.toFixed(1)}x`;
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
  async unmounted() {
    await this.updateTrackProgress();
    Howler.unload();
  },
  methods: {
    async fetchData() {
      await getPodcastItemById(
        this.$route.params.id,
        this.$route.params.itemId
      ).then((response) => {
        this.podcastData = response.data;
        this.podcastProgress = this.podcastData.LastPlayPosition;
        let howlerConfig = {
          src: [this.podcastData.FileURL],
          volume: this.podcastVolume,
          loop: false,
          html5: true,
          preload: true,
          onload: () => {
            this.podcastData.Duration = this.howlerData.duration();
            this.isTrackLoaded = true;
          },
          onend: async () => {
            await this.updateTrackProgress();
            await setPodcastItemCompleted(
              this.podcastData.PodcastID,
              this.podcastData.ID
            );
          },
        };
        this.howlerData = new Howl(howlerConfig);
        this.setTrackSeek();
        this.timeTrackerId = setInterval(() => {
          this.podcastProgress = this.howlerData.seek();
        }, 1000);
      });
    },
    async updateTrackProgress() {
      await updatePodcastItemPlayedTime(
        this.podcastData.PodcastID,
        this.podcastData.ID,
        Math.floor(this.podcastProgress)
      );
    },
    onPodcastProgressUpdate() {
      this.setTrackSeek();
      this.updateTrackProgress();
    },
    switchTrackStatus() {
      if (this.isPlaying) {
        this.pauseTrack();
        this.updateTrackProgress();
      } else {
        this.playTrack();
      }
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
    setTrackSeek() {
      this.howlerData.seek(this.podcastProgress);
    },
    moveTrackSeek(timeChunk) {
      this.podcastProgress += timeChunk;
      this.setTrackSeek();
    },
    setTrackSpeed(value) {
      this.trackSpeed = value;
      this.howlerData.rate(value);
    },
  },
};
</script>
