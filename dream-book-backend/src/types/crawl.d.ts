import type { PageBean } from './global'
type BookSourceInfo = {
  bookSourceID: string
  bookSourceName: string
  bookSourceURL: string
  enabled: string
  lastUpdateTime: number
  respondTime: number
}

type BookSourceQuery = PageBean & {
  bookSourceName: string
  bookSourceURL: string
}

type BookSourceInfoList = {
  list: BookSourceInfo[]
  total: number
}
export { BookSourceInfo, BookSourceQuery, BookSourceInfoList }
