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
          <SingleUserCard
            v-for="(user, index) in users"
            :key="index"
            :user="user"
            class="mb-2"
            @edit="editUser"
            @reset-password="resetPassword"
            @delete="deleteUser"
          />
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
import { UserService } from "@/services";
import { ref, onMounted } from "vue";
import DialogUserEdit from "@/components/admin/DialogUserEdit.vue";
import SingleUserCard from "@/components/admin/SingleUserCard.vue";

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
  users.value = await UserService.getAll();
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
  const data = await UserService.resetPassword(user.ID);
  alertTitle.value = data.message;
  alertText.value = data.password;
  alert.value = true;
};

const deleteUser = async (user) => {
  await UserService.remove(user.ID);
  fetchUsers();
};

const saveUser = async (user) => {
  const promise = isEdit.value
    ? UserService.update(user.ID, user)
    : UserService.create(user);

  const data = await promise;
  fetchUsers();
  if (isEdit.value) return;
  alertTitle.value = data.message;
  alertText.value = data.password;
  alert.value = true;
};
</script>
