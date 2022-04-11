<template>
  <el-card shadow="never">
    <template #header>
      Base Config
    </template>
    <span data-wails-no-drag>
      <VForm
        :key="key"
        v-model="baseConf"
        :form-items="formItems"
        :options="{submitButton: false, cancelButton:false}"
        @fieldchange="onSubmit"
      />
    </span>
  </el-card>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import app from '../util/app'
import { ElMessage } from 'element-plus'
import { VForm } from '@okiss/vform'
import { useAppStore } from '@/store/app'
import { debounce } from 'lodash'
const appStore = useAppStore()

const baseConf = ref({})

const formItems = [
  {
    field: 'port',
    label: 'Port',
    rules: 'required'
  },
  {
    field: 'socksPort',
    label: 'SocksPort',
    rules: 'required'
  },
  {
    field: 'redirPort',
    label: 'RedirPort',
    rules: 'required'
  },
  {
    field: 'allowLan',
    type: 'switch',
    label: 'AllowLan'
  },
  {
    field: 'externalController',
    label: 'ExternalCtrl',
    rules: 'required'
  }
]

const key = ref(1)
app.GetBaseConf().then((res: Record<string, any>) => {
  baseConf.value = res
  key.value++
})

const onSubmit = debounce(({field, value}) => {
  app.SaveBaseConf(baseConf.value).then((res: boolean) => {
    if (res) {
      ElMessage.success('保存成功')
      appStore.setChange()
    }
  })
}, 1500)
</script>
