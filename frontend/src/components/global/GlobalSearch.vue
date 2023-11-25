<script setup>
import { Icon } from '@iconify/vue'
import { ref } from 'vue'
import { onKeyStroke, useDebounceFn } from '@vueuse/core'
import Fuse from 'fuse.js'
import { usePodcastsStore } from '@/stores/podcasts'
import SearchResultCard from '@/components/SearchResultCard.vue'

const dialog = ref(false)
const searchString = ref('')
const searchResults = ref([])
const searchLoading = ref(false)

const podcastsStore = usePodcastsStore()

onKeyStroke(['ctrl', 'k'], (e) => {
  e.preventDefault()
  if (!dialog.value && e.ctrlKey && e.key === 'k')
    dialog.value = true
})

const debounceSearch = useDebounceFn(() => {
  const fuse = new Fuse(podcastsStore.podcasts, {
    keys: ['Title'],
    threshold: 0.3,
  })
  searchResults.value = fuse.search(searchString.value).map(r => r.item)
  searchLoading.value = false
}, 1000)

function clearData() {
  clearSearchResults()
  searchLoading.value = false
  searchString.value = ''
}

function clearSearchResults() {
  searchResults.value = []
}

function onInputChange() {
  if (searchString.value.length === 0)
    return
  clearSearchResults()
  searchLoading.value = true
  debounceSearch()
}
</script>

<template>
  <v-dialog
    v-model="dialog"
    height="700"
    width="600"
    scrollable
    @after-leave="clearData"
  >
    <template #activator="{ props }">
      <v-btn
        variant="plain"
        color="secondary"
        class="text-capitalize font-weight-regular ml-2 bg-surface-lighten-1"
        v-bind="props"
      >
        <Icon icon="carbon:search" class="mr-2" />
        Search
        <v-chip size="x-small" variant="outlined" label class="ml-4">
          Ctrl K
        </v-chip>
      </v-btn>
    </template>
    <v-row no-gutters class="mb-2">
      <v-col cols="12">
        <v-card color="primary" flat rounded>
          <v-card-title>
            <v-text-field
              v-model="searchString"
              placeholder="Start typing..."
              autofocus
              hide-details
              single-line
              prepend-inner-icon="mdi-magnify"
              bg-color="transparent"
              color="primary"
              variant="outlined"
              @input="onInputChange"
            />
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>
    <v-card height="100%">
      <v-card-text>
        <v-row v-if="searchLoading" class="fill-height">
          <v-col cols="12" align-self="center" class="d-flex justify-center">
            <Icon
              icon="svg-spinners:3-dots-fade"
              color="secondary"
              width="50"
            />
          </v-col>
        </v-row>
        <v-row v-else-if="searchResults.length > 0">
          <v-col cols="12">
            <SearchResultCard
              v-for="(result, index) in searchResults"
              :key="index"
              :podcast="result"
              @click="dialog = false"
            />
          </v-col>
        </v-row>
        <v-row v-else-if="searchString.length > 0" class="fill-height">
          <v-col
            cols="12"
            align-self="center"
            class="d-flex flex-column align-center text-disabled"
          >
            <Icon icon="mdi:emoticon-dead-outline" width="100" />
            <div class="mt-1">
              No results found
            </div>
          </v-col>
        </v-row>
        <v-row v-else class="fill-height">
          <v-col
            cols="12"
            align-self="center"
            class="d-flex flex-column align-center text-disabled"
          >
            <Icon icon="mdi:selection-search" width="100" />
            <div class="mt-1">
              Your search result will appear here
            </div>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>
