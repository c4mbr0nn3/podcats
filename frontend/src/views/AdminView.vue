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
import { UsersService } from "@/api";
import { ref } from "vue";
import { useQuery } from "@tanstack/vue-query";
import DialogUserEdit from "@/components/admin/DialogUserEdit.vue";
import SingleUserCard from "@/components/admin/SingleUserCard.vue";

const dialog = ref(false);
const dialogUser = ref({});
const isEdit = ref(false);
const alert = ref(false);
const alertTitle = ref("");
const alertText = ref("");

const fetchUsers = async () => {
  return await UsersService.getAll().then((response) => response.data);
};

const { data: users } = useQuery({
  queryKey: ["users"],
  queryFn: fetchUsers,
});

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
