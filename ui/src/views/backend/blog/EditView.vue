<template>
  <div>
    <!-- 页头 -->
    <a-page-header
      :title="MODE"
      @back="router.go(-1)"
    >
    </a-page-header>
    <!-- 编辑表单 -->
    <a-form ref="formRef" layout="vertical" :model="form">
    <a-form-item required field="title" label="标题">
      <a-input
        v-model="form.title"
        placeholder="请输入文章标题"
      />
    </a-form-item>
    <a-form-item required field="summary" label="概要">
      <a-input v-model="form.summary" placeholder="请输入文章概要" />
    </a-form-item>
    <a-form-item required field="content" label="内容">
      <MdEditor 
      class="editor" 
      v-model="form.content"
      @onSave="handleSave"
      >
    </MdEditor>
    </a-form-item>
  </a-form>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { GET_BLOG, CRATE_BLOG,UPDATE_BLOG } from '../../../api/blog'
import { state } from '../../../stores/app'
import { Message } from '@arco-design/web-vue';

const router = useRouter()

// 编辑？创建
const MODE = ref('编辑文章')
// 判断页面参数
const id = router.currentRoute.value.query.id
if (!id) {
  MODE.value = '创建文章'
}



// 创建文章
const formRef = ref(null)
const form = ref({
  title: '',
  summary: '',
  author: state.value.token.username,
  content: '',
})
const handleSave = async () => {
  // 校验表单
  const err = await formRef.value.validate()
  if (!err) {
    switch (MODE.value) {
      case '编辑文章':
        // 具体操作, 提交表单
        await UPDATE_BLOG(router.currentRoute.value.query.id, form.value)
        MODE.value = '编辑文章'
        Message.success('保存成功')
        break;
      case '创建文章':
        // 具体操作, 提交表单
        var resp = await CRATE_BLOG(form.value)
        router.replace({query: {id: resp.id}})
        MODE.value = '编辑文章'
        Message.success('创建成功')
        break;
    }
  }
}

onMounted( async () => {
  if (id) {
    const resp = await GET_BLOG(router.currentRoute.value.query.id)
    form.value = resp
  }
})

</script>

<style lang="css" scoped>
.editor {
  height: calc(100vh - 100px);
}
</style>
