<template>
  <v-dialog height="700" width="600" scrollable>
    <v-card>
      <v-card-title class="bg-primary d-flex"
        >{{ props.infoDialogData.Title }}<v-spacer />
        <v-btn
          icon="mdi-close"
          size="x-small"
          variant="text"
          @click="$emit('close-info-dialog')"
      /></v-card-title>
      <v-card-text>
        <div v-for="(item, index) in podcastInfoDialogSchema" :key="index">
          <div class="text-h6 text-primary">{{ item.label }}</div>
          <p-markdown :markdown="getItemText(item)" />
          <v-divider class="my-1"></v-divider>
        </div>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup>
import PMarkdown from "@/components/global/PMarkdown.vue";
import { formatDateToIso } from "@/utils/date";
import { podcastInfoDialogSchema } from "@/schemas/podcast-info-dialog";

defineEmits(["close-info-dialog"]);

const props = defineProps({
  infoDialogData: {
    type: Object,
    required: true,
  },
});

const getItemText = (item) => {
  if (item.type === "date") {
    return formatDateToIso(props.infoDialogData[item.key]);
  }
  return props.infoDialogData[item.key];
};
</script>
