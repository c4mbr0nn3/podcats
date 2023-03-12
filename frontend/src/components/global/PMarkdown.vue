<template>
  <div v-html="html"></div>
</template>

<script setup>
import { markdownToHtml } from "@/plugins/marked";

const props = defineProps({
  markdown: {
    type: [String, Number],
    required: true,
  },
});

const html = markdownToHtml(props.markdown);
</script>
