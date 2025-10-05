<script setup lang="ts">
import { RouterView } from 'vue-router'
import { ToastContainer } from '@yuelioi/toast'
import TheHeader from './components/TheHeader.vue'

import TheFooter from '@yuelioi/components/the-footer'
</script>

<template>
  <ToastContainer />
  <TheHeader site-title="沙雕动画收集站" />

  <RouterView v-slot="{ Component }">
    <keep-alive>
      <component :is="Component" />
    </keep-alive>
  </RouterView>

  <TheFooter />
</template>
