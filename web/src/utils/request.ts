import axios, { type AxiosRequestConfig } from 'axios'

const TOKEN_KEY = 'x_admin_token'

export function getToken(): string {
  return localStorage.getItem(TOKEN_KEY) ?? ''
}

export function setToken(token: string): void {
  localStorage.setItem(TOKEN_KEY, token)
}

export function removeToken(): void {
  localStorage.removeItem(TOKEN_KEY)
}

export interface Response<T = any> {
  code: number
  data: T
  message?: string
  msg?: string
}

const instance = axios.create({
  baseURL: '/api',
  timeout: 15000,
})

instance.interceptors.request.use((config) => {
  const token = getToken()
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

instance.interceptors.response.use(
  (res) => {
    const body = res.data
    if (body && typeof body === 'object' && 'code' in body) {
      res.data = body
      return res
    }
    res.data = { code: 0, data: body, message: '' }
    return res
  },
  (err) => {
    const body = err.response?.data as Response | undefined
    const msg = body?.message ?? body?.msg ?? err.message ?? '网络错误'
    if (err.response?.status === 401) {
      removeToken()
      if (!location.hash.includes('#/login')) {
        location.hash = '#/login'
      }
    }
    return Promise.reject({ code: err.response?.status ?? -1, data: null, message: msg, msg })
  }
)

export default function request<T = any>(config: AxiosRequestConfig): Promise<Response<T>> {
  return instance.request(config).then((res) => res.data as Response<T>)
}
