<script setup>
import { computed } from 'vue'
import PMarkdown from '@/components/global/PMarkdown.vue'

const props = defineProps({
  summary: {
    type: String,
    required: true,
  },
  applyFade: {
    type: Boolean,
    default: true,
  },
})

const getSummaryStyle = computed(() => {
  return {
    'overflow-hidden': props.applyFade,
    'fade': props.applyFade,
  }
})

const shortenedSummary = computed(() => {
  if (props.applyFade && props.summary.length > 250)
    return `${props.summary.substring(0, 250)}...`
  return props.summary
})
</script>

<template>
  <v-card-text>
    <PMarkdown :class="getSummaryStyle" :markdown="shortenedSummary" />
    <div class="d-flex align-center mt-2">
      <slot name="actions" />
    </div>
  </v-card-text>
</template>

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
