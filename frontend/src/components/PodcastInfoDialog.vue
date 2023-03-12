<template>
  <v-dialog width="700">
    <v-card>
      <v-card-text>
        <div v-for="(item, index) in podcastInfoDialogSchema" :key="index">
          <div class="text-h6 font-italic text-primary">{{ item.label }}</div>
          <p-markdown :markdown="getItemText(item)" />
          <v-divider class="my-1"></v-divider>
        </div>
      </v-card-text>
      <v-card-actions>
        <v-btn color="secondary" block @click="$emit('close-info-dialog')"
          >Close</v-btn
        >
      </v-card-actions>
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
