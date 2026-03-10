import type { Response } from '@/utils/request'
import req from '@/utils/request'

export interface LoginReq {
  username: string
  password: string
}

export interface LoginRes {
  token: string
  expire: number
  message?: string
}

export interface ChangePasswordReq {
  oldPassword: string
  newPassword: string
}

export function login(data: LoginReq) {
  return req<Response<LoginRes>>({
    url: '/auth/login',
    method: 'post',
    data,
  })
}

export function changePassword(data: ChangePasswordReq) {
  return req<Response<{ message?: string }>>({
    url: '/auth/change-password',
    method: 'post',
    data,
  })
}
