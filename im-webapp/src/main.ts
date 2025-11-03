import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import './assets/main.css'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'


import './assets/iconfont/iconfont.js'
import SvgIcon from './components/SvgIcon.vue'



const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus)

app.component('SvgIcon', SvgIcon);

app.mount('#app')
