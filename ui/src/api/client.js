// 封装一个统一的http client 实例

import axios from 'axios'
import { Message } from '@arco-design/web-vue'

// 设置实例统一通用配置
var instance = axios.create({
  timeout: 5000
})

// 为实例添加统一的拦截器
instance.interceptors.response.use(
  // 成功处理 (200 ok)
  (response) => {
    return response.data
  },
  // 请求异常这么处理(!200)
  (err) => {
    console.log(err)
    var msg = err.message
    if (err.response.data && err.response.data.message) {
      msg = err.response.data.message
      // 针对特定的异常, 做特殊的逻辑处理
      if (err.response.data.code === 401) {
         window.location.assign('/login')
      }
    }


    // 提示异常
    Message.error({
      content: msg,
      duration: 2000
    })
    // 异常不进行传递也可以
    return Promise.reject(err)
  }
)

export default instance
