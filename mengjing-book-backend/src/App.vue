<template>
  <n-config-provider :theme-overrides="themeOverrides">
    <n-message-provider>
      <n-dialog-provider>
        <MainLayout />
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>
<script lang="ts" setup>
import { type GlobalThemeOverrides } from 'naive-ui'
import MainLayout from '@/layout/MainLayout.vue'

const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#03acff',
    primaryColorHover: '#03acff',
    primaryColorPressed: '#03acff',
    borderRadius: '4px'
  },
  Input: {
    textColor: '#000135',
    heightLarge: '48px',
    placeholderColor: '#909295',

    borderFocusError: '1px solid #f56c6c',
    borderHoverError: '1px solid #f56c6c',
    borderActiveError: '1px solid #f56c6c',
    borderError: '1px solid #f56c6c'
  },

  Form: {
    feedbackHeightMedium: '20px',
    feedbackFontSizeMedium: '12px',
    feedbackPadding: '1px 0 0 0',

    feedbackTextColorError: '#f56c6c',
    asteriskColor: '#f56c6c'
  },
  Button: {
    colorInfo: '#03ACFF',
    colorHoverInfo: 'rgba(3, 172, 255, 0.8)',
    colorPressedInfo: '#03ACFF',
    colorFocusInfo: 'rgba(3, 172, 255, 0.8)',
    colorDisabledInfo: '#EDF6FB',

    borderInfo: '#03ACFF',
    borderHoverInfo: 'rgba(3, 172, 255, 0.8)',
    borderPressedInfo: '#03ACFF',
    borderFocusInfo: 'rgba(3, 172, 255, 0.8)',
    borderDisabledInfo: '#EDF6FB',

    borderRadiusMedium: '8px',
    fontSizeMedium: '14px'
  },

  Checkbox: {
    borderRadius: '6px'
  },

  Message: {
    borderRadius: '12px',
    iconMargin: '0 4px 0 0',

    textColorError: '#f56c6c',
    colorError: '#fef0f0',
    iconColorError: '#f56c6c',

    textColorSuccess: '#67c23a',
    colorSuccess: '#e1f3d8',
    iconColorSuccess: '#67c23a'
  },

  Collapse: {
    dividerColor: '#E4EAF0',
    titleTextColor: '#000135',
    titleFontSize: '20px',
    titleFontWeight: '500'
  },

  Tabs: {
    tabBorderRadius: '32px'
  },

  Dropdown: {
    borderRadius: '8px',
    padding: '6px'
  },

  Pagination: {
    itemColorActive: '#03acff',
    itemTextColorActive: '#fff',
    itemColorActiveHover: '#03acff',
    itemColorHover: '#03acff',
    itemTextColorActiveHover: '#fff',
    itemTextColorHover: '#fff'
  }
}
</script>
