<template>
  <Overview />
  <el-divider />
  <el-alert
    v-if="!haveProxies"
    type="warning"
  >请先添加代理服务</el-alert>
  <ProxiesList v-if="clashRunning" />
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import ProxiesList from './ProxiesList.vue'
import Overview from './Overview.vue'
import { useAppStore } from '@/store/app'
import app from '@/util/app'

const appStore = useAppStore()

const clashRunning = computed(() => {
  return appStore.clashRunning
})

const haveProxies = ref(false)

Promise.all([app.GetProxy(), app.GetProxyProvider()]).then(res => {
  haveProxies.value = !!(res[0] || res[1])
})

</script>
