import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { videoApi } from '@/api/index'
import type { VideoData, Filters } from '@/models'

export const useVideoStore = defineStore('video', () => {
  const videos = ref<VideoData[]>([])

  const loadVideos = async () => {
    try {
      const resp = await videoApi.getAllVideos()
      videos.value = resp.data
    } catch (err) {
      console.error(err)
    }
  }

  return {
    videos,
    loadVideos,
  }
})
