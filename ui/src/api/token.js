import client from './client'

// 封装一个LOGIN API(Restful) ==> LOGIN({})
export const LOGIN = (data) => client({
    url: '/vblog/api/v1/tokens/',
    method: 'post',
    data: data,
})