<template>
  <n-layout>
    <n-layout-header>
      <div class="main-header">梦境小说管理系统</div>
    </n-layout-header>
    <n-layout has-sider>
      <n-layout-sider
        bordered
        collapse-mode="width"
        :collapsed-width="64"
        :width="240"
        :collapsed="collapsed"
        show-trigger
        @collapse="collapsed = true"
        @expand="collapsed = false"
      >
        <n-menu
          :value="selectedMenu"
          :collapsed="collapsed"
          :collapsed-width="64"
          :collapsed-icon-size="22"
          :options="menuOptions"
        />
      </n-layout-sider>
      <n-layout-content content-style="padding: 20px"> <RouterView /> </n-layout-content>
    </n-layout>
  </n-layout>
</template>
<script lang="ts" setup>
import { computed, h, ref } from 'vue'
import { RouterLink, RouterView, useRoute } from 'vue-router'
import { NIcon } from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import VueIcon from '@/components/VueIcon.vue'

const collapsed = ref(false)

const route = useRoute()
const selectedMenu = computed(() => route.name as string)
const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'crawl'
          }
        },
        { default: () => '爬虫管理' }
      ),
    key: 'crawl',
    icon: renderIcon('search')
  }
]

function renderIcon(iconName: string) {
  return () => h(NIcon, null, { default: () => h(VueIcon, { name: iconName }, '') })
}
</script>

<style scoped lang="scss">
.admin-title {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  font-weight: 900;
}
.main-header {
  height: 60px;
  padding: 20px;
  border-bottom: solid 1px #f0f0f0;
}
</style>
