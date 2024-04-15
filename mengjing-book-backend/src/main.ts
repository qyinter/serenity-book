import './assets/styles/global.scss'
import 'virtual:svg-icons-register'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import VueIcon from '@/components/VueIcon.vue'

const app = createApp(App)

app.component('vue-icon', VueIcon)
app.use(createPinia())
app.use(router)

app.mount('#app')
