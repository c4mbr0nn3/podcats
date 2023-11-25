import { PodcastItemApi } from '@/api'
import { showErrorSnackbar, showSuccessSnackbar } from '@/utils/snackbar'

export async function getLatest(pageId) {
  const { data } = await PodcastItemApi.getLatest(pageId)
  return data
}

export async function getFavourites(pageId) {
  const { data } = await PodcastItemApi.getFavourites(pageId)
  return data
}

export async function getById(itemId) {
  const { data } = await PodcastItemApi.getById(itemId)
  return data
}

export async function updatePlayedTime(itemId, timePlayed) {
  const { data } = await PodcastItemApi.updatePlayedTime(itemId, timePlayed)
  return data
}

export async function setCompleted(itemId) {
  const { data } = await PodcastItemApi.setCompleted(itemId)
  return data
}

export async function switchStatus(itemId, isPlayed) {
  let data = null
  try {
    const { data: responseData } = await PodcastItemApi.switchStatus(itemId)
    data = responseData
    showSuccessSnackbar(
      isPlayed ? 'Podcast marked as not played' : 'Podcast marked as played',
    )
  }
  catch {
    showErrorSnackbar('Error switching podcast status')
  }
  return data
}

export async function switchFavourite(itemId, isFavourite) {
  let data = null
  try {
    const { data: responseData } = await PodcastItemApi.switchFavourite(itemId)
    data = responseData
    showSuccessSnackbar(
      isFavourite
        ? 'Podcast removed from favourites'
        : 'Podcast added to favourites',
    )
  }
  catch {
    showErrorSnackbar('Error switching podcast favourite status')
  }
  return data
}
