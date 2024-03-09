# ui

Vblog 前端UI

## 关于设置新淘宝源

```sh
$ yrm add tn https://registry.npmmirror.com

$ yrm use tn
```

## 安装UI插件

[](https://arco.design/vue/docs/start)


## 布局

几个布局
+ 登录页面 (/login)
+ 后台页面 (/backend/xxx)
+ 前台页面 (/frontend/xxx)

## 登录页面

### 前后端对接

熟悉的API 客户端调用方式(Postman):  [axios](http://www.axios-js.com/)
+ post('') 
+ get('')

```sh
npm install axios
```

```js
axios.post('/user', {
    firstName: 'Fred',
    lastName: 'Flintstone'
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
```



cros中间件： 允许那个网站来访问你的后端API
+ 部署策略(非跨越场景)  /xxxx.com/api/ ---> backend  /xxxx.com/* ---> frontend
  + 前端是单独开发(需要一个代理)

![](./docs/proxy.drawio)

配置 vite服务器 的代理设置
```json
  server: {
    proxy: {
      // string shorthand: http://localhost:5173/foo -> http://localhost:4567/foo
      // '/foo': 'http://localhost:4567',
      // /vblog/api/v1 --> http://127.0.0.1:8080/vblog/api/v1
      '/vblog/api/v1': 'http://127.0.0.1:8080'
    }
  }
```

非代理模式(后端处理跨越):
```js
axios.post('http://127.0.0.1:8080/vblog/api/v1/tokens/', form.value)
```
代理模式(前端处理跨越):
```js
axios.post('/vblog/api/v1/tokens/', form.value)
```


### 关于axios

+ 临时创建一个http 实例进行请求: axios.post(), create http client
+ 复用一个http 

```
client = create()
client.get()
```

// CRSF

## 后台


### 后台布局与嵌套路由

```json
{
  path: '/backend',
  name: 'BackendLayout',
  component: BackenLayout,
  children: [
    {
      // blogs 相对路径 /backend/blogs
      // /blogs 绝对路径
      path: 'blogs/list',
      name: 'BackendListBlog',
      component: () => import('../views/backend/blog/ListView.vue'),
    },
    {
      //   /backend/blogs/22  id=22
      path: 'blogs/detail/:id',
      name: 'BackendDetailBlog',
      component: () => import('../views/backend/blog/DetailView.vue'),
    },
    {
      //   /backend/blogs/22  id=22
      path: 'blogs/edit/:id',
      name: 'BackendEditBlog',
      component: () => import('../views/backend/blog/EditView.vue'),
    },
  ],
}
```

### 登录后调整到后台




## 前台