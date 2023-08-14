<template>
  <v-dialog v-model="dialog" width="500">
    <v-card>
      <v-card-title>{{ dialogTitle }}</v-card-title>
      <v-card-text>
        <v-form ref="manage_user_form" v-model="formValidation">
          <v-text-field
            v-model="userDialog.Username"
            label="Username"
            color="primary"
            :disabled="props.isEdit"
            :rules="[isRequiredRule]"
          ></v-text-field>
          <v-text-field
            v-model="userDialog.Name"
            label="Name"
            color="primary"
            :rules="[isRequiredRule, isAlphaWithSpacesRule]"
          ></v-text-field>
          <v-text-field
            v-model="userDialog.Surname"
            label="Surname"
            color="primary"
            :rules="[isRequiredRule, isAlphaWithSpacesRule]"
          ></v-text-field>
          <v-text-field
            v-model="userDialog.Email"
            label="E-mail"
            color="primary"
            :disabled="props.isEdit"
            :rules="[isRequiredRule, isEmailRule]"
          ></v-text-field>
          <v-checkbox
            v-model="userDialog.IsAdmin"
            label="Admin"
            color="primary"
          ></v-checkbox>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="secondary" variant="plain" @click="dialog = false">
          Close
        </v-btn>
        <v-btn color="success" :disabled="!formValidation" @click="onSave">
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed, ref } from "vue";
import {
  isEmailRule,
  isRequiredRule,
  isAlphaWithSpacesRule,
} from "@/utils/validation";
import { watch } from "vue";

const emit = defineEmits(["update:modelValue", "save-user"]);

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true,
  },
  user: {
    type: Object,
    required: true,
  },
  isEdit: {
    type: Boolean,
    required: true,
  },
});

const manage_user_form = ref(null);
const formValidation = ref(false);
const userDialog = ref(props.user);

watch(
  () => props.modelValue,
  () => {
    userDialog.value = {};
    // assign all key from props.user to userDialog
    Object.keys(props.user).forEach((key) => {
      userDialog.value[key] = props.user[key];
    });
  }
);

const dialog = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});

const dialogTitle = computed(() => {
  if (props.isEdit) {
    return "Edit User";
  }
  return "Add User";
});

const onSave = () => {
  emit("save-user", userDialog.value);
  dialog.value = false;
};
</script>
