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

export function login(data: LoginReq) {
  return req<Response<LoginRes>>({
    url: '/auth/login',
    method: 'post',
    data,
  })
}
