<template>
  <div>
    <!-- 页头 -->
    <a-page-header
      title="文章详情"
      @back="router.go(-1)"
    ></a-page-header>
    <!-- 博客内容 -->
    <a-spin v-if="loadding" />
    <div v-if="!loadding && content">
      <MdPreview class="parent" :editorId="id" :modelValue="content"></MdPreview>
      <MdCatalog class="child" :editorId="id" :scrollElement="scrollElement" />
    </div>
    
    
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { GET_BLOG } from '../../../api/blog'
import { onMounted, ref } from 'vue';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

const id = 'preview-only';
const scrollElement = document.documentElement;

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

// const handleClick = (e, t) => {
//   // 滚动到具有特定ID的元素
//   console.log(document.getElementById(t.text))
//   document.getElementById(t.text).scrollIntoView({
//     behavior: 'smooth', // 可选，平滑滚动
//     block: 'center' // 可选，对齐方式（start, center, end, nearest）
//   });
//   console.log(e, t);
// }
</script>

<style lang="css" scoped>
.parent {
  position: relative;
}

.child {
  position: absolute;
  top: 80px; /* 负值表示向上偏移 */
  right: 0;
  background-color: white;
}
</style>
