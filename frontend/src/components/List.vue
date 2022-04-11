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
  },
  list: {
    type: Array as PropType<List>,
    default: () => []
  }
})

</script>
