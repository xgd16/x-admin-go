import type { Response } from '@/utils/request'
import req from '@/utils/request'

export interface UserItem {
  id: number
  username: string
  nickname: string
  avatar: string
  status: number
  createTime?: string
}

export interface UserListRes {
  list: UserItem[]
  total: number
}

export interface CreateUserReq {
  username: string
  password: string
  nickname?: string
  avatar?: string
  status?: number
}

export interface UpdateUserReq {
  id: number
  username?: string
  password?: string
  nickname?: string
  avatar?: string
  status?: number
}

export function getUserList(data: { pageNum?: number; pageSize?: number; username?: string; nickname?: string; status?: number }) {
  return req<Response<UserListRes>>({ url: '/user/list', method: 'post', data })
}

export function getUser(id: number) {
  return req<Response<UserItem>>({ url: '/user/get', method: 'get', params: { id } })
}

export function createUser(data: CreateUserReq) {
  return req<Response<{ id: number }>>({ url: '/user/create', method: 'post', data })
}

export function updateUser(data: UpdateUserReq) {
  return req<Response<unknown>>({ url: '/user/update', method: 'post', data })
}

export function deleteUser(id: number) {
  return req<Response<unknown>>({ url: '/user/delete', method: 'post', data: { id } })
}
