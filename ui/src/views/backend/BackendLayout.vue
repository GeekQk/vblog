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
import TopBar from '../../components/TopBar.vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const handleMenuItemClick = (key) => {
  // 点那个 页面就切换到那个
  router.push({ name: key })
}
</script>

<style lang="css" scoped>
.main {
  height: calc(100vh - 60px);
  display: flex;
}

.side-bar {
  height: 100%;
  border-right: 1px solid rgb(229, 230, 235);
}

.page {
  padding: 6px;
}
</style>
