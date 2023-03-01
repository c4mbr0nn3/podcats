<template>
  <v-card class="mt-3">
    <div class="d-flex flex-no-wrap justify-space-between">
      <v-col cols="10">
        <v-card-title class="text-truncate">{{ podcast.Title }}</v-card-title>
        <PodcastCardText :summary="podcast.Summary">
          <template #actions>
            <slot name="actions"></slot>
          </template>
        </PodcastCardText>
      </v-col>
      <v-col cols="2" class="d-flex justify-end ml-2">
        <PodcastAvatar
          :image="podcast.Image"
          :router-link="{ name: 'single-podcast', params: { id: podcast.ID } }"
        />
      </v-col></div
  ></v-card>
</template>

<script>
import PodcastAvatar from "@/components/global/PodcastAvatar.vue";
import PodcastCardText from "@/components/global/PodcastCardText.vue";

export default {
  components: {
    PodcastAvatar,
    PodcastCardText,
  },
  props: {
    podcast: { type: Object, default: () => {} },
    routerLink: { type: Object, default: null },
  },
};
</script>
