// 封装程序的状态处理
// 将整个程序的状态存储在localstorage 并且做成响应式

import { useStorage } from '@vueuse/core'

export const state = useStorage(
    'vblog-store',
    {token: null,
    },
    localStorage,
    // 读取localstorage 里面的vblog-store的值作为参数 一起合并
    { mergeDefaults: true } 
  )