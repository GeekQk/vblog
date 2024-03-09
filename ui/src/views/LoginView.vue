<template>
  <div class="content">
    <!-- Login Form表单 -->
    <a-form class="login-form" :rules="formRules" :model="form" @submit="handleSubmit">
      <div class="title">欢迎登录博客系统</div>
      <a-form-item hide-label field="username" label="">
        <a-input v-model="form.username" placeholder="请输入用户名">
          <template #prefix>
            <icon-user />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item hide-label field="password" label="">
        <a-input-password v-model="form.password" placeholder="请输入用户密码">
          <template #prefix>
            <icon-lock />
          </template>
        </a-input-password>
      </a-form-item>
      <a-form-item hide-label field="remind_me">
        <a-checkbox v-model="form.remind_me"> 记住 </a-checkbox>
      </a-form-item>
      <a-form-item hide-label>
        <a-button type="primary" html-type="submit" style="width: 100%">登录</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { LOGIN } from '../api/token'
import { state } from '../stores/app'
import { useRouter } from 'vue-router'

const router = useRouter()

const form = ref({
  username: '',
  password: '',
  remind_me: false
})
const handleSubmit = (data) => {
  // 调用后端api
  console.log(data)

  if (!data.errors) {
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
}

const formRules = {
  username: [
    {
      required: true,
      message: '请输入用户名'
    }
  ],
  password: [
    {
      required: true,
      message: '请输入密码'
    },
    {
      minLength: 6,
      message: '密码不能少于6位'
    }
  ]
}
</script>

<style lang="css" scoped>
.content {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
}

.login-form {
  height: 100%;
  width: 460px;
  display: flex;
  justify-content: center;
}

.title {
  display: flex;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 12px;
}
</style>
