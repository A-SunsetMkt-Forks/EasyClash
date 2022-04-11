import { createApp } from 'vue'
import App from './App.vue'
import { store, key } from './store'
import router from './router'
import '@/styles/index.scss'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import '@okiss/vform/style.css'
import { createPinia } from 'pinia'
import app from '@/util/app'

app.GetBaseConf().then(base => {
  window.BaseApi = base.externalController
  window.HomeDir = base.homeDir
  createApp(App)
    .use(createPinia())
    .use(ElementPlus)
    .use(store, key)
    .use(router)
    .mount('#app')
})

