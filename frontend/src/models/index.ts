// 标签表
export interface Tag {
  id: number
  name: string
  type: string
  displayName: string
  typeDisplayName: string
  icon: string
  description: string
  videos?: Video[] // 可选，多对多关联视频
  createdAt: string
  updatedAt: string
}

// 分类表
export interface Category {
  id: number
  name: string
  type: string
  description?: string
  videos?: Video[]
  createdAt: string
  updatedAt: string
}

// 作者表
export interface Author {
  id: number
  uid: number
  name: string
  avatar?: string
  videos?: Video[]
  createdAt?: string
  updatedAt?: string
}

// 视频表
export class Video {
  id = 0
  aid = 0
  title = ''
  authorId = 0
  author?: Author
  url = ''
  cover = ''
  description = ''
  duration = 0
  views = 0
  tags: Tag[] = []
  categories: Category[] = []
  reviewed = false
  ctime = new Date().toISOString()
  createdAt = new Date().toISOString()
  updatedAt = new Date().toISOString()

  constructor(init?: Partial<Video>) {
    Object.assign(this, init)
  }

  getBooleanTag(type: string): boolean | null {
    if (!this.tags) return null
    const tag = this.tags.find((t) => t.type === type)
    if (!tag) return null
    if (tag.name === 'true') return true
    if (tag.name === 'false') return false
    return null
  }

  isOriginal(): boolean | null {
    return this.getBooleanTag('isOriginal')
  }

  isCompleted(): boolean | null {
    return this.getBooleanTag('isCompleted')
  }

  hasSystem(): boolean | null {
    return this.getBooleanTag('hasSystem')
  }

  background(): Tag | undefined {
    return this.tags.find((tag) => {
      return tag.type == 'background'
    })
  }

  world(): Tag | undefined {
    return this.tags.find((tag) => {
      return tag.type == 'world'
    })
  }

  style(): Tag[] {
    return this.tags.filter((tag) => {
      return tag.type == 'style'
    })
  }

  static create(init?: Partial<Video>) {
    return new Video(init)
  }

  clone(): Video {
    return new Video(JSON.parse(JSON.stringify(this)))
  }
}

export interface VideoData {
  id?: number
  title: string
  aid: number
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
  ctime: string
}

export interface VideoNewData {
  video: VideoData
  author: Author
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
