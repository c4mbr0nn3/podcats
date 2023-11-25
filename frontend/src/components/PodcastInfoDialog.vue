<script setup>
import { computed } from 'vue'
import PMarkdown from '@/components/global/PMarkdown.vue'
import { formatDateToIso } from '@/utils/date'
import { podcastInfoDialogSchema } from '@/schemas/podcast-info-dialog'

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true,
  },
  infoDialogData: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['update:modelValue'])

const dialog = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val),
})

function getItemText(item) {
  if (item.type === 'date')
    return formatDateToIso(props.infoDialogData[item.key])

  return props.infoDialogData[item.key]
}
</script>

<template>
  <v-dialog v-model="dialog" height="700" width="600" scrollable>
    <v-card>
      <v-card-title class="bg-primary d-flex align-center">
        <div class="text-truncate" style="width: 85%">
          {{ props.infoDialogData.Title }}
        </div>
        <v-spacer />
        <v-btn
          icon="mdi-close"
          density="compact"
          variant="text"
          @click="dialog = false"
        />
      </v-card-title>
      <v-card-text>
        <div v-for="(item, index) in podcastInfoDialogSchema" :key="index">
          <div class="text-h6 text-primary">
            {{ item.label }}
          </div>
          <PMarkdown :markdown="getItemText(item)" />
          <v-divider class="my-1" />
        </div>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>
