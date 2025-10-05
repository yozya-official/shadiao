import { createRouter, createWebHistory } from 'vue-router'

import VideoNew from '@/views/VideoNew.vue'
import VideoView from '@/views/VideoView.vue'
import ReviewedView from '@/views/ReviewedView.vue'
import VideoUpdate from '@/views/VideoUpdate.vue'

import AuthorsView from '@/views/AuthorsView.vue'
import AuthorView from '@/views/AuthorView.vue'

import TagsView from '@/views/TagsView.vue'
import TagView from '@/views/TagView.vue'

import CommunicateView from '@/views/CommunicateView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'videos',
      component: VideoView,
    },
    {
      path: '/videos/:id',
      name: 'videos-update',
      component: VideoUpdate,
    },
    {
      path: '/video/new',
      name: 'video-new',
      component: VideoNew,
    },

    {
      path: '/videos/review',
      name: 'videos-review',
      component: ReviewedView,
    },
    {
      path: '/tags',
      name: 'tags',
      component: TagsView,
    },
    {
      path: '/tags/:tagId',
      name: 'tag',
      component: TagView,
    },
    {
      path: '/authors/:authorId',
      name: 'author',
      component: AuthorView,
    },
    {
      path: '/authors',
      name: 'authors',
      component: AuthorsView,
    },
    {
      path: '/communicate',
      name: 'communicate',
      component: CommunicateView,
    },
  ],
})

export default router
