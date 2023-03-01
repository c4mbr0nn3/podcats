<template>
  <v-card-text
    ><Markdown class="overflow-hidden fade" :source="summary" />
    <div class="d-flex align-center mt-2">
      <slot name="actions"></slot>
    </div>
  </v-card-text>
</template>

<script>
import Markdown from "vue3-markdown-it";

export default {
  components: {
    Markdown,
  },
  props: {
    summary: { type: String, default: "" },
  },
};
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
