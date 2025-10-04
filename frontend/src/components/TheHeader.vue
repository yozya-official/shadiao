<template>
  <header
    class="sticky top-0 z-50 w-full border-b border-border/40 bg-background/80 backdrop-blur-xl supports-[backdrop-filter]:bg-background/60 transition-all duration-300"
  >
    <div class="container mx-auto flex h-16 max-w-screen-2xl items-center justify-between px-6">
      <!-- Logo 区域 -->
      <div class="flex items-center space-x-3">
        <a
          href="/"
          aria-label="首页"
          class="group flex items-center space-x-3 transition-all duration-200 hover:opacity-90"
        >
          <div class="relative">
            <!-- Logo 容器 -->
            <div
              class="w-10 h-10 rounded-xl bg-gradient-to-br from-primary to-chart-2 p-2 shadow-md group-hover:shadow-lg transition-all duration-300 group-hover:scale-105"
            >
              <img
                :src="logo"
                :alt="siteTitle"
                class="w-full h-full object-contain filter brightness-0 invert"
              />
            </div>
            <!-- 装饰光晕 -->
            <div
              class="absolute inset-0 rounded-xl bg-gradient-to-br from-primary to-chart-2 opacity-0 group-hover:opacity-20 transition-opacity duration-300 blur-md -z-10"
            ></div>
          </div>
          <!-- 站点标题 -->
          <div class="hidden sm:block">
            <h1
              class="text-lg font-bold text-foreground group-hover:text-primary transition-colors duration-200"
            >
              {{ siteTitle }}
            </h1>
          </div>
        </a>
      </div>

      <!-- 桌面端导航 -->
      <!-- 桌面端导航 -->
      <nav class="hidden items-center md:flex">
        <div class="flex items-center space-x-1 bg-muted rounded-full p-1">
          <button
            v-for="item in navigationItems"
            :key="item.name"
            @click="navigateTo(item.path)"
            class="relative px-4 py-2 cursor-pointer text-sm font-medium transition-all duration-200 rounded-full"
            :class="[
              isActive(item.path)
                ? 'text-primary-foreground bg-primary shadow-sm'
                : 'text-muted-foreground hover:text-foreground hover:bg-background/60',
            ]"
          >
            {{ item.name }}
            <!-- 活跃状态指示器 -->
            <div
              v-if="isActive(item.path)"
              class="absolute bottom-0 left-1/2 transform -translate-x-1/2 w-1 h-1 bg-primary-foreground rounded-full translate-y-1"
            ></div>
          </button>
        </div>
      </nav>

      <!-- 右侧功能区 -->
      <div class="flex items-center">
        <!-- 主题切换器+设置 -->
        <div class="relative flex gap-3 pr-2">
          <ThemeToggle class="hover:scale-105 transition-transform duration-200"></ThemeToggle>
          <button onclick="secret_setting.show()" class="cursor-pointer hover:scale-105">
            <svg
              class="size-5"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
            >
              <g
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
              >
                <path
                  d="M9.671 4.136a2.34 2.34 0 0 1 4.659 0a2.34 2.34 0 0 0 3.319 1.915a2.34 2.34 0 0 1 2.33 4.033a2.34 2.34 0 0 0 0 3.831a2.34 2.34 0 0 1-2.33 4.033a2.34 2.34 0 0 0-3.319 1.915a2.34 2.34 0 0 1-4.659 0a2.34 2.34 0 0 0-3.32-1.915a2.34 2.34 0 0 1-2.33-4.033a2.34 2.34 0 0 0 0-3.831A2.34 2.34 0 0 1 6.35 6.051a2.34 2.34 0 0 0 3.319-1.915"
                />
                <circle cx="12" cy="12" r="3" />
              </g>
            </svg>
          </button>
        </div>

        <!-- 移动端菜单按钮 -->
        <button
          type="button"
          id="mobile-menu-toggle"
          class="relative inline-flex h-10 w-10 items-center justify-center rounded-xl bg-muted/40 hover:bg-muted/60 text-muted-foreground hover:text-foreground transition-all duration-200 hover:scale-105 focus-visible:ring-2 focus-visible:ring-primary/50 focus-visible:outline-none md:hidden"
          aria-label="菜单"
          @click="toggleMobileMenu"
        >
          <!-- 汉堡图标动画 -->
          <div class="relative w-5 h-5">
            <span
              class="absolute left-0 top-1 block h-0.5 w-5 transform bg-current transition-all duration-300 ease-in-out"
              :class="mobileMenuOpen ? 'rotate-45 translate-y-1.5' : ''"
            ></span>
            <span
              class="absolute left-0 top-2.5 block h-0.5 w-5 bg-current transition-all duration-200 ease-in-out"
              :class="mobileMenuOpen ? 'opacity-0' : ''"
            ></span>
            <span
              class="absolute left-0 top-4 block h-0.5 w-5 transform bg-current transition-all duration-300 ease-in-out"
              :class="mobileMenuOpen ? '-rotate-45 -translate-y-1.5' : ''"
            ></span>
          </div>
        </button>
      </div>
    </div>
  </header>

  <!-- 设置密钥 -->
  <dialog id="secret_setting" ref="secretModal" class="modal">
    <div class="modal-box bg-card">
      <h3 class="text-lg font-bold pb-4">请输入密钥</h3>
      <div class="flex justify-between gap-3">
        <input type="text" v-model="apiKey" class="input input-primary" />
        <button class="btn btn-primary" onclick="secret_setting.close()">确认</button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>

  <!-- 移动端菜单 -->
  <Transition
    enter-active-class="transition-all duration-300 ease-out"
    enter-from-class="opacity-0 -translate-y-4 scale-95"
    enter-to-class="opacity-100 translate-y-0 scale-100"
    leave-active-class="transition-all duration-200 ease-in"
    leave-from-class="opacity-100 translate-y-0 scale-100"
    leave-to-class="opacity-0 -translate-y-4 scale-95"
  >
    <div
      v-if="mobileMenuOpen"
      id="mobile-menu"
      class="fixed top-16 right-0 left-0 z-40 border-b border-border/40 bg-card/95 backdrop-blur-xl drop-shadow-2xl md:hidden"
    >
      <div class="container mx-auto max-w-screen-2xl px-6 py-6">
        <!-- 移动端导航 -->
        <nav class="space-y-2">
          <button
            v-for="item in navigationItems"
            :key="item.name"
            @click="navigateTo(item.path)"
            class="group flex items-center justify-between rounded-xl px-4 py-3 text-sm font-medium transition-all duration-200 w-full text-left"
            :class="[
              isActive(item.path)
                ? 'text-primary bg-primary/10 border border-primary/20'
                : 'text-foreground hover:text-primary hover:bg-accent/50 border border-transparent hover:border-border/50',
            ]"
          >
            <span>{{ item.name }}</span>
            <!-- 箭头图标 -->
            <svg
              class="w-4 h-4 opacity-0 group-hover:opacity-100 transform translate-x-0 group-hover:translate-x-1 transition-all duration-200"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5l7 7-7 7"
              ></path>
            </svg>
          </button>
        </nav>

        <!-- 移动端底部装饰 -->
        <div class="mt-6 pt-4 border-t border-border/20">
          <div class="flex items-center justify-center space-x-2 text-xs text-muted-foreground">
            <div class="w-2 h-2 rounded-full bg-primary/30"></div>
            <span>{{ siteTitle }}</span>
            <div class="w-2 h-2 rounded-full bg-chart-2/30"></div>
          </div>
        </div>
      </div>
    </div>
  </Transition>

  <!-- 移动端菜单背景遮罩 -->
  <Transition
    enter-active-class="transition-opacity duration-300 ease-out"
    enter-from-class="opacity-0"
    enter-to-class="opacity-100"
    leave-active-class="transition-opacity duration-200 ease-in"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div
      v-if="mobileMenuOpen"
      class="fixed inset-0 z-30 bg-background/80 backdrop-blur-sm md:hidden"
      @click="closeMobileMenu"
    ></div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import ThemeToggle from '@yuelioi/components/theme-toggle'
import '@yuelioi/components/theme-toggle.css'
import logo from '/logo.png'
import { useRoute } from 'vue-router'

import { useAuthStore } from '@/stores/authStore'
import { storeToRefs } from 'pinia'
import router from '@/router'

const store = useAuthStore()
const { apiKey } = storeToRefs(store)

defineProps(['siteTitle'])

const secretModal = ref<HTMLDialogElement>()

const navigationItems = ref([
  { name: '视频', path: '/', active: false },
  { name: 'UP主', path: '/authors', active: false },
  { name: '审核', path: '/videos/review', active: false },
  { name: '交流', path: '/communicate', active: false },
  { name: '提交视频', path: '/video/new', active: false },
])

const route = useRoute()

const navigateTo = (path: string) => {
  router.push(path)
}

const isActive = (path: string) => {
  if (path === '/') {
    return route.path === '/'
  }
  return route.path.startsWith(path)
}

// 移动端菜单开关
const mobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
  // 防止页面滚动
  if (mobileMenuOpen.value) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
}

const closeMobileMenu = () => {
  mobileMenuOpen.value = false
  document.body.style.overflow = ''
}

// 点击外部关闭菜单
function handleClickOutside(e: MouseEvent) {
  const menu = document.getElementById('mobile-menu')
  const toggle = document.getElementById('mobile-menu-toggle')
  if (!menu?.contains(e.target as Node) && !toggle?.contains(e.target as Node)) {
    closeMobileMenu()
  }
}

// ESC 键关闭菜单
function handleEscapeKey(e: KeyboardEvent) {
  if (e.key === 'Escape') {
    closeMobileMenu()
  }
}

// 滚动效果 - 添加阴影和背景模糊
function initScrollEffect() {
  const header = document.querySelector('header')
  let lastScrollY = window.scrollY

  const onScroll = () => {
    const scrollY = window.scrollY

    // 滚动阴影效果
    if (scrollY > 10) {
      header?.classList.add('shadow-lg', 'bg-background/90')
      header?.classList.remove('bg-background/80')
    } else {
      header?.classList.remove('shadow-lg', 'bg-background/90')
      header?.classList.add('bg-background/80')
    }

    // 滚动方向检测（可用于隐藏/显示导航栏）
    if (scrollY > lastScrollY && scrollY > 100) {
      // 向下滚动 - 可以添加隐藏逻辑
      header?.classList.add('transform', '-translate-y-1')
    } else {
      // 向上滚动
      header?.classList.remove('transform', '-translate-y-1')
    }

    lastScrollY = scrollY
  }

  window.addEventListener('scroll', onScroll, { passive: true })
  return () => window.removeEventListener('scroll', onScroll)
}

let cleanupScroll: (() => void) | undefined

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  document.addEventListener('keydown', handleEscapeKey)
  cleanupScroll = initScrollEffect()
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
  document.removeEventListener('keydown', handleEscapeKey)
  cleanupScroll?.()
  // 确保页面可滚动
  document.body.style.overflow = ''
})
</script>

<style scoped lang="css">
/* Logo 悬停效果 */
.group:hover .w-10 {
  transform: scale(1.05) rotate(2deg);
}

/* 导航链接活跃状态动画 */
nav a {
  position: relative;
  overflow: hidden;
}

nav a::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg,
    transparent,
    var(--primary-rgb, 84 114 183) / 0.1,
    transparent
  );
  transition: left 0.5s;
}

nav a:hover::before {
  left: 100%;
}

/* 汉堡菜单动画优化 */
#mobile-menu-toggle span {
  transform-origin: center;
}

/* 移动端菜单项悬停效果 */
#mobile-menu a {
  backdrop-filter: blur(8px);
}

/* 头部毛玻璃效果增强 */
header {
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
}

/* 响应式字体大小 */
@media (max-width: 640px) {
  h1 {
    font-size: 1rem;
  }
}
</style>
