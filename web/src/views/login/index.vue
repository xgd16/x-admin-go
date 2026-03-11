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
  <div class="login-page">
    <div
      v-motion
      class="login-card"
      :initial="{ opacity: 0, y: 32, scale: 0.95 }"
      :enter="{
        opacity: 1,
        y: 0,
        scale: 1,
        transition: { type: 'spring', stiffness: 90, damping: 18, mass: 0.8 },
      }"
    >
      <div
        v-motion
        class="login-header"
        :initial="{ opacity: 0, y: -8 }"
        :enter="{
          opacity: 1,
          y: 0,
          transition: { delay: 100, duration: 0.4, ease: 'easeOut' },
        }"
      >
        <h1 class="login-title">x-admin</h1>
        <p class="login-subtitle">后台管理系统</p>
      </div>

      <el-form
        ref="formRef"
        class="login-form"
        :model="form"
        :rules="rules"
        label-position="top"
        @submit.prevent="handleSubmit"
      >
        <el-form-item
          v-motion
          label="用户名"
          prop="username"
          :initial="{ opacity: 0, x: -12 }"
          :enter="{ opacity: 1, x: 0, transition: { delay: 180, duration: 0.35, ease: 'easeOut' } }"
        >
          <el-input
            v-model="form.username"
            placeholder="请输入用户名"
            size="large"
            clearable
            :disabled="loading"
          />
        </el-form-item>
        <el-form-item
          v-motion
          label="密码"
          prop="password"
          :initial="{ opacity: 0, x: -12 }"
          :enter="{ opacity: 1, x: 0, transition: { delay: 240, duration: 0.35, ease: 'easeOut' } }"
        >
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
        <el-form-item
          v-motion
          class="login-btn-wrap"
          :initial="{ opacity: 0, y: 16 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: 300, duration: 0.4, ease: 'easeOut' } }"
        >
          <el-button
            v-motion
            type="primary"
            size="large"
            class="login-btn"
            :loading="loading"
            :hovered="{ scale: 1.02, y: -1, transition: { duration: 0.2 } }"
            :tapped="{ scale: 0.98, transition: { duration: 0.1 } }"
            @click="handleSubmit"
          >
            登 录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-layer-2);
  padding: 24px;
}

.login-card {
  width: 100%;
  max-width: 400px;
  background: var(--bg-layer-0);
  border-radius: var(--radius-panel);
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.08);
  padding: 40px;
}

html.dark .login-card {
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.3);
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.login-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
}

.login-subtitle {
  font-size: 0.875rem;
  color: var(--text-secondary);
  margin: 6px 0 0;
}

/* 修复 Element Plus 表单布局 */
.login-form {
  --el-form-label-font-size: 14px;
}

.login-form :deep(.el-form-item) {
  margin-bottom: 24px;
}

.login-form :deep(.el-form-item__label) {
  display: block;
  width: 100% !important;
  margin-bottom: 8px;
  padding: 0 !important;
  line-height: 1.5;
  font-size: 14px;
  color: var(--text-primary);
}

.login-form :deep(.el-form-item__content) {
  display: block;
  margin-left: 0 !important;
  line-height: normal;
}

.login-form :deep(.el-input__wrapper) {
  width: 100%;
}

.login-btn-wrap {
  margin-bottom: 0;
}

.login-btn {
  width: 100%;
}

.login-hint {
  margin: 20px 0 0;
  font-size: 12px;
  color: var(--text-tertiary);
  text-align: center;
}
</style>
