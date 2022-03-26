import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import '@/assets/css/global.css'
//要有@才能选择文件

createApp(App).use(store).use(ElementPlus,{size:'small'}).use(router).mount('#app')
