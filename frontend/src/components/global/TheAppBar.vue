<template>
  <v-app-bar flat color="info">
    <v-row justify="center"
      ><v-btn
        v-for="(item, index) in menuList"
        :key="index"
        :to="item.route"
        variant="text"
        >{{ item.label }}
      </v-btn></v-row
    >
  </v-app-bar>
</template>

<script>
export default {
  data: () => ({
    menuList: [
      { label: "Home", route: { name: "home" } },
      { label: "Podcasts", route: { name: "podcasts" } },
    ],
  }),
};
</script>
