<template>
  <v-card color="background" rounded="lg">
    <v-card-title class="d-flex align-center">
      {{ user.Name }} {{ user.Surname }}
      <v-icon v-if="user.IsAdmin" size="xsmall" color="primary" class="ml-2"
        >mdi-shield-crown-outline</v-icon
      >
      <v-spacer></v-spacer>
      <v-btn
        variant="text"
        density="comfortable"
        color="primary"
        icon="mdi-pencil"
        @click="$emit('edit', user)"
      />
      <v-btn
        variant="text"
        density="comfortable"
        color="info"
        icon="mdi-lock-reset"
        class="ml-1"
        @click="$emit('reset-password', user)"
      />
      <v-btn
        variant="text"
        density="comfortable"
        color="error"
        icon="mdi-delete"
        class="ml-1"
        :disabled="user.ID === 1"
        @click="$emit('delete', user)"
      />
    </v-card-title>
    <v-card-text>
      <v-row>
        <v-col cols="3"> ID: {{ user.ID }} </v-col>
        <v-col cols="3"> Username: {{ user.Username }} </v-col>
        <v-col cols="3"> E-mail: {{ user.Email }} </v-col>
        <v-col cols="3">
          Created At:
          {{ formatDateToIso(user.CreatedAt.split(" ")[0]) }}
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script setup>
import { formatDateToIso } from "@/utils/date";

defineEmits(["edit", "reset-password", "delete"]);

defineProps({
  user: {
    type: Object,
    required: true,
  },
});
</script>
