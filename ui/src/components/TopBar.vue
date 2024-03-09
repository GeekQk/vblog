<template>
  <div class="top-bar">
    <div>
      <div>老喻的博客</div>
    </div>
    <div class="top-bar-left">
      <a-button type="text" class="full-btn" v-if="isLogin" @click="Logout">退出</a-button>
      <a-button type="text" class="full-btn" v-else @click="Login">登录</a-button>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()

// 获取当前登录状态
import { state } from '@/stores/app'
import { computed } from 'vue'
const isLogin = computed(() => {
  return state.value.token !== null
})
console.log(state.value)

// 退出, 重新登录
const Logout = () => {
  state.value.token = null
  router.push({ name: 'LoginView' })
}

// 登录
const Login = () => {
  router.push({ name: 'LoginView' })
}
</script>

<style lang="css" scoped>
.top-bar {
  padding: 0 12px;
  height: 60px;
  border-bottom: 1px solid rgb(229, 230, 235);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.top-bar-left {
  height: 100%;
  display: flex;
  align-items: center;
}

.full-btn {
  height: 100%;
}
</style>
