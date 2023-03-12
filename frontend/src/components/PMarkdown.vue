<template>
  <div v-html="html"></div>
</template>

<script setup>
import { defineProps } from "vue";
import { markdownToHtml } from "@/plugins/marked";

const props = defineProps({
  markdown: {
    type: String,
    required: true,
  },
});

const html = markdownToHtml(props.markdown);
</script>
