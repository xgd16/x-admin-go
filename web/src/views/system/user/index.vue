<script setup lang="ts">
import { ref } from 'vue'

const tableData = ref([
  { id: 1, name: '管理员', username: 'admin', status: '启用' },
  { id: 2, name: '测试用户', username: 'test', status: '启用' },
])
</script>

<template>
  <div class="user-page">
    <h2 class="page-title mb-6 text-gray-800 dark:text-gray-100">用户管理</h2>

    <el-card class="rounded-card" shadow="hover">
      <template #header>
        <div class="flex justify-between items-center">
          <span>用户列表</span>
          <el-space>
            <el-button type="primary">新增用户</el-button>
            <el-button>导出</el-button>
          </el-space>
        </div>
      </template>
      <el-table :data="tableData" stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="姓名" />
        <el-table-column prop="username" label="用户名" />
        <el-table-column prop="status" label="状态" width="100" />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default>
            <el-space>
              <el-button type="primary" link>编辑</el-button>
              <el-button type="danger" link>删除</el-button>
            </el-space>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<style scoped>
.page-title {
  font-size: 1.25rem;
  font-weight: 600;
}
</style>
