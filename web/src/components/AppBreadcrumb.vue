<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

interface BreadcrumbItem {
  title: string
  path: string
}

const breadcrumbList = computed<BreadcrumbItem[]>(() => {
  const result: BreadcrumbItem[] = []
  let path = ''
  for (const r of route.matched) {
    const seg = r.path === '/' ? '' : r.path.replace(/^\/+|\/+$/g, '')
    path = path ? `${path}/${seg}` : r.path || '/'
    if (r.meta?.title) {
      result.push({ title: r.meta.title as string, path })
    }
  }
  return result
})
</script>

<template>
  <el-breadcrumb v-if="breadcrumbList.length" separator="/" class="app-breadcrumb">
    <el-breadcrumb-item v-for="(item, index) in breadcrumbList" :key="item.path">
      <router-link
        v-if="index < breadcrumbList.length - 1"
        :to="item.path"
        class="breadcrumb-link"
      >
        {{ item.title }}
      </router-link>
      <span v-else class="breadcrumb-current">{{ item.title }}</span>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<style scoped>
.app-breadcrumb {
  font-size: 14px;
}

.breadcrumb-link {
  color: var(--el-text-color-secondary);
  text-decoration: none;
  transition: color 0.2s;
}

.breadcrumb-link:hover {
  color: var(--el-color-primary);
}

.breadcrumb-current {
  color: var(--el-text-color-primary);
  font-weight: 500;
}
</style>
