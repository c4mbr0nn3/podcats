<template>
  <div>
    <v-row v-if="!isTrackLoaded">
      <v-progress-linear indeterminate color="primary"></v-progress-linear>
    </v-row>
    <div v-else>
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
    </div>
  </div>
</template>

<script>
import { updatePodcastItemPlayedTime, setPodcastItemCompleted } from "@/api";
import { Howl, Howler } from "howler";
import { getTotalTime, getTimePlayed } from "@/utils/player";
export default {
  props: {
    data: { type: Object, default: () => {} },
  },
  data() {
    return {
      podcastData: this.data,
      isTrackLoaded: false,
      howlerData: null,
      isPlaying: false,
      isMuted: false,
      podcastProgress: 0,
      podcastVolume: 0.5,
      timeTrackerId: null,
      trackSpeed: 1,
      availableSpeed: [0.8, 1, 1.2, 1.8, 2.5, 3.5],
    };
  },
  computed: {
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
      return getTimePlayed(this.howlerData, this.podcastProgress);
    },
    totalTime() {
      return getTotalTime(this.podcastData);
    },
  },
  created() {
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
        this.isPlaying = false;
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
  },
  async unmounted() {
    await this.updateTrackProgress();
    Howler.unload();
  },
  methods: {
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

<style></style>
