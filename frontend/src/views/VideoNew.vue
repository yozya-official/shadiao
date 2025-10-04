<template>
  <!-- 页面标题 -->
  <PageHeader title="添加精彩作品" subtitle="记录和分享你发现的优秀沙雕动画">
    <template #icon>
      <svg class="w-12 h-12" fill="white" viewBox="0 0 20 20">
        <path
          d="M8 12a4 4 0 1 1 8 0 4 4 0 0 1-8 0zM2 5a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V5zM6 5a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1zM6 9a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1zM6 13a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z"
        ></path>
      </svg>
    </template>
  </PageHeader>

  <div class="container mx-auto px-6 py-8 max-w-6xl relative">
    <!-- 提示信息 -->
    <div class="card bg-card rounded-2xl p-5 mb-8 shadow">
      <div class="flex items-center space-x-2 mb-4">
        <div class="p-1.5 bg-gradient-to-br from-primary to-chart-2 rounded-lg shadow-sm">
          <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path
              fill-rule="evenodd"
              d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
        <h3 class="text-base font-semibold bg-gradient-to-r text-primary bg-clip-text">使用提示</h3>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
        <div class="card border backdrop-blur-sm rounded-xl p-4 space-y-2">
          <div class="flex items-center space-x-2 mb-2">
            <svg
              class="w-4 h-4 text-amber-500"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M13 10V3L4 14h7v7l9-11h-7z"
              />
            </svg>
            <span class="text-sm font-semibold text-gray-600 uppercase tracking-wide"
              >快速开始</span
            >
          </div>
          <p class="text-sm">先输入<span class="font-semibold text-primary">作品链接</span></p>
          <p class="text-sm">
            接着使用 <span class="font-semibold text-primary">"自动生成"</span> 快速填充信息
          </p>
          <p class="text-xs">支持 BV号 · av号 · b23.tv短链</p>
        </div>

        <div class="bg-card card border-1 backdrop-blur-sm rounded-xl p-4 space-y-2">
          <div class="flex items-center space-x-2 mb-2">
            <svg
              class="w-4 h-4 text-green-500"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <span class="text-sm font-semibold uppercase tracking-wide">投稿须知</span>
          </div>
          <p class="text-sm">
            未完结作品请添加 <span class="font-semibold text-primary">第1集</span> (帮作者引引流)
          </p>
          <p class="text-sm">提交后会进入审核</p>
        </div>
      </div>
    </div>

    <!-- 主内容区域 -->
    <div class="grid grid-cols-1 xl:grid-cols-4 gap-8">
      <!-- 表单区域 -->
      <div class="xl:col-span-3 space-y-8">
        <!-- 基本信息卡片 -->
        <div
          class="bg-card/90 backdrop-blur-sm border rounded-3xl shadow p-8 hover:shadow-lg transition-all duration-300"
        >
          <div class="flex items-center justify-between mb-6">
            <div class="flex items-center">
              <div
                class="w-4 h-4 bg-gradient-to-br from-primary to-chart-2 rounded-full mr-4"
              ></div>
              <h3 class="font-bold">基本信息</h3>
            </div>
            <button
              @click="handleAutoFill"
              class="btn border bg-gradient-to-r from-chart-2 to-chart-3 text-primary-foreground rounded-xl transition-all duration-200 font-semibold flex items-center group"
            >
              <svg
                class="w-5 h-5 group-hover:rotate-12 transition-transform duration-300"
                fill="none"
                stroke="white"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M13 10V3L4 14h7v7l9-11h-7z"
                ></path>
              </svg>
              <span>自动生成</span>
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <!-- 作品名称 -->
            <div class="space-y-3">
              <label class="flex items-center space-x-2 font-semibold">
                <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M4 3a2 2 0 100 4h12a2 2 0 100-4H4z"></path>
                  <path
                    fill-rule="evenodd"
                    d="M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
                <span>作品名称</span>
                <span class="text-destructive text-sm">*</span>
              </label>
              <input
                v-model="formData.video.title"
                type="text"
                placeholder="输入作品名称..."
                class="w-full input input-primary"
              />
            </div>

            <!-- 作者UID -->
            <div class="space-y-3">
              <label class="flex items-center space-x-2 font-semibold">
                <svg class="w-5 h-5 text-chart-2" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
                <span>作者UID</span>
                <span class="text-destructive text-sm">*</span>
              </label>
              <input
                v-model.number="formData.author.uid"
                type="text"
                placeholder="输入作者UID..."
                class="w-full input input-primary"
              />
            </div>

            <!-- 链接 -->
            <div class="space-y-3 col-span-full">
              <label class="flex items-center space-x-2 font-semibold">
                <svg
                  class="w-5 h-5 text-chart-3"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"
                  ></path>
                </svg>
                <span>作品链接</span>
                <span class="text-destructive text-sm">*</span>
              </label>
              <input
                v-model="formData.video.url"
                type="url"
                placeholder="输入B站视频链接..."
                class="w-full input input-primary"
              />
            </div>
            <!-- 封面 -->
            <div class="space-y-3 col-span-full">
              <label class="flex items-center space-x-2 font-semibold">
                <svg
                  class="w-5 h-5 text-chart-3"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"
                  ></path>
                </svg>
                <span>作品封面</span>
                <span class="text-destructive text-sm">*</span>
              </label>
              <input
                v-model="formData.video.cover"
                type="url"
                placeholder="输入B站视频封面链接..."
                class="w-full input input-primary"
              />
            </div>

            <div
              v-if="formData.video.cover"
              class="relative group w-full max-w-md mx-auto rounded-2xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300"
            >
              <img
                :src="formData.video.cover"
                referrerpolicy="no-referrer"
                alt="封面预览"
                class="w-full h-60 object-cover transform group-hover:scale-105 transition-transform duration-300"
              />

              <!-- 渐变遮罩 -->
              <div
                class="absolute inset-0 bg-gradient-to-t from-background/70 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"
              ></div>

              <!-- 删除按钮 -->
              <button
                @click.stop="formData.video.cover = ''"
                class="absolute z-20 top-3 right-3 w-9 h-9 flex items-center justify-center rounded-full bg-destructive/90 text-destructive-foreground opacity-0 group-hover:opacity-100 hover:bg-destructive hover:scale-110 transition-all duration-200 shadow-lg backdrop-blur-sm"
                title="删除封面"
                aria-label="删除封面"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12"
                  ></path>
                </svg>
              </button>
            </div>

            <!-- 描述 -->
            <div class="space-y-3 col-span-full">
              <label class="flex items-center space-x-2 font-semibold">
                <svg
                  class="w-5 h-5 text-chart-4"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z"
                  ></path>
                </svg>
                <span>推荐理由</span>
              </label>
              <textarea
                v-model="formData.video.description"
                placeholder="分享你的推荐理由..."
                rows="4"
                class="w-full textarea rounded-xl input-primary resize-none"
              />
            </div>

            <!-- 时长和播放量 -->
            <div class="space-y-3">
              <label class="flex items-center space-x-2 font-semibold">
                <svg class="w-5 h-5 text-chart-5" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
                <span>时长</span>
              </label>
              <select
                v-model.number="formData.video.duration"
                class="w-full select select-primary rounded-xl"
              >
                <option :value="k" v-for="(v, k) in durationOptions" :key="k">{{ v }}</option>
              </select>
            </div>

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
                <span>播放量</span>
              </label>
              <select
                v-model.number="formData.video.views"
                class="w-full select select-primary rounded-xl"
              >
                <option :value="k" v-for="(v, k) in viewsOptions" :key="k">{{ v }}</option>
              </select>
            </div>
          </div>
        </div>

        <!-- 状态设置卡片 -->
        <div
          class="bg-card/90 backdrop-blur-sm border rounded-3xl shadow p-8 hover:shadow-lg transition-all duration-300"
        >
          <div class="flex items-center mb-6">
            <div class="w-4 h-4 bg-gradient-to-br from-chart-2 to-chart-3 rounded-full mr-4"></div>
            <h3 class="font-bold">状态设置</h3>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <!-- 是否原创 -->
            <label
              class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-2xl p-3 border transition-all duration-200 flex items-center justify-between"
            >
              <div class="flex items-center space-x-4">
                <div
                  class="size-6 rounded-xl bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 transition-colors duration-200"
                >
                  <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                    ></path>
                  </svg>
                </div>
                <div>
                  <span class="font-semibold">原创作品</span>
                  <p class="text-muted-foreground text-sm">非搬运，原创内容</p>
                </div>
              </div>
              <input
                v-model="formData.video.isOriginal"
                type="checkbox"
                class="w-5 h-5 text-primary border-2 rounded focus:ring-2 focus:ring-primary/50 transition-all duration-200"
              />
            </label>

            <!-- 是否完结 -->
            <label
              class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-2xl p-3 border transition-all duration-200 flex items-center justify-between"
            >
              <div class="flex items-center space-x-4">
                <div
                  class="w-10 h-10 rounded-xl bg-chart-2/10 flex items-center justify-center group-hover:bg-chart-2/20 transition-colors duration-200"
                >
                  <svg class="w-5 h-5 text-chart-2" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                      clip-rule="evenodd"
                    ></path>
                  </svg>
                </div>
                <div>
                  <span class="font-semibold">已完结</span>
                  <p class="text-muted-foreground text-sm">系列动画已完成</p>
                </div>
              </div>
              <input
                v-model="formData.video.isCompleted"
                type="checkbox"
                class="w-5 h-5 text-chart-2 border-2 rounded focus:ring-2 focus:ring-chart-2/50 transition-all duration-200"
              />
            </label>
          </div>
        </div>

        <!-- 分类设置卡片 -->
        <div
          class="bg-card/90 backdrop-blur-sm border rounded-3xl shadow p-8 hover:shadow-lg transition-all duration-300"
        >
          <div class="flex items-center mb-6">
            <div class="w-4 h-4 bg-gradient-to-br from-chart-3 to-chart-4 rounded-full mr-4"></div>
            <h3 class="font-bold">分类设置</h3>
          </div>

          <div class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <!-- 背景设定 -->
              <div class="space-y-4">
                <label class="flex items-center space-x-2 font-semibold">
                  <svg class="size-5 text-chart-3" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                      clip-rule="evenodd"
                    ></path>
                  </svg>
                  <span>背景设定</span>
                  <span class="text-sm font-normal">(架空背景也算)</span>
                </label>
                <select
                  v-model="formData.video.background"
                  class="w-full select input input-primary"
                >
                  <option value="">选择背景设定...</option>
                  <option
                    v-for="(icon, background) in backgroundOptions"
                    :key="background"
                    :value="background"
                  >
                    {{ icon }} {{ background }}
                  </option>
                </select>
              </div>

              <!-- 世界设定 -->
              <div class="space-y-4">
                <label class="flex items-center space-x-2 font-semibold">
                  <svg class="size-5 text-chart-4" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                      clip-rule="evenodd"
                    ></path>
                  </svg>
                  <span>世界设定</span>
                </label>
                <select v-model="formData.video.world" class="w-full select input input-primary">
                  <option value="">选择世界设定...</option>
                  <option v-for="(icon, option) in worldOptions" :key="option" :value="option">
                    {{ icon }} {{ option }}
                  </option>
                </select>
              </div>
            </div>

            <!-- 风格选择 -->

            <label class="flex items-center space-x-2 font-semibold">
              <svg class="size-5 text-chart-5" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M7 4a3 3 0 016 0v4a3 3 0 11-6 0V4zm4 10.93A7.001 7.001 0 0017 8a1 1 0 10-2 0A5 5 0 015 8a1 1 0 00-2 0 7.001 7.001 0 006 6.93V17H6a1 1 0 100 2h8a1 1 0 100-2h-3v-2.07z"
                  clip-rule="evenodd"
                ></path>
              </svg>
              <span>动画风格（可多选）</span>
            </label>

            <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3">
              <label
                v-for="(icon, style) in styleOptions"
                :key="style"
                class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-xl p-2 border transition-all duration-200 flex items-center justify-between"
                :class="formData.video.style.includes(style) ? 'border-primary bg-primary/10' : ''"
              >
                <div class="flex items-center space-x-3">
                  <span class="">{{ icon }}</span>
                  <span class="font-medium">{{ style }}</span>
                </div>
                <input
                  v-model="formData.video.style"
                  type="checkbox"
                  :value="style"
                  class="w-4 h-4 text-primary border-2 rounded focus:ring-2 focus:ring-primary/50 transition-all duration-200"
                />
              </label>
            </div>
            <!-- 系统外挂 -->
            <label
              class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-2xl p-3 border transition-all duration-200 flex items-center justify-between"
              :class="formData.video.hasSystem ? 'border-primary bg-primary/10' : ''"
            >
              <div class="flex items-center space-x-4">
                <div
                  class="w-10 h-10 rounded-xl bg-chart-3/10 flex items-center justify-center group-hover:bg-chart-3/20 transition-colors duration-200"
                >
                  <svg class="w-5 h-5 text-chart-3" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                      clip-rule="evenodd"
                    ></path>
                  </svg>
                </div>
                <div>
                  <span class="font-semibold">系统外挂</span>
                  <p class="text-muted-foreground text-sm">主角拥有系统、金手指等外挂能力</p>
                </div>
              </div>
              <input
                v-model="formData.video.hasSystem"
                type="checkbox"
                class="w-5 h-5 text-chart-3 border-2 rounded focus:ring-2 focus:ring-chart-3/50 transition-all duration-200"
              />
            </label>
          </div>
        </div>
      </div>

      <!-- 预览和操作区域 -->
      <div class="xl:col-span-1">
        <div class="sticky top-8 space-y-6">
          <!-- 操作按钮 -->
          <div class="bg-card/90 backdrop-blur-sm border rounded-2xl shadow p-3">
            <h3 class="font-bold mb-6 flex items-center">
              <svg class="w-5 h-5 text-primary mr-2" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 11-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 012 0v1.586l2.293-2.293a1 1 0 111.414 1.414L6.414 15H8a1 1 0 010 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 010-2h1.586l-2.293-2.293a1 1 0 111.414-1.414L15 13.586V12a1 1 0 011-1z"
                  clip-rule="evenodd"
                ></path>
              </svg>
              操作
            </h3>
            <div class="space-y-4">
              <button
                @click="handleSubmit"
                :disabled="!formData.video.title || !formData.author.uid || !formData.video.url"
                class="w-full btn btn-lg btn-primary border text-primary-foreground shadow hover:shadow-lg disabled:cursor-not-allowed transition-all duration-200 font-semibold flex items-center justify-center space-x-2"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"
                  ></path>
                </svg>
                <span>保存作品</span>
              </button>
              <button
                @click="handleReset"
                class="w-full btn btn-lg border bg-muted/40 hover:bg-muted/60 text-muted-foreground transition-all duration-200 font-medium flex items-center justify-center space-x-2"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
                  ></path>
                </svg>
                <span>重新填写</span>
              </button>
            </div>
          </div>

          <!-- 表单预览 -->
          <div class="bg-card/60 backdrop-blur-sm border rounded-2xl shadow p-3">
            <h3 class="font-bold mb-4 flex items-center">
              <svg class="w-5 h-5 text-chart-2 mr-2" fill="currentColor" viewBox="0 0 20 20">
                <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"></path>
                <path
                  fill-rule="evenodd"
                  d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                  clip-rule="evenodd"
                ></path>
              </svg>
              表单预览
            </h3>
            <div class="space-y-3 text-sm">
              <div class="flex justify-between">
                <span class="text-muted-foreground">作品名称:</span>
                <span class="font-medium">{{ formData.video.title || '未填写' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-muted-foreground">作者UID:</span>
                <span class="font-medium">{{ formData.author.uid || '未填写' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-muted-foreground">背景设定:</span>
                <span class="font-medium">{{ formData.video.background || '未选择' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-muted-foreground">世界设定:</span>
                <span class="font-medium">{{ formData.video.world || '未选择' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-muted-foreground">风格数量:</span>
                <span class="font-medium">{{ formData.video.style.length }}个</span>
              </div>
              <div class="pt-2 border-t">
                <div class="flex items-center space-x-4 text-xs">
                  <span
                    :class="formData.video.isOriginal ? 'text-primary' : 'text-muted-foreground'"
                  >
                    {{ formData.video.isOriginal ? '✓' : '○' }} 原创
                  </span>
                  <span
                    :class="formData.video.isCompleted ? 'text-chart-2' : 'text-muted-foreground'"
                  >
                    {{ formData.video.isCompleted ? '✓' : '○' }} 完结
                  </span>
                  <span
                    :class="formData.video.hasSystem ? 'text-chart-3' : 'text-muted-foreground'"
                  >
                    {{ formData.video.hasSystem ? '✓' : '○' }} 系统
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from '@yuelioi/toast'

import {
  backgroundOptions,
  styleOptions,
  durationOptions,
  viewsOptions,
  worldOptions,
} from '@/stores/options'

const formData = reactive<VideoNewData>({
  author: {
    uid: 0,
    name: '',
    avatar: '',
  },
  video: {
    title: '',
    aid: 0,
    author: {
      uid: 0,
      name: '',
      avatar: '',
    },
    url: '',
    cover: '',
    duration: 0,
    description: '',
    views: 0,
    isOriginal: false,
    isCompleted: false,
    world: '',
    background: '',
    style: [],
    hasSystem: false,
  },
})

const handleSubmit = async () => {
  if (!formData.video.title || !formData.video.url || !formData.author.uid) {
    toast.error('请填写作品名称/网址/作者UID')
    return
  }

  try {
    const response = await videoApi.createVideo(formData)
    console.log(response)
    toast.success('作品保存成功！')
    handleReset()
  } catch (error: unknown) {
    handleApiError(error, '创建')
  }
}

const handleAutoFill = async () => {
  if (!formData.video.url) {
    toast.error('请先输入B站视频链接')
    return
  }
  toast.info('生成中, 请稍后...', {
    duration: 2000,
  })

  try {
    const resp = await videoApi.getBiliBiliVideoInfo(formData.video.url)
    const data = resp.data.data

    formData.author = {
      uid: data.author.mid,
      name: data.author.name,
      avatar: data.author.avatar,
    }
    formData.video.aid = Number(data.parse_params.id)
    formData.video.url = 'https://www.bilibili.com/video/av' + formData.video.aid
    formData.video.cover = data.item.cover
    formData.video.title = data.item.title
    formData.video.ctime = new Date(data.stat.time * 1000).toISOString()

    toast.success('自动填充成功！')
  } catch (err: unknown) {
    handleApiError(err, '自动填充失败')
  }
}

const handleReset = () => {
  Object.assign(formData.video, {
    title: '',
    aid: 0,
    url: '',
    cover: '',
    description: '',
    duration: 0,
    views: 0,
  })

  Object.assign(formData.author, {
    uid: 0,
    name: '',
    avatar: '',
  })
}
</script>
