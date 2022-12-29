<template>
  <v-app-bar flat>
    <v-row justify="center"
      ><v-btn
        v-for="(item, index) in menuList"
        :key="index"
        :to="item.route"
        variant="text"
        color="primary"
        :prepend-icon="item.icon ? item.icon : null"
        class="mx-1"
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
      { label: "Latest", route: { name: "latest-podcasts" } },
      {
        label: "Favs",
        icon: "mdi-heart",
        route: { name: "favourites-podcasts" },
      },
    ],
  }),
};
</script>
