<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useTitle } from '@vueuse/core'
import { PodcastItemService } from '@/services'
import PodcastPlayerCard from '@/components/PodcastPlayerCard.vue'

const item = ref(null)

const route = useRoute()

const title = computed(() => {
  return item.value ? `PodCats | ${item.value.Title}` : 'PodCats'
})

useTitle(title)

onMounted(async () => {
  item.value = await PodcastItemService.getById(route.params.itemId)
})
</script>

<template>
  <PodcastPlayerCard v-if="item" :podcast-data="item" />
</template>
