export interface AuthorData {
  id?: number
  uid: number
  name: string
  avatar: string
  master?: VideoData // 代表作
  videos?: VideoData[]
}

export interface VideoData {
  id?: number
  title: string
  aid: number
  author: AuthorData
  url: string
  cover: string
  duration: number
  description: string
  views: number
  isOriginal: boolean
  isCompleted: boolean
  world: string
  background: string
  style: string[]
  hasSystem: boolean
  ctime?: string
}
export interface VideoNewData {
  video: VideoData
  author: AuthorData
}

export interface Filters {
  search: string
  background: string
  style: string
  world: string
  isOriginal: boolean | null
  isCompleted: boolean | null
  hasSystem: boolean | null
}
