<template>
  <part-mng
    :name="name"
    :list-handle="listHandle"
    :save-handle="saveHandle"
    :form="form"
    :header="header"
    :import-btn="false"
    :can-del="canDel"
  >
    <template #button>
      <el-button
        class="button"
        type="text"
        @click="onclick"
      >订阅链接转换</el-button>
    </template>
  </part-mng>
</template>
<script setup lang="tsx">
import PartMng from '../../components/PartMng.vue'
import app from '../../util/app'
import { ElSwitch, ElMessage } from 'element-plus'

const name = 'ProxyProvider'
const listHandle = () : Promise<Array<Record<string, any>>> => {
  return new Promise(r => {
    app.GetProxyProvider().then((res: Record<string, any>) => {
      const tmp : Array<Record<string, any>> = []
      Object.keys(res || {}).forEach(item => {
        tmp.push({
          name: item,
          ...res[item]
        })
      })
      r(tmp)
    })
  })
}
const saveHandle = (list: Array<Record<string, any>>) : Promise<boolean> => {
  const tmp : Record<string, any> = {}
  list.forEach(item => {
    tmp[item.name] = item
  })
  return new Promise(r => r(app.SaveProxyProvider(tmp)))
}
const form = [
  {
    field: 'name',
    title: 'name',
    readonly: true,
    rules: {
      required: true,
      pattern: /^[A-Za-z0-9_]+$/,
      message: '只能包含字母数字下划线',
      trigger: 'blur'
    }
  },
  {
    type: 'hidden',
    field: 'type',
    title: '类型',
    value: 'http'
  },
  {
    type: 'input',
    field: 'url',
    title: 'url',
    rules: 'required'
  },
  {
    type: 'number',
    field: 'interval',
    title: 'Interval',
    props: {
      control: false
    },
    value: 3600,
    rules: 'required'
  },
  {
    type: 'sub-form',
    field: 'healthCheck',
    label: '监控检查',
    props: {
      formItems: [
        {
          type: 'hidden',
          field: 'enable',
          value: true
        },
        {
          type: 'hidden',
          field: 'lazy',
          value: true
        },
        {
          field: 'url',
          title: 'url',
          rules: 'required',
          value: 'http://www.gstatic.com/generate_204'
        },
        {
          type: 'number',
          field: 'interval',
          title: 'Interval',
          props: {
            control: false
          },
          value: 3600,
          rules: 'required'
        }
      ]
    }
  },
  {
    type: 'hidden',
    field: 'path',
    title: 'path',
    validate: [{ required: true, message: '请输入path', trigger: 'blur' }]
  },
  {
    type: 'hidden',
    field: 'enable',
    value: true
  },
  {
    type: 'hidden',
    field: 'canDel',
    value: true
  },
  {
    type: 'hidden',
    field: 'canEdit',
    value: true
  },
  {
    type: 'hidden',
    field: 'canDisable',
    value: true
  }
]
const header = [
  {
    field: 'name'
  },
  {
    field: 'type'
  },
  {
    field: 'url'
  },
  {
    field: 'enable',
    render: (scope: Record<string, any>, cb: () => void) => {
      const check = (val) => {
        !val && handler(scope.row).then(res => {
          if (res.value) {
            cb()
          } else {
            ElMessage.error(`this provider used by ProxyGroup: ${res.name}, can not disable it.`)
            scope.row.enable = true
          }
        })
        if (val) {
          cb()
        }
      }
      return <ElSwitch v-model={scope.row.enable} onChange={check} />
    }
  }
]

const handler = (data: Record<string, any>) : Promise<boolean> => {
  return new Promise(r => {
    app.GetProxyGroup().then((list: Array<Record<string, any>>) => {
      (list || []).forEach(item => {
        if (item?.use && item?.use.indexOf(data.name) > -1) {
          r({ name: data.name, value: false })
        }
      })
      r({ name: data.name, value: true })
    })
  })
}

const canDel = (data: Record<string, any>) : Promise<boolean> => {
  return new Promise(r => {
    handler(data).then(res => {
      if (!res.value) {
        ElMessage.error(`this provider used by ProxyGroup: ${data.name}, can not delete it.`)
      }
      r(res.value)
    })
  })
}

const onclick = () => {
  window.runtime.BrowserOpenURL('https://bianyuan.xyz/')
}

</script>
