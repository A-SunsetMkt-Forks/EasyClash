import { ElMessage } from 'element-plus'
export * from './app'
export * from './type'
export * from './env'
export * from './lodash'
export * from './request'
export * from './ws'
export * from './fetch'

export function message(options) {
  ElMessage(options)
}

export function copyToClipboard(value) {
  navigator.clipboard
    .writeText(value)
    .then((_) => {
      message({
        message: '已复制到剪贴板',
        type: 'success',
        duration: 1000
      })
    })
    .catch((err) => {
      message({
        message: '复制到剪贴板失败' + err,
        type: 'error',
        duration: 1000
      })
    })
}

export const mix = (color1, color2, weight) => {
  weight = Math.max(Math.min(Number(weight), 1), 0)
  const r1 = parseInt(color1.substring(1, 3), 16)
  const g1 = parseInt(color1.substring(3, 5), 16)
  const b1 = parseInt(color1.substring(5, 7), 16)
  const r2 = parseInt(color2.substring(1, 3), 16)
  const g2 = parseInt(color2.substring(3, 5), 16)
  const b2 = parseInt(color2.substring(5, 7), 16)
  let r = Math.round(r1 * (1 - weight) + r2 * weight)
  let g = Math.round(g1 * (1 - weight) + g2 * weight)
  let b = Math.round(b1 * (1 - weight) + b2 * weight)
  r = ('0' + (r || 0).toString(16)).slice(-2)
  g = ('0' + (g || 0).toString(16)).slice(-2)
  b = ('0' + (b || 0).toString(16)).slice(-2)
  return '#' + r + g + b
}

export function bytesToSize(bytes) {
  var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  if (!bytes) return '0 Byte'
  var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)) + '')
  return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i]
}
