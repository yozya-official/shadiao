import { videoApi } from '@/api/index'
import { Video, type Filters } from '@/models'

export const useVideoStore = defineStore('video', () => {
  const videos = ref<Video[]>([])

  const loadVideos = async () => {
    try {
      const resp = await videoApi.getAllVideos()
      videos.value = resp.data.videos.map((v: Partial<Video>) => new Video(v))
    } catch (err) {
      console.error(err)
    }
  }

  return {
    videos,
    loadVideos,
  }
})
