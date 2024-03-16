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

```js
    // 只有当前端校验通过 才需要向后端提交数据
    // axios 调用API请求
    // 临时创建一个http 实例进行请求
    LOGIN(form.value).then((response) => {
      // 把返回的Token对象保持到浏览器, 选用浏览器存储: localstorage
      // 需要一个比较集中管理 应用状态的存储方案: app {token: '', isLogin}
      // localstorage 本书不是响应式的, 有没有办法把localstorage 做成响应式
      state.value.token = response

      // 跳转到后台管理页面: BackendListBlog
      // vue router库提供 router对象的获取方法
      router.push({ name: 'BackendListBlog' })
    })
  }
```

### 后台侧边栏导航

```vue
<template>
  <div>
    <!-- 顶部导航 -->
    <TopBar></TopBar>
    <!-- 内容区 -->
    <div class="main">
      <!-- 侧边栏导航 -->
      <div class="side-bar">
        <a-menu
          @menu-item-click="handleMenuItemClick"
          :style="{ width: '200px', height: '100%' }"
          :default-open-keys="['blog']"
          :default-selected-keys="['blog_list']"
          show-collapse-button
          breakpoint="xl"
        >
          <a-sub-menu key="blog">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>文章管理</template>
            <a-menu-item key="BackendListBlog">文章列表</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="comment">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>评论管理</template>
            <a-menu-item key="BackendListComment">评论列表</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </div>
      <!-- 业务页面 -->
      <div class="page">
        <RouterView></RouterView>
      </div>
    </div>
  </div>
</template>

<script setup>
import TopBar from '../../components/TopBar.vue';
import { useRouter } from 'vue-router';

const router = useRouter()

const handleMenuItemClick = (key) => {
  // 点那个 页面就切换到那个
  router.push({name: key})
}
</script>
```

### TopBar状态处理

+ 登录/退出

```js
// 获取当前登录状态
import { state } from '@/stores/app';
import { computed } from 'vue';
const isLogin = computed(() => {
  return state.value.token !== null
})
console.log(state.value);

// 退出, 重新登录
const Logout = () => {
  state.value.token = null
  router.push({name: 'LoginView'})
}

// 登录
const Login = () => {
  router.push({name: 'LoginView'})
}
```

### 守卫导航

如果用户绕开你的登录页面 访问你后台页面 应该如何防卫

补充全局前置守卫
```js
router.beforeEach((to, from) => {
  // 是不是访问后台页面
  if (to.fullPath.startsWith('/backend/')) {
    // 判断是否登录
    if (!state.value.token) {
      // 跳转去登录页面
      return {name: 'LoginView'}
    }
  }
})
```

### 文章列表页

+ 面包屑
+ 表格操作区
+ 表格数据
+ 分页数据

试图
```vue
<template>
  <div>
    <div>
      <a-breadcrumb>
        <a-breadcrumb-item>文章管理</a-breadcrumb-item>
        <a-breadcrumb-item>文章列表</a-breadcrumb-item>
      </a-breadcrumb>
    </div>
    <div class="op">
      <div>
        <a-button type="primary">创建文章</a-button>
      </div>
      <div>
        <a-input :style="{ width: '320px' }" placeholder="请输入文字名称敲回车键搜索" allow-clear />
      </div>
    </div>
    <div>
      <a-table :pagination="false" :columns="columns" :data="data" />
    </div>
    <div class="pagi">
      <a-pagination :total="50" show-total show-jumper show-page-size/>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'

const columns = [
  {
    title: 'Name',
    dataIndex: 'name'
  },
  {
    title: 'Salary',
    dataIndex: 'salary'
  },
  {
    title: 'Address',
    dataIndex: 'address'
  },
  {
    title: 'Email',
    dataIndex: 'email'
  }
]
const data = reactive([
  {
    key: '1',
    name: 'Jane Doe',
    salary: 23000,
    address: '32 Park Road, London',
    email: 'jane.doe@example.com'
  },
  {
    key: '2',
    name: 'Alisa Ross',
    salary: 25000,
    address: '35 Park Road, London',
    email: 'alisa.ross@example.com'
  },
  {
    key: '3',
    name: 'Kevin Sandra',
    salary: 22000,
    address: '31 Park Road, London',
    email: 'kevin.sandra@example.com'
  },
  {
    key: '4',
    name: 'Ed Hellen',
    salary: 17000,
    address: '42 Park Road, London',
    email: 'ed.hellen@example.com'
  },
  {
    key: '5',
    name: 'William Smith',
    salary: 27000,
    address: '62 Park Road, London',
    email: 'william.smith@example.com'
  }
])
</script>

<style lang="css" scoped>
.op {
  margin-top: 8px;
  margin-bottom: 8px;
  display: flex;
  justify-content: space-between;
}
.pagi {
  margin-top: 4px;
  display: flex;
  flex-direction: row-reverse;
}
</style>
```

定义 API 获取列表数据

### 关键字搜索 

1. BlogList(page_size,page_number,...)

后端:
+ keywords 关键字查询条件
+ http 请求解析 该用户参数
+ BlogList接口 实现参数的处理

签订:
+ request.keywords  <---> a-input
+ 用户回车时 触发 blog list API调用

### 文章的创建与编辑

独立使用一个页面: EditView.vue, 点击后跳转


+ 页头
+ 文本编辑和创建表单

选一个支持markdown的编辑器 [md-editor-v3](https://www.npmjs.com/package/md-editor-v3)

如何找到UI组件元素，并且调用该组件上的方法: Ref机制: https://cn.vuejs.org/guide/essentials/template-refs.html




### 文章的删除


## 前台