<script setup lang="ts">
import { ref, computed, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Fold, Expand, Moon, Sunny, User } from '@element-plus/icons-vue'
import IconButton from './IconButton.vue'
import { useAuthStore } from '@/store/auth'
import { changePassword } from '@/api/auth'

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

const pwdDialogVisible = ref(false)
const pwdLoading = ref(false)
const pwdFormRef = ref<InstanceType<typeof import('element-plus').ElForm>>()
const pwdForm = reactive({ oldPassword: '', newPassword: '', confirmPassword: '' })
const pwdRules = {
  oldPassword: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, max: 32, message: '密码 6-32 字符', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    { validator: (_: any, v: string, cb: (e?: Error) => void) => {
      if (v !== pwdForm.newPassword) return cb(new Error('两次输入的密码不一致'))
      cb()
    }, trigger: 'blur' },
  ],
}

function handleCommand(cmd: string) {
  if (cmd === 'changePassword') {
    pwdForm.oldPassword = ''
    pwdForm.newPassword = ''
    pwdForm.confirmPassword = ''
    pwdDialogVisible.value = true
  } else if (cmd === 'logout') {
    handleLogout()
  }
}

async function handleChangePassword() {
  if (!pwdFormRef.value) return
  await pwdFormRef.value.validate(async (valid) => {
    if (!valid) return
    pwdLoading.value = true
    try {
      const res = await changePassword({
        oldPassword: pwdForm.oldPassword,
        newPassword: pwdForm.newPassword,
      })
      if (res.code === 0) {
        ElMessage.success(res.data?.message ?? '修改成功')
        pwdDialogVisible.value = false
      } else {
        ElMessage.error(res.message ?? res.msg ?? '修改失败')
      }
    } catch (e: any) {
      ElMessage.error(e?.message ?? '修改失败')
    } finally {
      pwdLoading.value = false
    }
  })
}

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
      <el-dropdown trigger="click" @command="handleCommand">
        <div class="flex items-center gap-2 cursor-pointer">
          <el-icon><User /></el-icon>
          <span class="text-sm">{{ nickname }}</span>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="changePassword">修改密码</el-dropdown-item>
            <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </el-header>

  <el-dialog
      v-model="pwdDialogVisible"
      title="修改密码"
      width="380px"
      :close-on-click-modal="false"
    >
      <el-form ref="pwdFormRef" :model="pwdForm" :rules="pwdRules" label-width="110px">
        <el-form-item label="原密码" prop="oldPassword">
          <el-input v-model="pwdForm.oldPassword" type="password" placeholder="请输入原密码" show-password />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input v-model="pwdForm.newPassword" type="password" placeholder="6-32 字符" show-password />
        </el-form-item>
        <el-form-item label="确认新密码" prop="confirmPassword">
          <el-input v-model="pwdForm.confirmPassword" type="password" placeholder="再次输入新密码" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="pwdDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="pwdLoading" @click="handleChangePassword">确定</el-button>
      </template>
    </el-dialog>
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
