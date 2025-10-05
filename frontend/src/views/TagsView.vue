<template>
  <PageHeader title="标签" subtitle="浏览和管理标签">
    <template #icon>
      <svg class="w-8 h-8" fill="white" viewBox="0 0 20 20">
        <path
          fill-rule="evenodd"
          d="M17.707 9.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-7-7A.997.997 0 012 10V5a3 3 0 013-3h5c.256 0 .512.098.707.293l7 7zM5 6a1 1 0 100-2 1 1 0 000 2z"
          clip-rule="evenodd"
        ></path>
      </svg>
    </template>
    <template #actions>
      <button
        v-if="isLogin"
        @click="openAddModal"
        class="btn btn-primary gap-2 shadow-lg hover:shadow-xl transition-all duration-300"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 6v6m0 0v6m0-6h6m-6 0H6"
          ></path>
        </svg>
        添加标签
      </button>
    </template>
  </PageHeader>

  <div class="container px-6 py-8 relative">
    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="card-primary p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-muted-foreground mb-1">总标签数</p>
            <p class="text-3xl font-bold text-primary">{{ tags.length }}</p>
          </div>
          <div class="w-14 h-14 bg-primary/10 rounded-full flex items-center justify-center">
            <svg class="w-7 h-7 text-primary" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M17.707 9.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-7-7A.997.997 0 012 10V5a3 3 0 013-3h5c.256 0 .512.098.707.293l7 7zM5 6a1 1 0 100-2 1 1 0 000 2z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- 标签卡片网格 -->
    <!-- 标签卡片网格 -->
    <div
      v-if="tags.length > 0"
      class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-3 mb-8"
    >
      <div v-for="tag in pagedTags" :key="tag.id" class="group card card-hover overflow-hidden">
        <div class="p-3 relative">
          <!-- 头部：图标、标题、类型 -->
          <div class="flex items-start gap-2 mb-2">
            <div
              class="size-10 rounded-lg bg-gradient-to-br from-primary/20 to-chart-2/20 flex items-center justify-center flex-shrink-0 group-hover:scale-110 transition-transform duration-300"
            >
              <span v-if="tag.icon" class="text-base">{{ tag.icon }}</span>
              <svg v-else class="w-4 h-4 text-primary" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M17.707 9.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-7-7A.997.997 0 012 10V5a3 3 0 013-3h5c.256 0 .512.098.707.293l7 7zM5 6a1 1 0 100-2 1 1 0 000 2z"
                  clip-rule="evenodd"
                ></path>
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <h2
                class="text-sm font-semibold text-foreground group-hover:text-primary transition-colors duration-200 line-clamp-1"
              >
                {{ tag.displayName || tag.name }}
              </h2>
              <span class="text-xs text-muted-foreground">
                {{ tag.videos?.length || 0 }} 个视频
              </span>
            </div>
            <div
              class="absolute right-2 top-2 px-1.5 py-0.5 rounded badge-secondary text-xs font-medium text-muted-foreground flex-shrink-0"
            >
              {{ tag.typeDisplayName || tag.type }}
            </div>
          </div>

          <!-- 描述 -->
          <p class="text-xs text-muted-foreground mb-2 line-clamp-2 leading-relaxed">
            {{ tag.description || '暂无描述' }}
          </p>

          <!-- 操作按钮 -->
          <div class="flex gap-1">
            <button
              @click="goToTag(tag.id)"
              class="flex-1 btn btn-sm btn-primary rounded text-xs font-medium transition-all duration-200 hover:shadow-md py-1"
            >
              详情
            </button>
            <button
              v-if="isLogin"
              @click="openEditModal(tag)"
              class="btn btn-sm btn-ghost rounded transition-all duration-200 hover:bg-accent px-2 py-1"
              title="编辑"
            >
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                ></path>
              </svg>
            </button>
            <button
              v-if="isLogin"
              @click="delTagConfirm(tag)"
              class="btn btn-sm btn-ghost text-destructive rounded transition-all duration-200 hover:bg-destructive/10 px-2 py-1"
              title="删除"
            >
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                ></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 分页控制 -->
    <div v-if="totalPages > 1" class="flex justify-center mt-8">
      <div
        class="flex items-center space-x-2 bg-card border border-border rounded-xl p-2 shadow-lg"
      >
        <button
          @click="currentPage = Math.max(1, currentPage - 1)"
          :disabled="currentPage === 1"
          :class="[
            'p-2 rounded-lg transition-all duration-200',
            currentPage === 1
              ? 'text-muted-foreground/50 cursor-not-allowed'
              : 'text-muted-foreground hover:text-foreground hover:bg-accent',
          ]"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 19l-7-7 7-7"
            ></path>
          </svg>
        </button>

        <button
          v-for="page in totalPages"
          :key="page"
          @click="currentPage = page"
          :class="[
            'px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200',
            currentPage === page
              ? 'bg-primary text-primary-foreground shadow-sm scale-110'
              : 'text-muted-foreground hover:text-foreground hover:bg-accent',
          ]"
        >
          {{ page }}
        </button>

        <button
          @click="currentPage = Math.min(totalPages, currentPage + 1)"
          :disabled="currentPage === totalPages"
          :class="[
            'p-2 rounded-lg transition-all duration-200',
            currentPage === totalPages
              ? 'text-muted-foreground/50 cursor-not-allowed'
              : 'text-muted-foreground hover:text-foreground hover:bg-accent',
          ]"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </button>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-else-if="tags.length === 0" class="text-center py-20">
      <div
        class="w-32 h-32 mx-auto mb-6 rounded-full bg-gradient-to-br from-primary/10 to-chart-2/10 flex items-center justify-center"
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
            d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"
          ></path>
        </svg>
      </div>
      <h3 class="text-2xl font-bold text-foreground mb-2">暂无标签数据</h3>
      <p class="text-muted-foreground mb-6">还没有添加任何标签信息</p>
      <button
        v-if="isLogin"
        @click="openAddModal"
        class="btn btn-primary gap-2 shadow-lg hover:shadow-xl"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 6v6m0 0v6m0-6h6m-6 0H6"
          ></path>
        </svg>
        创建第一个标签
      </button>
    </div>

    <!-- 删除确认模态框 -->
    <dialog
      id="delete_tag_modal"
      ref="deleteModal"
      class="dialog dialog-slide-down dialog-bounce"
      @click.self="deleteModal?.close()"
    >
      <div class="dialog-body dialog-body-md">
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
          <h3 class="text-xl font-bold text-foreground mb-2">确认删除标签</h3>
          <p class="text-muted-foreground mb-1">
            将删除标签
            <span class="font-medium text-foreground">{{ selectedTag?.name }}</span>
          </p>
          <p class="text-sm text-destructive">此操作不可撤销</p>
        </div>

        <div class="flex gap-3">
          <button
            @click="closeDeleteModal"
            class="flex-1 btn btn-ghost rounded-lg transition-colors duration-200"
          >
            取消
          </button>
          <button
            @click="delTag"
            class="flex-1 btn btn-destructive rounded-lg transition-colors duration-200"
          >
            确认删除
          </button>
        </div>
      </div>
    </dialog>

    <!-- 添加/编辑标签模态框 -->
    <dialog id="edit_tag_modal" ref="editModal" class="dialog" @click.self="editModal?.close()">
      <div class="dialog-body-lg dialog-body m-2">
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
          <span>{{ isEditMode ? '编辑标签' : '添加标签' }}</span>
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
    </dialog>
  </div>
</template>

<script setup lang="ts">
const { isLogin } = useAuthStore()

const tags = ref<Tag[]>([])
const selectedTag = ref<Tag | null>(null)
const currentPage = ref(1)
const pageSize = 12

const router = useRouter()

const deleteModal = ref<HTMLDialogElement | null>(null)
const editModal = ref<HTMLDialogElement | null>(null)
const isSaving = ref(false)
const isEditMode = ref(false)

const editForm = ref({
  name: '',
  displayName: '',
  type: '',
  typeDisplayName: '',
  icon: '',
  description: '',
})

const totalPages = computed(() => Math.ceil(tags.value.length / pageSize))

const pagedTags = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return tags.value.slice(start, start + pageSize)
})

const loadTags = async () => {
  try {
    const resp = await tagApi.getAllTags()
    tags.value = resp.data.tags
  } catch (err) {
    handleApiError(err, '获取标签')
  }
}

const delTagConfirm = (tag: Tag) => {
  selectedTag.value = tag
  deleteModal.value?.showModal()
}

const delTag = async () => {
  try {
    const id = selectedTag.value?.id
    if (!id) throw 'id 不能为空'
    await tagApi.deleteTag(id)
    toast.success('删除成功')
    closeDeleteModal()
    await loadTags()
  } catch (err) {
    handleApiError(err, '删除')
  }
}

const closeDeleteModal = () => {
  deleteModal.value?.close()
  selectedTag.value = null
}

const openAddModal = () => {
  isEditMode.value = false
  editForm.value = {
    name: '',
    displayName: '',
    type: '',
    typeDisplayName: '',
    icon: '',
    description: '',
  }
  editModal.value?.showModal()
}

const openEditModal = (tag: Tag) => {
  isEditMode.value = true
  selectedTag.value = tag
  editForm.value = {
    name: tag.name,
    displayName: tag.displayName,
    type: tag.type,
    typeDisplayName: tag.typeDisplayName,
    icon: tag.icon,
    description: tag.description,
  }
  editModal.value?.showModal()
}

const closeEditModal = () => {
  editModal.value?.close()
  selectedTag.value = null
}

const handleSave = async () => {
  try {
    // 验证必填字段
    if (!editForm.value.name || !editForm.value.displayName || !editForm.value.type) {
      toast.error('请填写必填字段')
      return
    }

    isSaving.value = true

    if (isEditMode.value && selectedTag.value) {
      // 编辑模式
      await tagApi.updateTag(selectedTag.value.id, editForm.value)
      toast.success('标签更新成功')
    } else {
      // 添加模式
      await tagApi.createTag(editForm.value)
      toast.success('标签创建成功')
    }

    closeEditModal()
    await loadTags()
  } catch (err) {
    handleApiError(err, isEditMode.value ? '更新标签' : '创建标签')
  } finally {
    isSaving.value = false
  }
}

const goToTag = (id?: number) => {
  if (!id) return
  router.push({ name: 'tag', params: { tagId: id } })
}

onMounted(() => {
  loadTags()
})
</script>
