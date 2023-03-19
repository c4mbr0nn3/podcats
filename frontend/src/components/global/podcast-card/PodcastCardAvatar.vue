<template>
  <v-hover v-slot="{ isHovering, props }" :disabled="!routerLink"
    ><component :is="routerLink ? 'router-link' : 'span'" :to="routerLink">
      <v-avatar
        class="ma-3"
        :size="avatarSize"
        rounded="lg"
        :class="{ 'on-hover': isHovering || !routerLink }"
        v-bind="props"
      >
        <v-img :src="image || missingImage">
          <template #placeholder>
            <div class="d-flex align-center justify-center fill-height">
              <v-progress-circular
                indeterminate
                color="primary"
              ></v-progress-circular></div
          ></template>
          <div class="fill-height d-flex justify-center align-center">
            <v-icon
              color="transparent"
              size="x-large"
              :class="{ 'show-btns': isHovering }"
              >mdi-play-circle</v-icon
            >
          </div></v-img
        >
      </v-avatar></component
    >
  </v-hover>
</template>

<script>
import missingImage from "@/assets/missing_image.png";
export default {
  props: {
    image: { type: String, default: "" },
    routerLink: { type: Object, default: null },
    avatarSize: { type: Number, default: 125 },
  },
  data: () => ({
    missingImage: missingImage,
  }),
};
</script>

<style scoped>
.v-avatar {
  transition: opacity 0.4s ease-in-out;
}
.v-avatar:not(.on-hover) {
  opacity: 0.6;
}

.show-btns {
  color: rgba(var(--v-theme-primary), 1) !important;
}
</style>
