<script setup>
import { markdownToHtml } from '@/plugins/marked'

const props = defineProps({
  markdown: {
    type: [String, Number],
    required: true,
  },
})

const html = markdownToHtml(props.markdown)
</script>

<template>
  <div v-html="html" />
</template>
