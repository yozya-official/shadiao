<template>
  <!-- 页面标题 -->
  <PageActionHeader title="作者" subtitle="管理和浏览创作者">
    <template #icon>
      <svg class="w-8 h-8" fill="white" viewBox="0 0 20 20">
        <path
          d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z"></path>
      </svg>
    </template>
  </PageActionHeader>
  <div class="container px-6 py-8 relative">
    <!-- 作者信息卡片 -->
    <div
      class="bg-card border border-border rounded-3xl shadow-lg p-8 mb-8 hover:shadow-xl transition-all duration-300">
      <div
        class="flex flex-col md:flex-row items-start md:items-center space-y-6 md:space-y-0 md:space-x-8">
        <!-- 头像区域 -->
        <div class="relative">
          <div
            class="w-28 h-28 rounded-2xl bg-gradient-to-br from-primary/80 to-chart-2/80 p-1 shadow-lg">
            <img
              :src="author.avatar"
              referrerpolicy="no-referrer"
              alt="头像"
              class="w-full h-full rounded-xl object-cover hover:scale-105 transition-transform duration-300" />
          </div>
          <!-- 装饰元素 -->
          <div
            class="absolute -top-2 -right-2 w-6 h-6 bg-chart-4 rounded-full border-3 border-card shadow-md"></div>
        </div>

        <!-- 作者信息 -->
        <div class="flex-1">
          <h1 class="text-3xl font-bold text-foreground mb-2">
            {{ author.name }}
          </h1>
          <div class="flex items-center space-x-4 text-muted-foreground text-sm mb-4">
            <span>UID: {{ author.uid }}</span>
            <span class="w-1 h-1 bg-muted-foreground rounded-full"></span>
            <a
              :href="`https://space.bilibili.com/${author.uid}`"
              target="_blank"
              class="flex items-center space-x-1 text-primary hover:text-primary/80 transition-colors duration-200">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path>
              </svg>
              <span>访问B站主页</span>
            </a>
          </div>
        </div>
      </div>
    </div>

    <!-- 作品区域 -->
    <div class="bg-card border border-border rounded-2xl shadow-lg p-6">
      <div class="flex items-center space-x-3 mb-6">
        <div class="w-1 h-8 bg-gradient-to-b from-primary to-chart-2 rounded-full"></div>
        <h2 class="text-2xl font-bold text-foreground">全部作品</h2>
        <div class="flex-1 h-px bg-gradient-to-r from-border to-transparent ml-4"></div>
      </div>

      <div class="relative">
        <VideoContainer :videos="author.videos || []"></VideoContainer>

        <!-- 装饰性渐变叠加 -->
        <div class="absolute inset-0 pointer-events-none">
          <div class="absolute top-0 left-0 w-20 h-20 bg-primary/5 rounded-full blur-xl"></div>
          <div
            class="absolute bottom-10 right-10 w-32 h-32 bg-chart-2/5 rounded-full blur-2xl"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { PageActionHeader } from '@yuelioi/ui'

const route = useRoute()
const authorId = ref('')

const author = ref<Author>({
  name: '',
  avatar: '',
  uid: 0,
  id: 0,
})

watch(
  () => route.params.authorId,
  async (newId) => {
    if (!newId) return

    const id = Array.isArray(newId) ? newId[0] : newId

    try {
      const res = await authorApi.getAuthorById(id)
      author.value = res.data.author
      authorId.value = id
      author.value.videos = author.value.videos?.map((v: Partial<Video>) => new Video(v))
      console.log(author.value)
    } catch (err) {
      console.error('加载作者信息失败', err)
    }
  },
  { immediate: true },
)
</script>
