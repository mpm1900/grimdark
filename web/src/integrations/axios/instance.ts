import { getApiBaseUrl } from '#/utils/get-api-base-url'
import { getRequest } from '@tanstack/react-start/server'
import axios, { AxiosHeaders, type InternalAxiosRequestConfig } from 'axios'

const api = axios.create()

const authInterceptor = (config: InternalAxiosRequestConfig) => {
  config.baseURL ??= getApiBaseUrl()

  const request = getRequest()
  const cookie = request?.headers.get('cookie')
  if (!cookie) {
    return config
  }

  const headers = AxiosHeaders.from(config.headers)
  headers.set('Cookie', cookie)
  config.headers = headers
  return config
}

api.interceptors.request.use(authInterceptor)
api.interceptors.response.use(
  (response) => response,
  (error: any) => {
    error.status = error?.response?.status
    return Promise.reject(error)
  }
)

export { api }
