import client from './client'

export const LIST_BLOG = (params) =>
  client({
    url: '/vblog/api/v1/blogs/',
    method: 'get',
    params: params
  })