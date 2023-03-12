<template>
  <v-dialog width="700">
    <v-card>
      <v-card-text>
        <div class="text-h6 font-italic text-primary">Title</div>
        <p-markdown :markdown="infoDialogData.Title" />
        <v-divider class="my-1"></v-divider>
        <div class="text-h6 font-italic text-primary">Author</div>
        <p-markdown :markdown="infoDialogData.Author" />
        <v-divider class="my-1"></v-divider>
        <div class="text-h6 font-italic text-primary">Summary</div>
        <p-markdown :markdown="infoDialogData.Summary" />
        <v-divider class="my-1"></v-divider>
        <div class="text-h6 font-italic text-primary">Show URL</div>
        <p-markdown :markdown="infoDialogData.ShowURL" />
        <v-divider class="my-1"></v-divider>
        <div class="text-h6 font-italic text-primary">How many episodes?</div>
        <p-markdown :markdown="infoDialogData.EpisodesCount" />
        <v-divider class="my-1"></v-divider>
        <div class="text-h6 font-italic text-primary">How many played?</div>
        <p-markdown :markdown="infoDialogData.PlayedCount" />
        <v-divider class="my-1"></v-divider>
        <div class="text-h6 font-italic text-primary">Imported At</div>
        <p-markdown :markdown="formatDateToIso(infoDialogData.CreatedAt)" />
        <v-divider class="my-1"></v-divider>
      </v-card-text>
      <v-card-actions>
        <v-btn color="primary" block @click="$emit('close-info-dialog')"
          >Close</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import PMarkdown from "@/components/global/PMarkdown.vue";
import { formatDateToIso } from "@/utils/date";

defineEmits(["close-info-dialog"]);

defineProps({
  infoDialogData: {
    type: Object,
    required: true,
  },
});
</script>
