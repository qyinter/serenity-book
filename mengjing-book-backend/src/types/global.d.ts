type Result<T> = {
  msg: string
  code: number
  data: T
}

type CaptchaContent = {
  base64Img: string
  uuid: string
}

type PageBean = {
  pageNum: number
  pageSize: number
}
export { Result, CaptchaContent, PageBean }
