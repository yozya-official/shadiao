import axios from 'axios'
import { useAuthStore } from '@/stores/authStore'

// 创建 axios 实例
const api = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    // 加载Authorization
    const store = useAuthStore()
    const { apiKey } = store
    if (apiKey) {
      config.headers['Authorization'] = `ApiKey ${apiKey}`
    }
    console.log('Request:', config.method?.toUpperCase(), config.url)

    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    // 对响应数据做点什么
    return response
  },
  (error) => {
    // 对响应错误做点什么
    console.error('API Error:', error.response?.data || error.message)
    return Promise.reject(error)
  },
)

// Video API 接口
export const videoApi = {
  /**
   * 创建视频数据
   * @param {Object} videoData - 视频数据对象
   * @returns {Promise}
   */
  createVideo: (videoData: VideoNewData) => {
    return api.post('/videos', videoData)
  },

  /**
   * 获取所有视频数据
   * @param {Object} params - 查询参数（可选）
   * @returns {Promise}
   */
  getAllVideos: (params = {}) => {
    return api.get('/videos', { params })
  },

  /**
   * 根据ID获取视频数据
   * @param {string|number} id - 视频ID
   * @returns {Promise}
   */
  getVideoById: (id: string) => {
    return api.get(`/videos/${id}`)
  },

  /**
   * 更新视频数据
   * @param {string|number} id - 视频ID
   * @param {Object} videoData - 更新的视频数据
   * @returns {Promise}
   */
  updateVideo: (id: string, videoData: VideoData) => {
    return api.put(`/videos/${id}`, videoData)
  },

  /**
   * @returns {Promise}
   */
  getUnreviewedVideos: () => {
    return api.get('/videos/unreviewed')
  },

  /**
   * 修改视频审核状态（修改Public字段）
   * @param {number|string} id - 视频ID
   * @param {boolean} isReviewed - true表示审核通过，false表示未通过
   * @returns {Promise}
   */
  updateVideoReviewStatus: (id: number | string, isReviewed: boolean) => {
    return api.patch(`/videos/${id}/review`, { reviewed: isReviewed })
  },

  /**
   * 删除视频数据
   * @param {string|number} id - 视频ID
   * @returns {Promise}
   */
  deleteVideo: (id: number) => {
    return api.delete(`/videos/${id}`)
  },

  getBiliBiliVideoInfo: (url: string) => {
    return api.get(`/video/parse?url=${url}`)
  },
}

// Author API 接口
export const authorApi = {
  /**
   * 创建作者
   * @param {Object} authorData - 作者数据对象
   * @returns {Promise}
   */
  createAuthor: (authorData: { uid: number; name: string; avatar?: string }) => {
    return api.post('/authors', authorData)
  },

  /**
   * 获取所有作者
   * @param {Object} params - 查询参数（可选）
   * @returns {Promise}
   */
  getAllAuthors: (params = {}) => {
    return api.get('/authors', { params })
  },

  /**
   * 根据ID获取作者
   * @param {string|number} id - 作者ID
   * @returns {Promise}
   */
  getAuthorById: (id: string | number) => {
    return api.get(`/authors/${id}`)
  },

  /**
   * 更新作者
   * @param {string|number} id - 作者ID
   * @param {Object} authorData - 更新的作者数据
   * @returns {Promise}
   */
  updateAuthor: (
    id: string | number,
    authorData: { uid: number; name: string; avatar?: string },
  ) => {
    return api.put(`/authors/${id}`, authorData)
  },

  /**
   * 删除作者
   * @param {string|number} id - 作者ID
   * @returns {Promise}
   */
  deleteAuthor: (id: string | number) => {
    return api.delete(`/authors/${id}`)
  },
}
