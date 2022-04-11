<template>
  <el-config-provider>
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </el-config-provider>
</template>

<script setup lang="ts">
import { onUnmounted, watch, computed } from 'vue'
import { useAppStore } from './store/app'
import { ElNotification, ElMessageBox } from 'element-plus'
import { useRouter, useRoute } from 'vue-router'
import app from '@/util/app'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

appStore.clashConnInfoWatcher()

const check = setInterval(() => {
  fetch('http://' + window.BaseApi).then(() => {
    appStore.setClashRunning(true)
  }).catch(() => {
    appStore.setClashRunning(false)
  })
}, 1000)

interface Msg {Code: number, Message: string}
const msg = setInterval(() => {
  app.GetMsg().then((messages: Array<Msg>) => {
    messages.forEach(item => {
      if (item.Code !== 0) {
        ElMessageBox.alert(item.Message, '错误提示', { type: 'error' })
      } else {
        ElMessageBox.alert(item.Message, '提示')
      }
    })
  }).catch((err) => {
    console.log(err)
  })
}, 1000)

onUnmounted(() => {
  clearInterval(check)
  clearInterval(msg)
})

const version = computed(() => {
  return appStore.getVersion()
})

let el = undefined
watch(version, () => {
  if (appStore.getNotify()) {
    return
  }
  el = ElNotification({
    duration: 0,
    title: '配置更新',
    message: '配置已经更新, 但尚未应用, 点我去重启服务',
    type: 'warning',
    onClick: () => {
      if (route.path !== '/dashboard') {
        router.push('/')
      }
      el.close()
      el = undefined
      appStore.setNotify(false)
    }
  })
  appStore.setNotify(true)
})

</script>

<style lang="scss">
html,
body,
#app,
.el-container {
  /*设置内部填充为0，几个布局元素之间没有间距*/
  padding: 0px;
  /*外部间距也是如此设置*/
  margin: 0px;
  /*统一设置高度为100%*/
  height: 100%;
}
</style>
