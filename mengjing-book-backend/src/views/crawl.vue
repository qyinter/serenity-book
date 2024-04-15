<template>
  <div>
    <div class="query-group">
      <n-input v-model:value="query.bookSourceName" type="text" placeholder="书源名称" />
      <n-input v-model:value="query.bookSourceURL" type="text" placeholder="书源网址" />
      <NButton type="info" class="icon-btn" @click="init">
        <template #icon>
          <vue-icon name="search" size="0.85rem" />
        </template>
      </NButton>
      <NButton class="icon-btn" @click="reset">
        <template #icon>
          <vue-icon name="reset" size="1.1rem" />
        </template>
      </NButton>
      <NButton type="info" @click="addBookSource"> 添加书源 </NButton>
    </div>
    <n-data-table :columns="columns" :data="data" :pagination="false" :bordered="false" />
  </div>
</template>
<script lang="ts" setup>
import { h, onMounted, ref } from 'vue'
import { NButton, useDialog, useMessage } from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import type { BookSourceInfo, BookSourceQuery, BookSourceInfoList } from '@/types/crawl'
import CrawlForm from '@/components/CrawlForm.vue'
import { useRoute } from 'vue-router'
import request from '@/composables/request'
const dialog = useDialog()

const query = ref<BookSourceQuery>({
  bookSourceName: '',
  bookSourceURL: '',
  pageSize: 10,
  pageNum: 1
})

const message = useMessage()
const data = ref<BookSourceInfo[]>([])
const total = ref(0)
onMounted(() => {
  init()
})

const init = async () => {
  const result = await request.post<BookSourceInfoList>('/admin/book_source/list', query.value)
  if (result.code === 200) {
    data.value = result.data.list
    total.value = result.data.total
  }
}

const createColumns = ({
  edit,
  del,
  check
}: {
  edit: (row: BookSourceInfo) => void
  del: (row: BookSourceInfo) => void
  check: (row: BookSourceInfo) => void
}): DataTableColumns<BookSourceInfo> => {
  return [
    {
      title: '书源名称',
      key: 'bookSourceName'
    },
    {
      title: '书源网址',
      key: 'bookSourceURL',
      resizable: true
    },
    {
      title: '已启用',
      key: 'enabled',
      render(row) {
        return row.enabled ? '是' : '否'
      }
    },
    {
      title: '最后更新时间',
      key: 'lastUpdateTime'
    },
    {
      title: '响应时间',
      key: 'respondTime',
      minWidth: 100,
      maxWidth: 500
    },
    {
      title: '操作',
      key: 'actions',
      render(row) {
        return h('div', { style: { display: 'flex', gap: '10px' } }, [
          h(
            NButton,
            {
              size: 'small',
              onClick: () => check(row)
            },
            { default: () => '有效性测试' }
          ),
          h(
            NButton,
            {
              size: 'small',
              onClick: () => edit(row)
            },
            { default: () => '修改' }
          ),
          h(
            NButton,
            {
              size: 'small',
              onClick: () => del(row)
            },
            { default: () => '删除' }
          )
        ])
      }
    }
  ]
}

const columns = createColumns({
  edit(row: BookSourceInfo) {
    console.log(row)
  },
  async del(row: BookSourceInfo) {
    const result = await request.del(`/admin/book_source/${row.bookSourceID}`)
    if (result.code === 200) {
      message.success(result.msg)
      init()
    }
  },
  check(row: BookSourceInfo) {
    console.log(row)
  }
})

const reset = () => {
  query.value = {
    bookSourceName: '',
    bookSourceURL: '',
    pageSize: 10,
    pageNum: 1
  }
  init()
}
const route = useRoute()
const addBookSource = () => {
  dialog.create({
    title: '',
    showIcon: false,
    closeOnEsc: false,
    maskClosable: false,
    closable: false,
    class: 'proposal-dialog',
    content: () =>
      h(
        CrawlForm,
        { id: route.params.id, onCloseDialog: () => dialog.destroyAll(), onInit: () => init() },
        ''
      )
  })
}
</script>
<style lang="scss" scoped></style>
