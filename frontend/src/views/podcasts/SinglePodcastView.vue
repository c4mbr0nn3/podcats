<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useTitle } from '@vueuse/core'
import { PodcastService } from '@/services'
import SinglePodcastItemCard from '@/components/SinglePodcastItemCard.vue'

const podcastData = ref(null)
const podcastItemsData = ref([])
const currentPage = ref(1)
const nextPage = ref(null)
const pageCount = ref(null)

const route = useRoute()

const cardTitle = computed(() => {
  return podcastData.value ? podcastData.value.Title : ''
})

const podcastItems = computed(() => {
  return podcastItemsData.value ? podcastItemsData.value : []
})

const title = computed(() => {
  return podcastData.value ? `PodCats | ${cardTitle.value}` : 'PodCats'
})

useTitle(title)

onMounted(async () => {
  await fetchData()
})

async function fetchData() {
  podcastData.value = await PodcastService.getById(route.params.id)
  await fetchPodcastItems(1)
}

async function fetchPodcastItems(pageId) {
  const data = await PodcastService.getItemsById(route.params.id, pageId)
  if (data.podcastItems.length > 0)
    data.podcastItems.forEach(item => podcastItemsData.value.push(item))

  currentPage.value = data.thisPage
  pageCount.value = data.pageCount
  nextPage.value = data.nextPage
}

async function onIntersect() {
  if (currentPage.value === pageCount.value || !nextPage.value)
    return
  await fetchPodcastItems(nextPage.value)
}

function changePlayedStatus(event) {
  const item = podcastItemsData.value.find(el => el.ID === event)
  item.IsPlayed = !item.IsPlayed
}

function changeFavStatus(event) {
  const item = podcastItemsData.value.find(el => el.ID === event.id)
  item.BookmarkDate = event.bookmarkDate
}
</script>

<template>
  <v-row justify="center">
    <v-col cols="9">
      <v-card>
        <v-card-title>
          {{ cardTitle }}
        </v-card-title>
        <v-card-text>
          <SinglePodcastItemCard
            v-for="(podcastItem, index) in podcastItems"
            :key="index"
            class="mt-3"
            :podcast-id="$route.params.id"
            :podcast-item="podcastItem"
            @change-fav-status="changeFavStatus"
            @change-played-status="changePlayedStatus"
          />
        </v-card-text>
      </v-card>
      <v-card v-intersect="onIntersect" />
    </v-col>
  </v-row>
</template>
