<script setup>
import PodcastAvatar from '@/components/global/podcast-card/PodcastCardAvatar.vue'
import PodcastCardText from '@/components/global/podcast-card/PodcastCardText.vue'

const props = defineProps({
  podcast: {
    type: Object,
    default: () => {},
  },
  routerLink: {
    type: Object,
    default: null,
  },
  applyFade: {
    type: Boolean,
    default: true,
  },
  avatarSize: {
    type: Number,
    default: 125,
  },
})
</script>

<template>
  <v-card class="mt-3">
    <div class="d-flex flex-no-wrap justify-space-between">
      <v-col cols="10">
        <v-card-title class="text-truncate">
          {{
            props.podcast.Title
          }}
        </v-card-title>
        <PodcastCardText
          :summary="props.podcast.Summary"
          :apply-fade="props.applyFade"
        >
          <template #actions>
            <slot name="actions" />
          </template>
        </PodcastCardText>
      </v-col>
      <v-col cols="2" class="d-flex justify-end ml-2">
        <PodcastAvatar
          :image="props.podcast.Image"
          :router-link="props.routerLink"
          :avatar-size="props.avatarSize"
        />
      </v-col>
    </div>
    <slot name="other" />
  </v-card>
</template>
