import { createRouter, createWebHashHistory } from 'vue-router'
import Layout from '@/layout/Layout.vue'
import Dashboard from './views/Home/index.vue'
import BaseConf from './views/BaseConf.vue'
import Proxy from './views/Proxy/index.vue'
import Rule from './views/Rule/index.vue'
import Log from './views/Log.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard
      },
      {
        path: '/base_setting',
        name: 'BaseConf',
        component: BaseConf
      },
      {
        path: '/proxy',
        name: 'Proxy',
        component: Proxy
      },
      {
        path: '/rule',
        name: 'Rule',
        component: Rule
      },
      {
        path: '/log',
        name: 'Log',
        component: Log
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
