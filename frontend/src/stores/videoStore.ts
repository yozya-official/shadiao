import { videoApi } from '@/api/index'
import { Video, type Filters } from '@/models'

export const useVideoStore = defineStore('video', () => {
  const videos = ref<Video[]>([])

  const tags = ref<Tag[]>([])

  const loadVideos = async () => {
    try {
      const resp = await videoApi.getAllVideos()
      videos.value = resp.data.videos.map((v: Partial<Video>) => new Video(v))
    } catch (err) {
      console.error(err)
    }
  }

  const backgroundOptions = computed(() => {
    return tags.value.filter((tag) => {
      return tag.type === 'background'
    })
  })
  const styleOptions = computed(() => {
    return tags.value.filter((tag) => {
      return tag.type === 'style'
    })
  })
  const worldOptions = computed(() => {
    return tags.value.filter((tag) => {
      return tag.type === 'world'
    })
  })

  const loadTags = async () => {
    try {
      const resp = await tagApi.getAllTags()
      tags.value = resp.data.tags
    } catch (err) {
      console.error(err)
    }
  }

  return {
    videos,
    loadVideos,
    tags,
    loadTags,
    backgroundOptions,
    styleOptions,
    worldOptions,
  }
})
