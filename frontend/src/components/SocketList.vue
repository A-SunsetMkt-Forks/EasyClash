<template>
  <el-table
    :data="list"
    stripe
    style="width: 100%"
  >
    <el-table-column
      v-for="(item, index) in props.headers"
      :key="index + '-' + item.field"
      :label="item.label"
      :width="item.width"
    >
      <template #default="scope">
        <div style="display: flex; align-items: center">
          <span>{{ scope.row[item.field] }}</span>
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>
<script setup lang="ts">
import { ref, defineProps, PropType } from 'vue'
import ReconnectingWebSocket from 'reconnecting-websocket'
type List = Array<Record<string, any>>

const props = defineProps({
  headers: {
    type: Array as PropType<List>,
    default: () => []
  },
  ws: {
    type: String,
    default: ''
  },
  dataType: {
    type: String,
    default: 'push'
  },
  messageHandle: {
    type: Function,
    default: (e: any) => e
  },
  max: {
    type: Number,
    default: 200
  }
})

const options = {
  connectionTimeout: 1000,
  maxRetries: 10
}

const list = ref<List>([])
const ws = new ReconnectingWebSocket(props.ws, [], options)
ws.onopen = () => {}
ws.onmessage = (msg) => {
  const m = props.messageHandle(JSON.parse(msg.data))
  switch (props.dataType) {
    case 'push':
      if (list.value.length >= 200) {
        list.value = list.value.splice(0, 1)
      }
      list.value.push(m)
      break
    case 'overwrite':
      list.value = m
      break
  }
}
ws.onerror = (err) => {
}
</script>
