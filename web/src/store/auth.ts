import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getToken, setToken, removeToken } from '@/utils/request'
import type { Response } from '@/utils/request'
import req from '@/utils/request'

export interface UserInfo {
  id: number
  username: string
  nickname: string
  avatar?: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref(getToken())
  const userInfo = ref<UserInfo | null>(null)

  const isLoggedIn = computed(() => !!token.value)

  function setAuth(t: string) {
    token.value = t
    setToken(t)
  }

  function clearAuth() {
    token.value = ''
    userInfo.value = null
    removeToken()
  }

  async function fetchUserInfo(): Promise<UserInfo | null> {
    if (!token.value) return null
    try {
      const res = await req<Response<UserInfo>>({ url: '/auth/info', method: 'get' })
      if (res.code === 0 && res.data) {
        const info = res.data as unknown as UserInfo
        userInfo.value = info
        return info
      }
    } catch {
      clearAuth()
    }
    return null
  }

  return {
    token,
    userInfo,
    isLoggedIn,
    setAuth,
    clearAuth,
    fetchUserInfo,
  }
})
