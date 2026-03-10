<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { Fold, Expand, Moon, Sunny, User } from '@element-plus/icons-vue'
import IconButton from './IconButton.vue'
import { useAuthStore } from '@/store/auth'

defineProps<{
  collapsed: boolean
  darkMode: boolean
}>()

defineEmits<{
  toggleSidebar: []
  toggleDark: []
}>()

const router = useRouter()
const authStore = useAuthStore()
const nickname = computed(() => authStore.userInfo?.nickname ?? authStore.userInfo?.username ?? '管理员')

function handleLogout() {
  authStore.clearAuth()
  router.push('/login')
}
</script>

<template>
  <el-header class="admin-header flex items-center justify-between">
    <div class="flex items-center gap-2">
      <IconButton
        :content="collapsed ? '展开侧边栏' : '收起侧边栏'"
        @click="$emit('toggleSidebar')"
      >
        <el-icon class="text-lg"><Fold v-if="!collapsed" /><Expand v-else /></el-icon>
      </IconButton>
      <span class="text-sm text-gray-500 dark:text-gray-400">管理后台</span>
    </div>

    <div class="flex items-center gap-4">
      <div class="flex items-center gap-2">
        <el-icon>
          <Sunny v-if="!darkMode" />
          <Moon v-else />
        </el-icon>
        <el-switch
          :model-value="darkMode"
          @update:model-value="$emit('toggleDark')"
          inline-prompt
          active-text="暗"
          inactive-text="亮"
        />
      </div>
      <el-dropdown trigger="click" @command="(cmd: string) => cmd === 'logout' && handleLogout()">
        <div class="flex items-center gap-2 cursor-pointer">
          <el-icon><User /></el-icon>
          <span class="text-sm">{{ nickname }}</span>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </el-header>
</template>

<style scoped>
.admin-header {
  background: var(--bg-layer-0) !important;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.06);
  padding: 0 20px;
}

html.dark .admin-header {
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.2);
}
</style>
