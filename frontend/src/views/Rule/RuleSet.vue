<template>
  <part-mng
    :name="name"
    :list-handle="listHandle"
    :save-handle="saveHandle"
    :form="form"
    :header="header"
    :import-btn="false"
    insert-type="unshift"
  />
</template>
<script setup lang="tsx">
import PartMng from '../../components/PartMng.vue'
import app from '../../util/app'
import { ElSwitch } from 'element-plus'

const name = 'RuleProvider'
const listHandle = () : Promise<Array<Record<string, any>>> => {
  return new Promise(r => {
    app.GetRuleSet().then((res: Record<string, any>) => {
      const tmp : Array<Record<string, any>> = []
      Object.keys(res).forEach(item => {
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
  return new Promise(r => r(app.SaveRuleSet(tmp)))
}

const form = [
  {
    type: 'radio',
    field: 'type',
    title: '规则类型',
    options: [
      { value: 'file', label: 'File' },
      { value: 'url', label: 'Url' }
    ],
    rules: 'required'
  },
  {
    field: 'url',
    title: 'url',
    rules: 'required',
    depend: {
      field: 'type',
      value: 'url'
    }
  },
  {
    type: 'number',
    field: 'interval',
    title: 'interval',
    rules: 'required',
    depend: {
      field: 'type',
      value: 'url'
    }
  },
  {
    type: 'radio',
    field: 'behavior',
    title: '解析类型',
    options: [
      { value: 'ipcidr', label: 'ipcidr' },
      { value: 'domain', label: 'domain' }
    ],
    rules: 'required'
  },
  {
    field: 'proxyGroup',
    title: '代理组',
    type: 'select',
    options: () : Promise<Array<Record<any, any>>> => {
      return new Promise(r => {
        app.GetProxyGroup().then((res: Array<Record<any, any>>) => {
          const list : Array<Record<any, any>> = [
            { value: 'DIRECT', label: 'DIRECT' }
          ]
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
    rules: 'required'
  },
  {
    field: 'path',
    title: '存储位置',
    rules: 'required'
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
    field: 'type'
  },
  {
    field: 'name'
  },
  {
    field: 'behavior'
  },
  {
    field: 'enable',
    render: (scope: Record<string, any>, cb: () => void) => {
      return <ElSwitch v-model={scope.row.enable} onChange={cb} />
    }
  }
]
</script>
