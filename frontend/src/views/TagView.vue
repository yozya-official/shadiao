<template>
  <!-- 页面标题 -->
  <PageHeader title="标签" subtitle="浏览标签">
    <template #icon>
      <svg class="w-8 h-8" fill="white" viewBox="0 0 20 20">
        <path
          d="M5 3a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2V5a2 2 0 00-2-2H5zm2 3h6v2H7V6zm0 4h6v2H7v-2zm0 4h6v2H7v-2z"
        ></path>
      </svg>
    </template>
  </PageHeader>

  <div class="container px-6 py-8 relative">
    <!-- 标签信息卡片 -->
    <div
      class="bg-card border border-border rounded-3xl shadow-lg p-8 mb-8 hover:shadow-xl transition-all duration-300"
    >
      <div
        class="flex flex-col md:flex-row items-start md:items-center space-y-6 md:space-y-0 md:space-x-8"
      >
        <div class="flex-1">
          <h1 class="text-3xl font-bold text-foreground mb-2">
            {{ tag?.displayName || tag?.name }}
          </h1>
          <div class="flex items-center space-x-4 text-muted-foreground text-sm mb-4">
            <span>ID: {{ tag?.id }}</span>
            <span v-if="tag?.typeDisplayName">类型: {{ tag.typeDisplayName }}</span>
          </div>
          <p class="text-muted-foreground">{{ tag?.description }}</p>
        </div>

        <!-- 编辑按钮 -->
        <button @click="openEditModal" class="btn btn-primary gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
            ></path>
          </svg>
          编辑标签
        </button>
      </div>
    </div>

    <!-- 作品列表 -->
    <div class="bg-card border border-border rounded-2xl shadow-lg p-6">
      <div class="flex items-center space-x-3 mb-6">
        <div class="w-1 h-8 bg-gradient-to-b from-primary to-chart-2 rounded-full"></div>
        <h2 class="text-2xl font-bold text-foreground">全部作品</h2>
        <div class="flex-1 h-px bg-gradient-to-r from-border to-transparent ml-4"></div>
      </div>

      <div class="relative">
        <VideoContainer :videos="pagedVideos" />

        <!-- 装饰性渐变叠加 -->
        <div class="absolute inset-0 pointer-events-none">
          <div class="absolute top-0 left-0 w-20 h-20 bg-primary/5 rounded-full blur-xl"></div>
          <div
            class="absolute bottom-10 right-10 w-32 h-32 bg-chart-2/5 rounded-full blur-2xl"
          ></div>
        </div>

        <DataPagination
          :total="videos.length"
          :pageSize="pageSize"
          v-model:currentPage="currentPage"
        />
      </div>
    </div>
  </div>

  <!-- 编辑标签对话框 -->
  <dialog id="edit_tag_modal" ref="editModal" class="modal">
    <div class="modal-box bg-card max-w-2xl">
      <h3 class="text-2xl font-bold pb-6 flex items-center space-x-2">
        <svg class="w-6 h-6 text-primary" fill="currentColor" viewBox="0 0 20 20">
          <path
            d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"
          ></path>
          <path
            fill-rule="evenodd"
            d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
            clip-rule="evenodd"
          ></path>
        </svg>
        <span>编辑标签</span>
      </h3>

      <div class="space-y-6">
        <!-- 标签名称 -->
        <div class="space-y-3">
          <label class="flex items-center space-x-2 font-semibold">
            <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M17.707 9.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-7-7A.997.997 0 012 10V5a3 3 0 013-3h5c.256 0 .512.098.707.293l7 7zM5 6a1 1 0 100-2 1 1 0 000 2z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>标签名称</span>
            <span class="text-destructive text-sm">*</span>
          </label>
          <input
            type="text"
            v-model="editForm.name"
            class="w-full input input-primary"
            placeholder="输入标签名称..."
          />
        </div>

        <!-- 显示名称 -->
        <div class="space-y-3">
          <label class="flex items-center space-x-2 font-semibold">
            <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
              <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"></path>
              <path
                fill-rule="evenodd"
                d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>显示名称</span>
            <span class="text-destructive text-sm">*</span>
          </label>
          <input
            type="text"
            v-model="editForm.displayName"
            class="w-full input input-primary"
            placeholder="输入显示名称..."
          />
        </div>

        <!-- 标签类型 -->
        <div class="space-y-3">
          <label class="flex items-center space-x-2 font-semibold">
            <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
              <path
                d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z"
              ></path>
            </svg>
            <span>标签类型</span>
            <span class="text-destructive text-sm">*</span>
          </label>
          <input
            type="text"
            v-model="editForm.type"
            class="w-full input input-primary"
            placeholder="输入标签类型..."
          />
        </div>

        <!-- 类型显示名称 -->
        <div class="space-y-3">
          <label class="flex items-center space-x-2 font-semibold">
            <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>类型显示名称</span>
          </label>
          <input
            type="text"
            v-model="editForm.typeDisplayName"
            class="w-full input input-primary"
            placeholder="输入类型显示名称..."
          />
        </div>

        <!-- 图标 -->
        <div class="space-y-3">
          <label class="flex items-center space-x-2 font-semibold">
            <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-8 3 6 2-4 3 6z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>图标</span>
          </label>
          <input
            type="text"
            v-model="editForm.icon"
            class="w-full input input-primary"
            placeholder="输入图标..."
          />
        </div>

        <!-- 描述 -->
        <div class="space-y-3">
          <label class="flex items-center space-x-2 font-semibold">
            <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>描述</span>
          </label>
          <textarea
            v-model="editForm.description"
            class="w-full textarea textarea-primary h-28 resize-none"
            placeholder="输入标签描述..."
          ></textarea>
        </div>
      </div>

      <div class="flex justify-end gap-3 mt-8 pt-6 border-t border-border">
        <button class="btn btn-ghost" @click="closeEditModal">取消</button>
        <button class="btn btn-primary gap-2" @click="handleSave" :disabled="isSaving">
          <svg v-if="isSaving" class="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
            <circle
              class="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              stroke-width="4"
            ></circle>
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            ></path>
          </svg>
          <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 13l4 4L19 7"
            ></path>
          </svg>
          {{ isSaving ? '保存中...' : '保存' }}
        </button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<script setup lang="ts">
const route = useRoute()

const tag = ref<Tag | null>(null)
const videos = ref<Video[]>([])

const currentPage = ref(1)
const pageSize = ref(12)

const editModal = ref<HTMLDialogElement | null>(null)
const isSaving = ref(false)

const editForm = ref({
  name: '',
  displayName: '',
  type: '',
  typeDisplayName: '',
  icon: '',
  description: '',
})

const pagedVideos = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  return videos.value.slice(start, start + pageSize.value)
})

const loadTag = async (id: number | string) => {
  try {
    const res = await tagApi.getTagById(Number(id))
    tag.value = res.data.tag
    videos.value = res.data.tag.videos.map((v: Partial<Video>) => new Video(v))
  } catch (err) {
    console.error('加载标签信息失败', err)
  }
}

const openEditModal = () => {
  if (tag.value) {
    // 填充表单数据
    editForm.value = {
      name: tag.value.name,
      displayName: tag.value.displayName,
      type: tag.value.type,
      typeDisplayName: tag.value.typeDisplayName,
      icon: tag.value.icon,
      description: tag.value.description,
    }
  }
  editModal.value?.showModal()
}

const closeEditModal = () => {
  editModal.value?.close()
}

const handleSave = async () => {
  if (!tag.value) return

  try {
    isSaving.value = true
    // 调用更新API
    await tagApi.updateTag(tag.value.id, editForm.value)

    // 更新本地数据
    tag.value = {
      ...tag.value,
      ...editForm.value,
    }

    // 关闭对话框
    closeEditModal()

    // 可以添加成功提示
    console.log('标签更新成功')
  } catch (err) {
    console.error('更新标签失败', err)
    // 可以添加错误提示
  } finally {
    isSaving.value = false
  }
}

watch(
  () => route.params.tagId,
  (newId) => {
    if (!newId) return
    const id = Array.isArray(newId) ? newId[0] : newId
    loadTag(id)
  },
  { immediate: true },
)
</script>
