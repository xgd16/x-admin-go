<script setup lang="ts">
</script>

<template>
  <div class="dashboard-page">
    <h2
      v-motion
      class="page-title mb-3 text-gray-800 dark:text-gray-100"
      :initial="{ opacity: 0, y: -8 }"
      :enter="{ opacity: 1, y: 0, transition: { duration: 0.35 } }"
    >
      仪表盘
    </h2>

    <el-row :gutter="12">
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card
          v-motion
          class="dashboard-card rounded-card"
          shadow="hover"
          :initial="{ opacity: 0, y: 16 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: 80, duration: 0.4 } }"
          :hovered="{ scale: 1.02, transition: { duration: 0.2 } }"
        >
          <el-statistic title="今日访问" :value="1234" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card
          v-motion
          class="dashboard-card rounded-card"
          shadow="hover"
          :initial="{ opacity: 0, y: 16 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: 160, duration: 0.4 } }"
          :hovered="{ scale: 1.02, transition: { duration: 0.2 } }"
        >
          <el-statistic title="用户总数" :value="5678" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card
          v-motion
          class="dashboard-card rounded-card"
          shadow="hover"
          :initial="{ opacity: 0, y: 16 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: 240, duration: 0.4 } }"
          :hovered="{ scale: 1.02, transition: { duration: 0.2 } }"
        >
          <el-statistic title="订单数" :value="901" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card
          v-motion
          class="dashboard-card rounded-card"
          shadow="hover"
          :initial="{ opacity: 0, y: 16 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: 320, duration: 0.4 } }"
          :hovered="{ scale: 1.02, transition: { duration: 0.2 } }"
        >
          <el-statistic title="待处理" :value="23" />
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="12" class="mt-4">
      <el-col :span="24">
        <el-card
          v-motion
          class="rounded-card"
          shadow="hover"
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { delay: 400, duration: 0.4 } }"
          :hovered="{ scale: 1.01, transition: { duration: 0.2 } }"
        >
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
  padding: 12px;
}

.page-title {
  font-size: 1.25rem;
  font-weight: 600;
}
</style>
