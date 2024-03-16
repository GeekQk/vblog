<template>
  <div>
    <!-- 页头 -->
    <a-page-header
      title="文章详情"
      @back="router.go(-1)"
    ></a-page-header>
    <!-- 博客内容 -->
    <a-spin v-if="loadding" />
    <MdPreview v-if="!loadding && content" :modelValue="content"></MdPreview>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { GET_BLOG } from '../../../api/blog'
import { onMounted, ref } from 'vue';
import { MdPreview } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

const router = useRouter()
const blogId=router.currentRoute.value.params.id

// 先请求博客的内容
const query = ref({})
const loadding = ref(true)
const content = ref('')
const GetBlog = async () => {
  loadding.value = true
  try {
    const resp = await GET_BLOG(blogId, query.value)
    content.value = resp.content
  } finally {
    loadding.value = false
  }
}

onMounted(() => {
  GetBlog()
})

// 获取了数据后，说明时候渲染到页面
// 先获取数据，数据准备好了再渲染页面
// 先渲染页面, 在拉去数据(Loading)
</script>

<style lang="css" scoped>
</style>
