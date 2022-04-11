<template>
  <el-card shadow="never">
    <template #header>Current Connections</template>
    <List
      :headers="connsHeaders"
      :list="conns"
    />
  </el-card>
  <el-divider />
  <el-card shadow="never">
    <template #header>Proxy Log</template>
    <socket-list
      :headers="logHeaders"
      :ws="'ws://' + window.BaseApi + '/logs'"
    />
  </el-card>
</template>
<script setup lang="ts">
import SocketList from '../components/SocketList.vue'
import List from '../components/List.vue'
import { bytesToSize } from '../util'
import { useAppStore } from '@/store/app'
import { computed } from 'vue'

const store = useAppStore()

const logHeaders = [
  {
    label: 'Type',
    field: 'type'
  },
  {
    label: 'Payload',
    field: 'payload',
    width: 800
  }
]

const connsHeaders = [
  {
    label: '类型',
    field: 'type'
  },
  {
    label: 'HOST',
    field: 'host'
  },
  {
    label: '上传',
    field: 'upload'
  },
  {
    label: '下载',
    field: 'download'
  },
  {
    label: 'Chains',
    field: 'chains',
    width: 350
  }
]

const conns = computed(() => {
  const list = Object.values(store.getClashConns())

  return list.map((item: Record<string, any>) => {
    return {
      type: item.metadata.type,
      host: item.metadata.host,
      upload: bytesToSize(item.upload),
      download: bytesToSize(item.download),
      chains: item.chains
    }
  })
})

</script>

<style lang="scss">
.label {
  font-size: 10px;
}
</style>
