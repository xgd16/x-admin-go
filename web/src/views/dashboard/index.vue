<script setup lang="ts">
</script>

<template>
  <div class="dashboard-page">
    <h2 class="page-title mb-6 text-gray-800 dark:text-gray-100">仪表盘</h2>

    <el-row :gutter="16">
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="dashboard-card rounded-card" shadow="hover">
          <el-statistic title="今日访问" :value="1234" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="dashboard-card rounded-card" shadow="hover">
          <el-statistic title="用户总数" :value="5678" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="dashboard-card rounded-card" shadow="hover">
          <el-statistic title="订单数" :value="901" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="dashboard-card rounded-card" shadow="hover">
          <el-statistic title="待处理" :value="23" />
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="16" class="mt-6">
      <el-col :span="24">
        <el-card class="rounded-card" shadow="hover">
          <template #header>欢迎使用 x-admin</template>
          <p class="text-gray-600 dark:text-gray-300">
            管理后台基于 Vue 3 + TypeScript + Vite + Tailwind CSS + Element Plus 构建，
            支持深色模式切换，白色阶梯层级与圆角设计。
          </p>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.dashboard-card {
  background: var(--bg-layer-0);
}

.dashboard-card :deep(.el-card__body) {
  padding: 20px;
}

.page-title {
  font-size: 1.25rem;
  font-weight: 600;
}
</style>
