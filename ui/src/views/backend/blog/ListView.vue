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
        <a-button type="primary" @click="router.push({name: 'BackendEditBlog'})">创建文章</a-button>
      </div>
      <div>
        <a-input 
        v-model="request.keywords" 
        :style="{ width: '320px' }" 
        placeholder="请输入文字名称敲回车键搜索"
        allow-clear
        @press-enter="ListBlog"
        />
      </div>
    </div>
    <div>
      <a-table :loading="isLoading" column-resizable :bordered="{cell:true}" :pagination="false" :data="data.items" >
        <template #columns>
          <a-table-column title="编号" data-index="id" align="center"></a-table-column>
          <a-table-column title="名称" data-index="title" align="center"></a-table-column>
          <a-table-column title="作者" data-index="author" align="center"></a-table-column>
          <a-table-column title="分类" align="center">
            <template #cell="{ record }">
              <a-tag key="目录" color="pinkpurple" bordered>{{ record.tags['目录'] }}</a-tag>
            </template>
          </a-table-column>
          <a-table-column title="操作" align="center">
            <template #cell="{ record }">
              <a-space>
                <a-button @click="router.push({name: 'BackendDetailBlog', params: {id: record.id}})">预览</a-button>
                <a-button @click="router.push({name: 'BackendEditBlog', query: {id: record.id}})">编辑</a-button>
                <a-button @click="$modal.info({ title:'Name', content:record.title })">删除</a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
    </div>
    <div class="pagi">
      <a-pagination 
      :total="data.total" 
      :page-size-options="[2, 10, 20, 50, 100, 200]"
      @page-size-change="handlePageSizeChange"
      @change="hanlePageNumberChange"
      show-total 
      show-jumper
      show-page-size
       />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { LIST_BLOG } from '../../../api/blog'
import { useRouter } from 'vue-router';

const router = useRouter()

const isLoading = ref(false)
const data = ref({items: [], total: 0});
const ListBlog = async () => {
  isLoading.value = true
  try {
    const resp = await LIST_BLOG(request.value)
    data.value = resp
    console.log(resp);
  } finally {
    isLoading.value = false
  }
}
// 选择页面渲染完成后加载数据, 数据处于加载中的时候 给予一个Loadding反馈
onMounted(() => {
  ListBlog()
})

// 声明一个响应式变量保持当前用户输入
const request = ref({
  page_size: 10,
  page_number: 1,
  create_by: '',
  keywords: '',
})
// 处理页面大小变化
const handlePageSizeChange = (pageSize) => {
  request.value.page_size = pageSize
  ListBlog()
}
const hanlePageNumberChange = (pageNumber) => {
  request.value.page_number = pageNumber
  ListBlog()
}
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
