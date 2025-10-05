<template>
  <!-- 头部装饰背景 -->
  <div class="w-full bg-gradient-to-r from-primary/10 via-chart-2/10 to-chart-3/10 relative">
    <div class="container px-6 max-w-6xl">
      <div class="flex items-center justify-between py-6">
        <!-- 左侧：图标和文字 -->
        <div class="flex items-center space-x-4">
          <!-- 图标/徽章 -->
          <div
            class="inline-flex items-center justify-center size-16 rounded-2xl bg-gradient-to-br from-primary to-chart-2 text-primary-foreground shadow-xl hover:scale-105 transition-transform duration-300"
          >
            <slot name="icon">
              <!-- 默认图标 -->
              <svg class="w-8 h-8" fill="white" viewBox="0 0 20 20">
                <path
                  d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z"
                />
              </svg>
            </slot>
          </div>

          <!-- 标题和副标题 -->
          <div class="text-left">
            <h2 class="text-2xl font-bold text-foreground mb-1">{{ title }}</h2>
            <p v-if="subtitle" class="text-muted-foreground text-sm">{{ subtitle }}</p>
          </div>
        </div>

        <!-- 右侧：操作按钮 -->
        <div class="flex items-center space-x-3">
          <slot name="actions"></slot>
        </div>
      </div>

      <!-- 装饰线 -->
      <div
        class="h-1 bg-gradient-to-r w-full from-primary via-chart-2 to-transparent rounded-full"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string
  subtitle?: string
}>()
</script>
