<template>
  <div>
    <n-form ref="formRef" class="form" inline :label-width="80" :model="formValue" :rules="rules">
      <div class="form-item">
        <n-form-item label="书源名称" path="user.book_source_name">
          <n-input v-model:value="formValue.bookSourceInfo.bookSourceName" placeholder="书源名称" />
        </n-form-item>
        <n-form-item label="书源网址" path="book_source_url">
          <n-input v-model:value="formValue.bookSourceInfo.bookSourceURL" placeholder="书源网址" />
        </n-form-item>
      </div>
      <div class="form-item">
        <n-form-item label="书源评论" path="book_source_comment">
          <n-input
            v-model:value="formValue.bookSourceInfo.bookSourceComment"
            placeholder="书源评论"
          />
        </n-form-item>
        <n-form-item label="是否启用" path="enabled">
          <n-switch
            checked-value="1"
            unchecked-value="0"
            v-model:value="formValue.bookSourceInfo.enabled"
          />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="搜索地址" path="search_url">
          <n-input v-model:value="formValue.bookSourceInfo.searchURL" placeholder="搜索地址" />
        </n-form-item>
        <n-form-item label="权重" path="weight">
          <n-input-number v-model:value="formValue.bookSourceInfo.weight" clearable />
        </n-form-item>
      </div>

      <n-form-item label="请求头" path="header">
        <n-input
          v-model:value="formValue.bookSourceInfo.header"
          type="textarea"
          placeholder="请求头"
        />
      </n-form-item>

      <div class="form-item">
        <n-form-item label="书名" path="bookSourceRuleBookInfo.name">
          <n-input v-model:value="formValue.bookSourceRuleBookInfo.name" placeholder="书名" />
        </n-form-item>
        <n-form-item label="作者" path="bookSourceRuleBookInfo.author">
          <n-input v-model:value="formValue.bookSourceRuleBookInfo.author" placeholder="作者" />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="简介" path="bookSourceRuleBookInfo.intro">
          <n-input v-model:value="formValue.bookSourceRuleBookInfo.intro" placeholder="简介" />
        </n-form-item>
        <n-form-item label="封面地址" path="bookSourceRuleBookInfo.coverURL">
          <n-input
            v-model:value="formValue.bookSourceRuleBookInfo.coverURL"
            placeholder="封面地址"
          />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="最新章节" path="bookSourceRuleBookInfo.lastChapter">
          <n-input
            v-model:value="formValue.bookSourceRuleBookInfo.lastChapter"
            placeholder="简介"
          />
        </n-form-item>
        <n-form-item label="分类" path="bookSourceRuleBookInfo.kind">
          <n-input v-model:value="formValue.bookSourceRuleBookInfo.kind" placeholder="分类" />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="内容" path="bookSourceRuleContent.content">
          <n-input v-model:value="formValue.bookSourceRuleContent.content" placeholder="内容" />
        </n-form-item>
        <n-form-item label="下一页地址" path="bookSourceRuleContent.nextContentURL">
          <n-input
            v-model:value="formValue.bookSourceRuleContent.nextContentURL"
            placeholder="下一页地址"
          />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="作者" path="bookSourceRuleSearch.author">
          <n-input v-model:value="formValue.bookSourceRuleSearch.author" placeholder="作者" />
        </n-form-item>
        <n-form-item label="书籍列表" path="bookSourceRuleSearch.bookList">
          <n-input v-model:value="formValue.bookSourceRuleSearch.bookList" placeholder="书籍列表" />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="书籍地址" path="bookSourceRuleSearch.bookURL">
          <n-input v-model:value="formValue.bookSourceRuleSearch.bookURL" placeholder="书籍地址" />
        </n-form-item>
        <n-form-item label="检查关键字" path="bookSourceRuleSearch.checkKeyWord">
          <n-input
            v-model:value="formValue.bookSourceRuleSearch.checkKeyWord"
            placeholder="检查关键字"
          />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="封面地址" path="bookSourceRuleSearch.coverURL">
          <n-input v-model:value="formValue.bookSourceRuleSearch.coverURL" placeholder="封面地址" />
        </n-form-item>
        <n-form-item label="简介" path="bookSourceRuleSearch.intro">
          <n-input v-model:value="formValue.bookSourceRuleSearch.intro" placeholder="简介" />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="书名" path="bookSourceRuleSearch.name">
          <n-input v-model:value="formValue.bookSourceRuleSearch.name" placeholder="书名" />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="章节列表" path="bookSourceRuleToc.chapterList">
          <n-input v-model:value="formValue.bookSourceRuleToc.chapterList" placeholder="章节列表" />
        </n-form-item>
        <n-form-item label="章节名" path="bookSourceRuleToc.chapterName">
          <n-input v-model:value="formValue.bookSourceRuleToc.chapterName" placeholder="章节名" />
        </n-form-item>
      </div>

      <div class="form-item">
        <n-form-item label="章节地址" path="bookSourceRuleToc.chapterURL">
          <n-input v-model:value="formValue.bookSourceRuleToc.chapterURL" placeholder="章节地址" />
        </n-form-item>
      </div>
    </n-form>
    <div class="footer">
      <n-button type="primary" @click="handleValidateClick">提交</n-button>
      <n-button @click="emit('closeDialog')">取消</n-button>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
import { type FormInst, useMessage } from 'naive-ui'
import request from '@/composables/request'
const formRef = ref<FormInst | null>(null)
const message = useMessage()
const emit = defineEmits<{
  (event: 'closeDialog'): void
  (event: 'init'): void
}>()
const formValue = ref({
  bookSourceInfo: {
    bookSourceName: '', //书源名称
    bookSourceURL: '', //书源网址
    enabled: '1', //是否启用
    bookSourceComment: '', //书源评论
    header: '', //请求头
    searchURL: '', //搜索地址
    weight: 0 //权重
  },
  bookSourceRuleBookInfo: {
    author: '', //作者
    coverURL: '', //封面地址
    intro: '', //简介
    kind: '', //分类
    lastChapter: '', //最新章节
    name: '' //书名
  },
  bookSourceRuleContent: {
    content: '', //内容
    nextContentURL: '' //下一页地址
  },
  bookSourceRuleSearch: {
    author: '', //作者
    bookList: '', //书籍列表
    bookURL: '', //书籍地址
    checkKeyWord: '', //检查关键字
    coverURL: '', //封面地址
    intro: '', //简介
    name: '' //书名
  },
  bookSourceRuleToc: {
    chapterList: '', //章节列表
    chapterName: '', //章节名
    chapterURL: '' //章节地址
  }
})
const rules = ref({
  user: {
    name: {
      required: true,
      message: '请输入姓名',
      trigger: 'blur'
    },
    age: {
      required: true,
      message: '请输入年龄',
      trigger: ['input', 'blur']
    }
  },
  phone: {
    required: true,
    message: '请输入电话号码',
    trigger: ['input']
  }
})
const handleValidateClick = (e: MouseEvent) => {
  e.preventDefault()
  formRef.value?.validate(async (errors) => {
    if (!errors) {
      const result = await request.post('/admin/book_source', formValue.value)
      if (result.code === 200) {
        message.success('Success')
        emit('closeDialog')
        emit('init')
      }
    }
  })
}
</script>
<style lang="scss">
.footer {
  display: flex;
  gap: 8px;
  justify-content: flex-end;

  .n-button {
    width: 100px;
  }
}
.form {
  display: flex;
  flex-direction: column;
  .form-item {
    width: 100%;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 8px;
  }

  .n-form-item {
    width: 100% !important;
  }
}

@media screen and (max-width: 768px) {
  .form {
    .form-item {
      grid-template-columns: 1fr;
    }
  }
}
</style>
