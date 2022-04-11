<template>
  <el-row :gutter="20">
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">上行</el-col>
          <el-col>{{ bytesToSize(isRunning ? current.up : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">下行</el-col>
          <el-col>{{ bytesToSize(isRunning ? current.down : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">总下载</el-col>
          <el-col>{{ bytesToSize(isRunning ? conn.downloadTotal : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">总上传</el-col>
          <el-col>{{ bytesToSize(isRunning ? conn.uploadTotal : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">连接数</el-col>
          <el-col>{{ isRunning ? (conn.connNum || 0) : 0 }}</el-col>
        </el-row>
      </el-card>
    </el-col>
  </el-row>

  <el-divider />
  <el-button-group class="ml-4">
    <el-button
      :icon="isRunning ? VideoPause : VideoPlay"
      :type="isRunning ? 'success' : 'info'"
      @click="stopOrStartClash"
    >{{ isRunning ? '停止' : '启动' }} Clash Core</el-button>
<!--    <el-button-->
<!--      v-if="isRunning"-->
<!--      type="success"-->
<!--      :icon="() => SvgIcon({name: 'refresh'})"-->
<!--      @click="reloadClashConf"-->
<!--    />-->
  </el-button-group>

  <el-button
    :icon="Promotion"
    :type="isSystemProxy ? 'primary' : 'info'"
    @click="setOrUnSetSystemProxy"
  >{{ isSystemProxy ? '取消' : '设置为' }}系统代理</el-button>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue'
import { VideoPause, VideoPlay, Promotion } from '@element-plus/icons-vue'
import app from '@/util/app'
import { bytesToSize, wsMessage, get } from '@/util'
import { ElMessage } from 'element-plus'
import { useAppStore } from '@/store/app'
import { SvgIcon } from '@/components'

const appStore = useAppStore()

const current = ref<Record<string, number>>({})
wsMessage({
  ws: `ws://${window.BaseApi}/traffic`,
  onmessage: (msg) => current.value = msg
})

const conn = computed(() => {
  return appStore.getClashConnInfo()
})

const isRunning = computed(() => {
  return appStore.clashRunning
})
const stopOrStartClash = async() => {
  if (isRunning.value) {
    await app.StopClash()
    appStore.resetClashConnInfo()
    isSystemProxy.value = false
  } else {
    await app.StartClash()
    isSystemProxy.value = true
    const base = await app.GetBaseConf()
    if (window.BaseApi !== base.externalController) {
      window.BaseApi = base.externalController
      location.reload()
    }
  }
}

const isSystemProxy = ref(true)
const setOrUnSetSystemProxy = () => {
  if (isSystemProxy.value) {
    app.UnSetSystemProxy().then(() => {
      isSystemProxy.value = false
    })
  } else {
    app.SetSystemProxy().then(() => {
      isSystemProxy.value = true
    })
  }
}

const reloadClashConf = () => {
  get(`http://${window.BaseApi}/configs`, {
    method: 'PUT',
    body: JSON.stringify({ path: window.HomeDir + '/config.yml' })
  }).then(res => {
    ElMessage.success('clash配置应用成功, 但不会更改端口等基础配置')
  })
}
</script>

<style lang="scss">
.label {
  font-size: 10px;
}
</style>
