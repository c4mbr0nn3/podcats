<template>
  <v-row justify="center">
    <v-col cols="12">
      <v-alert
        v-model="alert"
        variant="tonal"
        type="success"
        :title="alertTitle"
        closable
        class="mb-2"
      >
        <div>The user password is:</div>
        <div>{{ alertText }}</div>
        <div>
          It is the only time you will see it, please share it with the user!
        </div>
      </v-alert>
      <v-card>
        <v-card-title>
          Users List
          <v-btn
            color="primary"
            append-icon="mdi-plus"
            variant="plain"
            density="comfortable"
            @click="addUser"
          >
            Add User</v-btn
          >
        </v-card-title>
        <v-card-text>
          <v-card
            v-for="(user, index) in users"
            :key="index"
            color="background"
            rounded="lg"
            class="mb-2"
          >
            <v-card-title class="d-flex align-center">
              {{ user.Name }} {{ user.Surname }}
              <v-icon
                v-if="user.IsAdmin"
                size="xsmall"
                color="primary"
                class="ml-2"
                >mdi-shield-crown-outline</v-icon
              >
              <v-spacer></v-spacer>
              <v-btn
                variant="text"
                density="comfortable"
                color="primary"
                icon="mdi-pencil"
                @click="editUser(user)"
              />
              <v-btn
                variant="text"
                density="comfortable"
                color="info"
                icon="mdi-lock-reset"
                class="ml-1"
                @click="resetPassword(user)"
              />
              <v-btn
                variant="text"
                density="comfortable"
                color="error"
                icon="mdi-delete"
                class="ml-1"
                :disabled="user.ID === 1"
                @click="deleteUser(user)"
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
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
  <dialog-user-edit
    v-model:model-value="dialog"
    :user="dialogUser"
    :is-edit="isEdit"
    @save-user="saveUser"
  />
</template>

<script setup>
import { UsersService } from "@/api";
import { ref, onMounted } from "vue";
import { formatDateToIso } from "@/utils/date";
import DialogUserEdit from "@/components/admin/DialogUserEdit.vue";

const users = ref([]);
const dialog = ref(false);
const dialogUser = ref({});
const isEdit = ref(false);
const alert = ref(false);
const alertTitle = ref("");
const alertText = ref("");

onMounted(async () => {
  await fetchUsers();
});

const fetchUsers = async () => {
  await UsersService.getAll().then((response) => {
    users.value = response.data;
  });
};

const addUser = () => {
  isEdit.value = false;
  const user = {
    Username: "",
    Name: "",
    Surname: "",
    Email: "",
    IsAdmin: false,
  };
  dialogUser.value = user;
  dialog.value = true;
};

const editUser = (user) => {
  isEdit.value = true;
  dialogUser.value = user;
  dialog.value = true;
};

const resetPassword = async (user) => {
  await UsersService.resetPassword(user.ID).then((response) => {
    alertTitle.value = response.data.message;
    alertText.value = response.data.password;
    alert.value = true;
  });
};

const deleteUser = async (user) => {
  await UsersService.remove(user.ID).then(() => fetchUsers());
};

const saveUser = async (user) => {
  const promise = isEdit.value
    ? UsersService.update(user.ID, user)
    : UsersService.create(user);

  await promise.then((response) => {
    fetchUsers();
    if (isEdit.value) return;
    alertTitle.value = response.data.message;
    alertText.value = response.data.password;
    alert.value = true;
  });
};
</script>
