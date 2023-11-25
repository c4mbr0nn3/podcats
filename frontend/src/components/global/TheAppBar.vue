<script setup>
import { onMounted } from 'vue'
import GlobalSearch from './GlobalSearch.vue'
import ImportBtn from '@/components/app-bar/AppBarImportBtn.vue'
import UserMenuBtn from '@/components/app-bar/AppBarUserMenuBtn.vue'
import NotificationBtn from '@/components/app-bar/AppBarNotificationBtn.vue'
import { usePodcastsStore } from '@/stores/podcasts'

const podcastsStore = usePodcastsStore()
const { fetchPodcasts } = podcastsStore

const menuList = [
  { label: 'Home', route: { name: 'home' } },
  { label: 'My Podcasts', route: { name: 'podcasts' } },
  { label: 'Latest', route: { name: 'latest-podcasts' } },
  {
    label: 'Favs',
    icon: 'mdi-heart',
    route: { name: 'favorites-podcasts' },
  },
]

onMounted(() => {
  fetchPodcasts()
})
</script>

<template>
  <v-app-bar flat>
    <v-row>
      <v-col cols="3">
        <GlobalSearch class="ml-6" />
      </v-col>
      <v-col cols="6" class="d-flex justify-center">
        <v-btn
          v-for="(item, index) in menuList"
          :key="index"
          :to="item.route"
          variant="text"
          color="primary"
          :prepend-icon="item.icon ? item.icon : null"
          class="mx-1"
        >
          {{ item.label }}
        </v-btn>
      </v-col>
      <v-col cols="3" class="d-flex justify-end">
        <ImportBtn />
        <NotificationBtn />
        <UserMenuBtn />
      </v-col>
    </v-row>
  </v-app-bar>
</template>
