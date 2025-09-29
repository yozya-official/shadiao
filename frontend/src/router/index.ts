import { createRouter, createWebHistory } from 'vue-router'

import VideoNew from '@/views/VideoNew.vue'
import VideoView from '@/views/VideoView.vue'
import ReviewedView from '@/views/ReviewedView.vue'

import AuthorsView from '@/views/AuthorsView.vue'
import AuthorView from '@/views/AuthorView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/videos',
    },
    {
      path: '/video/new',
      name: 'video-new',
      component: VideoNew,
    },
    {
      path: '/videos',
      name: 'videos',
      component: VideoView,
    },
    {
      path: '/videos/review',
      name: 'videos-review',
      component: ReviewedView,
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
  ],
})

export default router
