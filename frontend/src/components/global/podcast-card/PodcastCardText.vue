<template>
  <v-card-text
    ><p-markdown :class="getSummaryStyle" :markdown="props.summary" />
    <div class="d-flex align-center mt-2">
      <slot name="actions"></slot>
    </div>
  </v-card-text>
</template>

<script setup>
import { computed } from "vue";
import PMarkdown from "@/components/global/PMarkdown.vue";

const props = defineProps({
  summary: {
    type: String,
    required: true,
  },
  applyFade: {
    type: Boolean,
    default: true,
  },
});

const getSummaryStyle = computed(() => {
  return {
    "overflow-hidden": props.applyFade,
    fade: props.applyFade,
  };
});
</script>

<style scoped>
/* https://css-tricks.com/line-clampin/ */
.fade {
  position: relative;
  height: 6em; /* exactly four lines with line height = 1.5*/
}
.fade:after {
  content: "";
  text-align: right;
  position: absolute;
  bottom: 0;
  right: 0;
  width: 70%;
  height: 1.5em;
  background: linear-gradient(
    to right,
    rgba(var(--v-theme-surface), 0),
    rgba(var(--v-theme-surface), 1) 50%
  );
}
</style>
