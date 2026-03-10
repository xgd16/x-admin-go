<script setup lang="ts">
import { ref } from 'vue'
import { useDarkMode } from '@/composables/useDarkMode'
import AppHeader from '@/components/AppHeader.vue'
import AppSidebar from '@/components/AppSidebar.vue'

const { isDark, toggleDark } = useDarkMode()
const sidebarCollapsed = ref(false)
</script>

<template>
  <el-container class="admin-layout min-h-screen">
    <el-aside
      :width="sidebarCollapsed ? '64px' : '220px'"
      class="admin-aside transition-all duration-300"
    >
      <AppSidebar :collapsed="sidebarCollapsed" />
    </el-aside>

    <el-container direction="vertical" class="admin-main-wrapper overflow-hidden">
      <AppHeader
        :collapsed="sidebarCollapsed"
        :dark-mode="isDark"
        @toggle-sidebar="sidebarCollapsed = !sidebarCollapsed"
        @toggle-dark="toggleDark()"
      />

      <el-main class="admin-main bg-layer-1">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<style scoped>
.admin-layout {
  background: var(--bg-layer-2);
  gap: 16px;
  padding: 16px;
}

.admin-aside {
  background: var(--bg-layer-0) !important;
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.06);
  border-radius: var(--radius-panel) !important;
  overflow: hidden;
}

html.dark .admin-aside {
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.2);
}

.admin-main-wrapper {
  background: var(--bg-layer-1);
  gap: 16px;
  flex: 1;
  min-width: 0;
  border-radius: var(--radius-panel) !important;
  overflow: hidden;
}

.admin-main {
  background: var(--bg-layer-1) !important;
  padding: 20px;
  margin: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
