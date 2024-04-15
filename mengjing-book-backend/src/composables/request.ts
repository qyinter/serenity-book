import axios from 'axios'
import type { AxiosInstance, AxiosResponse, InternalAxiosRequestConfig } from 'axios'
import type { Result } from '@/types/global'
import { computed } from 'vue'
import { createDiscreteApi, type MessageProviderProps } from 'naive-ui'
const configProviderPropsRef = computed<MessageProviderProps>(() => ({
  themeOverrides: {
    textColorError: '#f56c6c',
    colorError: '#fef0f0',
    iconColorError: '#f56c6c',
    borderRadius: '12px',
    iconMargin: '0 4px 0 0'
  }
}))
const { message } = createDiscreteApi(['message'], {
  messageProviderProps: configProviderPropsRef
})

const baseUrl = import.meta.env.VITE_API_URL
const httpService: AxiosInstance = axios.create({
  url: baseUrl,
  timeout: 8000
})

httpService.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    if (config.url && !config.url.startsWith('http') && !config.url.startsWith('https')) {
      if (config.url.startsWith('/')) {
        config.url = config.url.slice(1)
      }
      config.url = baseUrl + config.url
    }
    // config.headers.Authorization = getToken()
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

httpService.interceptors.response.use(
  <T = any>(response: AxiosResponse<Result<T>>): AxiosResponse<Result<T>> | AxiosResponse<Blob> => {
    if (response.data instanceof Blob) {
      return response
    }
    if (response.data.code && response.data.code !== 200) {
      message.error(response.data.msg)
    }
    return response
  },
  function (error) {
    return Promise.reject(error)
  }
)

export const get = <T>(url: string, params = {}, headers = {}): Promise<Result<T>> => {
  return new Promise<Result<T>>((resolve, reject) => {
    httpService({
      url: url,
      method: 'get',
      params: params,
      headers: headers
    })
      .then((response: AxiosResponse<Result<T>>) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export const post = <T>(url: string, params = {}, headers = {}): Promise<Result<T>> => {
  return new Promise<Result<T>>((resolve, reject) => {
    httpService({
      url: url,
      method: 'post',
      data: params,
      headers: headers
    })
      .then((response: AxiosResponse<Result<T>>) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export const put = <T>(url: string, params = {}): Promise<Result<T>> => {
  return new Promise<Result<T>>((resolve, reject) => {
    httpService({
      url: url,
      method: 'put',
      data: params
    })
      .then((response: AxiosResponse<Result<T>>) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export const del = <T>(url: string): Promise<Result<T>> => {
  return new Promise<Result<T>>((resolve, reject) => {
    httpService({
      url: url,
      method: 'delete'
    })
      .then((response: AxiosResponse<Result<T>>) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export const fileUpload = <T>(url: string, params = {}): Promise<Result<T>> => {
  return new Promise((resolve, reject) => {
    httpService({
      url: url,
      method: 'post',
      data: params,
      headers: { 'Content-Type': 'multipart/form-data' }
    })
      .then((response: AxiosResponse<Result<T>>) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export const fileDownload = (url: string, params = {}): Promise<Blob> => {
  return new Promise((resolve, reject) => {
    httpService({
      url: url,
      method: 'post',
      data: params,
      responseType: 'blob'
    })
      .then((response: AxiosResponse<Blob>) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export default {
  get,
  post,
  put,
  del,
  fileDownload,
  fileUpload
}
