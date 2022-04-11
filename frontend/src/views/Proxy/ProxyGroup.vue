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
import { ElSwitch, ElMessage } from 'element-plus'
import { isArray } from '@okiss/utils'

const name = 'ProxyGroup'
const listHandle = app.GetProxyGroup
const saveHandle = app.SaveProxyGroup
const form = [
  {
    field: 'name',
    title: '代理组名称',
    rules: 'required',
    readonly: true
  },
  {
    type: 'radio',
    field: 'type',
    title: '类型',
    options: [
      { value: 'url-test', label: 'UrlTest' },
      { value: 'select', label: 'select' }
    ],
    value: 'url-test',
    rules: 'required'
  },
  {
    type: 'input',
    field: 'Url',
    title: 'url',
    value: 'http://www.gstatic.com/generate_204',
    rules: 'required',
    depend: {
      field: 'type',
      value: 'url-test'
    }
  },
  {
    type: 'number',
    field: 'interval',
    title: 'Interval',
    value: 3600,
    rules: 'required',
    depend: {
      field: 'type',
      value: 'url-test'
    }
  },
  {
    type: 'select',
    field: 'proxies',
    title: 'Proxies',
    options: () : Promise<Array<Record<any, any>>> => {
      return new Promise(r => {
        app.GetProxy().then((res: Array<Record<any, any>>) => {
          const list : Array<Record<any, any>> = []
          res.forEach((item: any) => {
            list.push({
              value: item.name,
              label: item.name
            })
          })
          r(list)
        })
      })
    },
    props: {
      multiple: true
    }
  },
  {
    type: 'select',
    field: 'use',
    title: 'use',
    options: () : Promise<Array<Record<any, any>>> => {
      return new Promise(r => {
        app.GetProxyProvider().then((res: Record<any, any>) => {
          const list : Array<Record<any, any>> = []
          Object.keys(res).forEach((item: any) => {
            list.push({
              value: item,
              label: item
            })
          })
          r(list)
        })
      })
    }, props: {
      multiple: true
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
    field: 'proxies'
  },
  {
    field: 'use'
  },
  {
    field: 'enable',
    render: (scope: Record<string, any>, cb: () => void) => {
      return <ElSwitch v-model={scope.row.enable} onChange={cb} disabled={!scope.row.canDisable} />
    }
  }
]

const importHandle = (data: any) => {
  return (isArray(data) ? data : (data['proxy-groups'] || [])).map((item: Record<string, any>) => {
    item.enable = true
    return item
  })
}

const canDel = (data: Record<string, any>) : Promise<boolean> => {
  return new Promise(r => {
    Promise.all([
      new Promise(r => {
        app.GetRule().then((list: Array<Record<string, any>>) => {
          list.forEach(item => {
            if (item.proxyGroup === data.name) {
              ElMessage.error(`this group used by rule: ${item.name}, can not delete it.`)
              r(false)
            }
          })
          r(true)
        })
      }),
      new Promise(r => {
        app.GetRuleSet().then((list: Record<string, Record<string, any>>) => {
          Object.keys(list).forEach(k => {
            const item = list[k]
            if (item?.proxyGroup === data.name) {
              ElMessage.error(`this group used by ruleSet: ${k}, can not delete it.`)
              r(false)
            }
          })
          r(true)
        })
      })
    ]).then(([a, b]) => {
      r(a && b)
    })
  })
}
</script>
