<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/store/auth'
import { login } from '@/api/auth'

const router = useRouter()
const authStore = useAuthStore()

const loading = ref(false)
const form = reactive({
  username: 'admin',
  password: '123456',
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
}

const formRef = ref<InstanceType<typeof import('element-plus').ElForm>>()

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    loading.value = true
    try {
      const res = await login(form)
      const data = res.data as { token?: string; message?: string } | null
      if (res.code === 0 && data?.token) {
        authStore.setAuth(data.token)
        ElMessage.success(data.message ?? '登录成功')
        const redirect = (router.currentRoute.value.query.redirect as string) || '/'
        router.replace(redirect)
      } else {
        ElMessage.error(res.message ?? res.msg ?? '登录失败')
      }
    } catch (err: any) {
      ElMessage.error(err?.message ?? err?.msg ?? '登录失败')
    } finally {
      loading.value = false
    }
  })
}

onMounted(() => {
  if (authStore.isLoggedIn) {
    router.replace((router.currentRoute.value.query.redirect as string) || '/')
  }
})
</script>

<template>
  <div class="login-page min-h-screen flex items-center justify-center bg-layer-2">
    <div class="login-card bg-layer-0 rounded-2xl shadow-lg p-8 w-full max-w-md">
      <div class="text-center mb-8">
        <h1 class="text-2xl font-semibold text-[var(--text-primary)]">x-admin</h1>
        <p class="text-sm text-[var(--text-secondary)] mt-1">后台管理系统</p>
      </div>
      <el-form ref="formRef" :model="form" :rules="rules" label-position="top" @submit.prevent="handleSubmit">
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="form.username"
            placeholder="请输入用户名"
            size="large"
            clearable
            :disabled="loading"
          />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            size="large"
            show-password
            clearable
            :disabled="loading"
            @keyup.enter="handleSubmit"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="large" class="w-full" :loading="loading" @click="handleSubmit">
            登 录
          </el-button>
        </el-form-item>
      </el-form>
      <p class="text-xs text-[var(--text-tertiary)] text-center mt-4">
        默认账号：admin / 123456
      </p>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  padding: 24px;
}
.login-card {
  border-radius: var(--radius-panel, 12px);
}
</style>
