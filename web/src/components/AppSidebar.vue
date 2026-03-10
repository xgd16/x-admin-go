<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { Odometer, Setting, User } from '@element-plus/icons-vue'
import AppLogo from './AppLogo.vue'

defineProps<{
  collapsed?: boolean
}>()

const router = useRouter()
const route = useRoute()

function handleSelect(path: string) {
  router.push(path)
}
</script>

<template>
  <div class="app-sidebar h-full flex flex-col">
    <AppLogo :collapsed="collapsed" />
    <el-menu
      :default-active="route.path"
      :collapse="collapsed"
      :collapse-transition="false"
      class="flex-1 border-r-0"
      background-color="transparent"
      text-color="var(--el-menu-text-color)"
      active-text-color="var(--el-color-primary)"
      @select="handleSelect"
    >
      <el-menu-item index="/dashboard">
        <el-icon><Odometer /></el-icon>
        <template #title>仪表盘</template>
      </el-menu-item>
      <el-sub-menu index="/system">
        <template #title>
          <el-icon><Setting /></el-icon>
          <span>系统管理</span>
        </template>
        <el-menu-item index="/system/user">
          <el-icon><User /></el-icon>
          <template #title>用户管理</template>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>
