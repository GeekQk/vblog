import client from './client'

export const LIST_BLOG = (params) =>
  client({
    url: '/vblog/api/v1/blogs/',
    method: 'get',
    params: params
  })

export const GET_BLOG = (id, params) =>
  client({
    url: `/vblog/api/v1/blogs/${id}`,
    method: 'get',
    params: params
  })

export const CRATE_BLOG = (data) =>
  client({
    url: '/vblog/api/v1/blogs/',
    method: 'post',
    data: data
  })

export const UPDATE_BLOG = (id, data) =>
  client({
    url: `/vblog/api/v1/blogs/${id}`,
    method: 'patch',
    data: data
  })