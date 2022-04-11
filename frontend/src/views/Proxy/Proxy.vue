<template>
  <part-mng
    :name="name"
    :list-handle="listHandle"
    :save-handle="saveHandle"
    :form="form"
    :header="header"
    :import-handle="importHandle"
    :can-del="canDel"
  />
</template>
<script setup lang="tsx">
import PartMng from '../../components/PartMng.vue'
import app from '../../util/app'
import { ElSwitch } from 'element-plus'
import { isArray } from '@okiss/utils'
import { ElMessage } from 'element-plus/es'

const name = 'Proxy'
const listHandle = app.GetProxy
const saveHandle = app.SaveProxy
const form = [
  {
    field: 'name',
    title: '节点名称',
    rules: 'required',
    readonly: true,
    props: {
      autocomplete: 'off'
    }
  },
  {
    field: 'server',
    title: '节点地址',
    rules: 'required'
  },
  {
    type: 'number',
    field: 'port',
    title: '端口',
    rules: 'required'
  },
  {
    type: 'radio',
    field: 'type',
    title: '节点类型',
    options: [
      { value: 'vmess', label: 'vmess' },
      { value: 'trojan', label: 'trojan' }
    ],
    value: 'vmess'
  },
  {
    field: 'uuid',
    title: 'uuid',
    rules: 'required',
    depend: {
      field: 'type',
      value: 'vmess'
    }
  },
  {
    type: 'number',
    field: 'alterId',
    title: 'alterId',
    props: {
      controls: false
    },
    rules: 'required',
    depend: {
      field: 'type',
      value: 'vmess'
    }
  },
  {
    field: 'cipher',
    title: 'cipher',
    value: 'auto',
    rules: 'required',
    depend: {
      field: 'type',
      value: 'vmess'
    }
  },
  {
    field: 'password',
    title: 'password',
    rules: 'required',
    depend: {
      field: 'type',
      value: 'trojan'
    }
  },
  {
    field: 'sni',
    title: 'sni',
    rules: 'required',
    depend: {
      field: 'type',
      value: 'trojan'
    }
  },
  {
    type: 'switch',
    field: 'udp',
    title: 'udp',
    value: true,
    rules: 'required',
    depend: {
      field: 'type',
      value: 'trojan'
    }
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
    field: 'server'
  },
  {
    field: 'enable',
    render: (scope: Record<string, any>, cb: () => void) => {
      return <ElSwitch v-model={scope.row.enable} onChange={cb} />
    }
  }
]

const importHandle = (data: any) => {
  return (isArray(data) ? data : (data.proxies || [])).map((item: Record<string, any>) => {
    item.enable = true
    return item
  })
}

const canDel = (data: Record<string, any>) : Promise<boolean> => {
  return new Promise(r => {
    app.GetProxyGroup().then((list: Array<Record<string, any>>) => {
      (list || []).forEach(item => {
        if (item?.proxies && item?.proxies.indexOf(data.name) > -1) {
          ElMessage.error(`this proxy used by ProxyGroup: ${item.name}, can not delete it.`)
          r(false)
        }
      })
      r(true)
    })
  })
}
</script>
