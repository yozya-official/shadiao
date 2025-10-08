<template>
  <!-- 页面标题 -->
  <PageActionHeader
    title="加入我们的社区"
    subtitle="与更多沙雕动画爱好者交流分享，发现更多精彩内容">
    <template #icon>
      <svg class="w-12 h-12" fill="white" viewBox="0 0 20 20">
        <path d="M2 5a2 2 0 012-2h7a2 2 0 012 2v4a2 2 0 01-2 2H9l-3 3v-3H4a2 2 0 01-2-2V5z"></path>
        <path
          d="M15 7v2a4 4 0 01-4 4H9.828l-1.766 1.767c.28.149.599.233.938.233h2l3 3v-3h2a2 2 0 002-2V9a2 2 0 00-2-2h-1z"></path>
      </svg>
    </template>
  </PageActionHeader>
  <div class="container px-6 py-8 relative">
    <!-- 主要内容卡片 -->
    <div
      class="bg-card/90 backdrop-blur-sm border border-border rounded-3xl shadow-xl p-8 mb-8 hover:shadow-2xl transition-all duration-300">
      <!-- QQ群信息 -->
      <div class="text-center space-y-6">
        <div
          class="inline-flex items-center justify-center w-16 rounded-xl text-primary bg-primary/10 mb-4">
          <svg
            t="1759173484906"
            class="icon size-12 p-2"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="4727"
            width="200"
            height="200">
            <path
              d="M946.9952 696.6784c-6.8608-98.8672-70.2464-181.76-106.9056-224.6144 5.12-12.4416 17.4592-84.6336-30.3616-133.888 0.1024-1.1776 0.1024-2.3552 0.1024-3.4816C809.8304 140.4416 677.7344 0.6144 512 0 346.2656 0.6656 214.1696 140.4416 214.1696 334.6944c0 1.1776 0 2.3552 0.1024 3.4816-47.8208 49.2544-35.4304 121.4464-30.3616 133.888-36.608 42.8544-99.9936 125.7472-106.9056 224.6144-1.2288 26.0096 2.6112 63.8976 14.7456 80.7424 14.7968 20.48 55.3472-4.1472 84.3776-69.7344 8.0896 30.4128 26.6752 76.8 68.864 135.6288-70.5536 16.896-90.6752 89.7024-66.9184 129.536 16.7424 28.0576 55.04 51.1488 121.088 51.1488 117.5552 0 169.472-33.024 192.6656-56.064 4.6592-5.0688 11.52-7.4752 20.1728-7.4752 8.704 0 15.5136 2.4064 20.1728 7.4752 23.1936 23.04 75.1104 56.064 192.6144 56.064 66.048 0 104.3968-23.04 121.1392-51.1488 23.7568-39.8336 3.584-112.64-66.9184-129.536 42.1888-58.88 60.8256-105.216 68.864-135.6288 29.0304 65.5872 69.632 90.2656 84.3776 69.7344 12.1344-16.896 15.9744-54.784 14.7456-80.7424z"
              fill="currentColor"
              p-id="4728"></path>
          </svg>
        </div>

        <h2 class="text-2xl text-foreground">沙雕动画俱乐部</h2>

        <p class="text-muted-foreground leading-relaxed max-w-xl mx-auto">
          这里聚集了大量的沙雕动画爱好者，大家一起分享推荐、讨论剧情、交流心得。不管你是资深观众还是新入坑的小伙伴，都欢迎加入！
        </p>

        <!-- 群特色 -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 my-8">
          <div class="p-4 bg-primary/5 rounded-xl border border-primary/10">
            <div class="text-primary text-2xl mb-2">💬</div>
            <h3 class="font-semibold text-foreground mb-1">实时交流</h3>
            <p class="text-sm text-muted-foreground">与群友实时讨论最新剧情</p>
          </div>
          <div class="p-4 bg-chart-2/5 rounded-xl border border-chart-2/10">
            <div class="text-chart-2 text-2xl mb-2">🎬</div>
            <h3 class="font-semibold text-foreground mb-1">资源分享</h3>
            <p class="text-sm text-muted-foreground">第一时间获取优质推荐</p>
          </div>
          <div class="p-4 bg-chart-3/5 rounded-xl border border-chart-3/10">
            <div class="text-chart-3 text-2xl mb-2">🎉</div>
            <h3 class="font-semibold text-foreground mb-1">审核快人一步</h3>
            <p class="text-sm text-muted-foreground">欢迎催审</p>
          </div>
        </div>

        <!-- 加群按钮 -->
        <div class="space-y-4 [&>div]:hover:translate-x-4">
          <a
            href="https://qm.qq.com/q/ajK30VmEPm"
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex hover:-translate-y-1 items-center justify-center px-8 py-4 bg-gradient-to-r from-primary to-chart-2 text-primary-foreground rounded-xl hover:shadow-xl transition-all duration-300 font-bold text-lg space-x-3 group">
            <svg
              class="w-6 h-6 group-hover:scale-110 transition-transform duration-200"
              fill="currentColor"
              viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
                clip-rule="evenodd"></path>
            </svg>
            <span>点击加入QQ群</span>
            <svg
              class="w-5 h-5 group-hover:translate-x-1 transition-transform duration-200"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
            </svg>
          </a>

          <p class="text-sm text-muted-foreground">或手动搜索群号加入:<span>673296047</span></p>
        </div>
      </div>
    </div>

    <!-- 社区规则 -->
    <div
      class="bg-card/90 backdrop-blur-sm border border-border rounded-3xl shadow-xl p-8 hover:shadow-2xl transition-all duration-300">
      <div class="flex items-center mb-6">
        <div class="w-4 h-4 bg-gradient-to-br from-chart-4 to-chart-5 rounded-full mr-4"></div>
        <h2 class="text-2xl font-bold text-foreground">社区公约</h2>
      </div>

      <div class="space-y-4 [&>div]:transition-transform [&>div]:hover:translate-x-2">
        <div
          class="flex items-start space-x-4 p-4 bg-muted/20 rounded-xl hover:bg-muted/30 transition-colors duration-200">
          <div
            class="flex-shrink-0 w-8 h-8 rounded-lg bg-primary/10 flex items-center justify-center text-primary font-bold">
            1
          </div>
          <div>
            <h3 class="font-semibold text-foreground mb-1">友善交流</h3>
            <p class="text-sm text-muted-foreground">尊重每一位成员，文明讨论，理性交流</p>
          </div>
        </div>

        <div
          class="flex items-start space-x-4 p-4 bg-muted/20 rounded-xl hover:bg-muted/30 transition-colors duration-200">
          <div
            class="flex-shrink-0 w-8 h-8 rounded-lg bg-chart-2/10 flex items-center justify-center text-chart-2 font-bold">
            2
          </div>
          <div>
            <h3 class="font-semibold text-foreground mb-1">分享精彩</h3>
            <p class="text-sm text-muted-foreground">积极分享优质内容，帮助大家发现好作品</p>
          </div>
        </div>

        <div
          class="flex items-start space-x-4 p-4 bg-muted/20 rounded-xl hover:bg-muted/30 transition-colors duration-200">
          <div
            class="flex-shrink-0 w-8 h-8 rounded-lg bg-chart-3/10 flex items-center justify-center text-chart-3 font-bold">
            3
          </div>
          <div>
            <h3 class="font-semibold text-foreground mb-1">拒绝广告</h3>
            <p class="text-sm text-muted-foreground">禁止发送无关广告和恶意刷屏信息</p>
          </div>
        </div>

        <div
          class="flex items-start space-x-4 p-4 bg-muted/20 rounded-xl hover:bg-muted/30 transition-colors duration-200">
          <div
            class="flex-shrink-0 w-8 h-8 rounded-lg bg-chart-4/10 flex items-center justify-center text-chart-4 font-bold">
            4
          </div>
          <div>
            <h3 class="font-semibold text-foreground mb-1">保护隐私</h3>
            <p class="text-sm text-muted-foreground">尊重他人隐私，不随意传播个人信息</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { PageActionHeader } from '@yuelioi/ui'
</script>
