<template>
  <div class="min-h-screen bg-background">
    <!-- 头部装饰背景 -->
    <div
      class="absolute top-0 left-0 w-full h-48 bg-gradient-to-r from-primary/10 via-chart-2/10 to-chart-3/10"
    ></div>

    <div class="container mx-auto px-6 py-8 max-w-6xl relative">
      <!-- 页面标题 -->
      <div class="text-center mb-8">
        <div
          class="inline-flex items-center justify-center size-16 rounded-2xl bg-gradient-to-br from-primary to-chart-2 text-primary-foreground mb-6 shadow-xl hover:scale-105 transition-transform duration-300"
        >
          <svg class="w-8 h-8" fill="white" viewBox="0 0 20 20">
            <path
              d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z"
            ></path>
          </svg>
        </div>
        <h2 class="text-foreground mt-4 mb-2">作者列表</h2>
        <p class="text-muted-foreground text-sm">管理和浏览所有创作者</p>
        <div
          class="w-24 h-1 bg-gradient-to-r from-primary to-chart-2 mx-auto mt-4 rounded-full"
        ></div>
      </div>

      <!-- 作者卡片网格 -->
      <div
        v-if="authors.length > 0"
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 mb-8"
      >
        <div
          v-for="author in pagedAuthors"
          :key="author.id"
          class="group bg-card border border-border rounded-2xl shadow-md hover:shadow-xl transition-all duration-300 overflow-hidden hover:-translate-y-2"
        >
          <div class="p-6 text-center">
            <!-- 头像 -->
            <div class="mb-4 cursor-pointer" @click="goToBiliSpace(author.uid)">
              <div
                class="w-20 h-20 relative mx-auto rounded-2xl bg-gradient-to-br from-primary/20 to-chart-2/20 p-1 group-hover:from-primary/40 group-hover:to-chart-2/40 transition-all duration-300"
              >
                <img
                  :src="author.avatar"
                  :alt="author.name + '的头像'"
                  referrerpolicy="no-referrer"
                  class="w-full h-full rounded-xl object-cover group-hover:scale-105 transition-transform duration-300"
                />
                <!-- 悬停效果指示器 -->
                <div
                  class="absolute inset-0 rounded-2xl bg-primary/0 group-hover:bg-primary/10 transition-colors duration-300 flex items-center justify-center opacity-0 group-hover:opacity-100"
                >
                  <svg
                    class="w-6 h-6 text-primary"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                    ></path>
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                    ></path>
                  </svg>
                </div>
              </div>
            </div>

            <!-- 作者信息 -->
            <h2
              class="text-lg font-bold text-foreground mb-1 group-hover:text-primary transition-colors duration-200"
            >
              {{ author.name }}
            </h2>
            <p class="text-sm text-muted-foreground mb-4">UID: {{ author.uid }}</p>

            <!-- 操作按钮 -->
            <div class="flex justify-center gap-2">
              <button
                @click="delAuthorConfirm(author)"
                class="btn btn-error hover:bg-destructive rounded-lg text-sm font-medium transition-all duration-200"
              >
                删除作者
              </button>
              <button
                @click="goToAuthor(author.id)"
                class="btn btn-primary rounded-lg text-sm font-medium transition-all duration-200 hover:shadow-md"
              >
                <a
                  :href="`https://space.bilibili.com/${author.uid}`"
                  target="_blank"
                  class="flex items-center space-x-1"
                  >作者主页</a
                >
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 分页控制 -->
      <div v-if="totalPages > 1" class="flex justify-center mt-8">
        <div
          class="flex items-center space-x-2 bg-card border border-border rounded-xl p-2 shadow-sm"
        >
          <button
            v-for="page in totalPages"
            :key="page"
            @click="currentPage = page"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200',
              currentPage === page
                ? 'bg-primary text-primary-foreground shadow-sm'
                : 'text-muted-foreground hover:text-foreground hover:bg-accent',
            ]"
          >
            {{ page }}
          </button>
        </div>
      </div>

      <!-- 空状态 -->
      <div v-else-if="authors.length === 0" class="text-center py-20">
        <div
          class="w-32 h-32 mx-auto mb-6 rounded-full bg-muted/30 flex items-center justify-center"
        >
          <svg
            class="w-16 h-16 text-muted-foreground"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="1.5"
              d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
            ></path>
          </svg>
        </div>
        <h3 class="text-2xl font-bold text-foreground mb-2">暂无作者数据</h3>
        <p class="text-muted-foreground">还没有添加任何创作者信息</p>
      </div>

      <!-- 删除确认模态框 -->
      <div
        v-if="selectedAuthor"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click="closeModal"
      >
        <div class="fixed inset-0 bg-background/80 backdrop-blur-sm"></div>
        <div
          class="relative bg-card border border-border rounded-2xl shadow-2xl p-6 w-full max-w-md"
          @click.stop
        >
          <div class="text-center mb-6">
            <div
              class="w-16 h-16 mx-auto mb-4 rounded-full bg-destructive/10 flex items-center justify-center"
            >
              <svg
                class="w-8 h-8 text-destructive"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-bold text-foreground mb-2">确认删除作者</h3>
            <p class="text-muted-foreground mb-1">
              将删除作者 <span class="font-medium text-foreground">{{ selectedAuthor.name }}</span>
            </p>
            <p class="text-sm text-destructive">此操作不可撤销</p>
          </div>

          <div class="flex space-x-3">
            <button
              @click="closeModal"
              class="flex-1 px-4 py-2 border border-border text-foreground rounded-lg hover:bg-accent transition-colors duration-200"
            >
              取消
            </button>
            <button
              @click="delAuthor"
              class="flex-1 px-4 py-2 bg-destructive text-destructive-foreground rounded-lg hover:bg-destructive/90 transition-colors duration-200"
            >
              确认删除
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { authorApi } from '@/api'
import type { AuthorData } from '@/models/video'
import { handleApiError } from '@/utils'
import { toast } from '@yuelioi/toast'

const authors = ref<AuthorData[]>([])
const selectedAuthor = ref<AuthorData | null>(null)
const currentPage = ref(1)
const pageSize = 9

const router = useRouter()

const totalPages = computed(() => Math.ceil(authors.value.length / pageSize))

const pagedAuthors = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return authors.value.slice(start, start + pageSize)
})

const loadAuthors = async () => {
  try {
    const resp = await authorApi.getAllAuthors()
    authors.value = resp.data
  } catch (err) {
    handleApiError(err, '获取作者')
  }
}

const delAuthorConfirm = (author: AuthorData) => {
  selectedAuthor.value = author
}

const delAuthor = async () => {
  try {
    const id = selectedAuthor.value?.id
    if (!id) throw 'id 不能为空'
    await authorApi.deleteAuthor(id)
    toast.success('删除成功')
    closeModal()
    await loadAuthors()
  } catch (err) {
    handleApiError(err, '删除')
  }
}

const closeModal = () => {
  selectedAuthor.value = null
}

const goToAuthor = (id?: number) => {
  if (!id) return
  router.push({ name: 'author', params: { authorId: id } })
}

const goToBiliSpace = (uid?: number) => {
  window.location.href = `https://space.bilibili.com/${uid}`
}

onMounted(() => {
  loadAuthors()
})
</script>

<style scoped>
/* 卡片悬停效果 */
.group:hover {
  transform: translateY(-8px);
}

/* 模态框动画 */
.fixed.inset-0 {
  animation: fadeIn 0.2s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.relative.bg-card {
  animation: slideUp 0.3s ease-out;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}
</style>
