<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import Table from '@/components/Table/Table.vue'
import { createUser, updateUser, deleteUser } from '@/api/user'
import type { UserItem } from '@/api/user'

const tableRef = ref<InstanceType<typeof Table>>()

const searchForm = reactive({
  username: '',
  nickname: '',
  status: undefined as number | undefined,
})

const dialogVisible = ref(false)
const dialogLoading = ref(false)
const isEdit = ref(false)
const formRef = ref<InstanceType<typeof import('element-plus').ElForm>>()
const form = reactive({
  id: 0,
  username: '',
  password: '',
  nickname: '',
  status: 1,
})

const formRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 32, message: '用户名 3-32 字符', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur', validator: (_: any, v: string, cb: (e?: Error) => void) => {
      if (isEdit.value && !v) return cb()
      if (!v) return cb(new Error('请输入密码'))
      if (v.length < 6 || v.length > 32) return cb(new Error('密码 6-32 字符'))
      cb()
    }},
  ],
}

function openCreate() {
  isEdit.value = false
  form.id = 0
  form.username = ''
  form.password = ''
  form.nickname = ''
  form.status = 1
  dialogVisible.value = true
}

function openEdit(row: UserItem) {
  isEdit.value = true
  form.id = row.id
  form.username = row.username
  form.password = ''
  form.nickname = row.nickname
  form.status = row.status
  dialogVisible.value = true
}

async function handleSubmit() {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    dialogLoading.value = true
    try {
      if (isEdit.value) {
        const data: any = { id: form.id, nickname: form.nickname, status: form.status }
        if (form.username) data.username = form.username
        if (form.password) data.password = form.password
        const res = await updateUser(data)
        if (res.code === 0) {
          ElMessage.success('更新成功')
          dialogVisible.value = false
          tableRef.value?.refresh()
        } else {
          ElMessage.error(res.message ?? res.msg ?? '更新失败')
        }
      } else {
        const res = await createUser({
          username: form.username,
          password: form.password,
          nickname: form.nickname,
          status: form.status,
        })
        if (res.code === 0) {
          ElMessage.success('创建成功')
          dialogVisible.value = false
          tableRef.value?.refresh()
        } else {
          ElMessage.error(res.message ?? res.msg ?? '创建失败')
        }
      }
    } catch (err: any) {
      ElMessage.error(err?.message ?? '操作失败')
    } finally {
      dialogLoading.value = false
    }
  })
}

function handleDelete(row: UserItem) {
  ElMessageBox.confirm(`确定删除用户「${row.username}」？`, '提示', {
    type: 'warning',
  }).then(async () => {
    try {
      const res = await deleteUser(row.id)
      if (res.code === 0) {
        ElMessage.success('删除成功')
        tableRef.value?.refresh()
      } else {
        ElMessage.error(res.message ?? res.msg ?? '删除失败')
      }
    } catch (err: any) {
      ElMessage.error(err?.message ?? '删除失败')
    }
  }).catch(() => {})
}
</script>

<template>
  <div class="user-page">
    <h2 class="page-title mb-3 text-gray-800 dark:text-gray-100">用户管理</h2>

    <el-card class="page-card rounded-card" shadow="hover">
      <template #header>
        <div class="flex justify-between items-center">
          <span>用户列表</span>
          <el-button type="primary" @click="openCreate">新增用户</el-button>
        </div>
      </template>

      <div class="flex flex-wrap gap-3 mb-3">
        <el-input v-model="searchForm.username" placeholder="用户名" clearable style="width: 160px" />
        <el-input v-model="searchForm.nickname" placeholder="昵称" clearable style="width: 160px" />
        <el-select v-model="searchForm.status" placeholder="状态" clearable style="width: 120px">
          <el-option label="启用" :value="1" />
          <el-option label="禁用" :value="0" />
        </el-select>
      </div>

      <Table
        ref="tableRef"
        url="/user/list"
        method="post"
        :search="[searchForm]"
        :page-num-key="'pageNum'"
        :page-size-key="'pageSize'"
      >
        <template #columns>
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="username" label="用户名" min-width="100" />
          <el-table-column prop="nickname" label="昵称" min-width="100" />
          <el-table-column prop="status" label="状态" width="90">
            <template #default="{ row }">{{ row.status === 1 ? '启用' : '禁用' }}</template>
          </el-table-column>
          <el-table-column prop="createTime" label="创建时间" width="170" />
          <el-table-column label="操作" width="160" fixed="right">
            <template #default="{ row }">
              <el-button type="primary" link @click="openEdit(row)">编辑</el-button>
              <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
            </template>
          </el-table-column>
        </template>
      </Table>
    </el-card>

    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑用户' : '新增用户'"
      width="420px"
      destroy-on-close
      @close="formRef?.resetFields()"
    >
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="3-32字符" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" :placeholder="isEdit ? '不修改留空' : '6-32字符'" show-password />
        </el-form-item>
        <el-form-item label="昵称">
          <el-input v-model="form.nickname" placeholder="选填" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="dialogLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.page-title {
  font-size: 1.25rem;
  font-weight: 600;
}
.page-card :deep(.el-card__body) {
  padding: 12px;
}
</style>
