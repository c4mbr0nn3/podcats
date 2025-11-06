<script setup>
import { useTitle } from '@vueuse/core'
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import PodcastPlayerCard from '@/components/PodcastPlayerCard.vue'
import { PodcastItemService } from '@/services'

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
