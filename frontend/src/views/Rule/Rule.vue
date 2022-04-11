<template>
  <part-mng
    :name="name"
    :list-handle="listHandle"
    :save-handle="saveHandle"
    :form="form"
    :header="header"
    :import-handle="importHandle"
    insert-type="unshift"
  />
</template>
<script setup lang="tsx">
import PartMng from '../../components/PartMng.vue'
import app from '../../util/app'
import { ElSwitch } from 'element-plus'
import { isArray } from '@okiss/utils'

const name = 'Rule'
const listHandle = app.GetRule
const saveHandle = app.SaveRule
const form = [
  {
    type: 'radio',
    field: 'type',
    title: '规则类型',
    options: [
      { value: 'DOMAIN-SUFFIX', label: 'DOMAIN-SUFFIX' },
      { value: 'DOMAIN-KEYWORD', label: 'DOMAIN-KEYWORD' },
      { value: 'DOMAIN', label: 'DOMAIN' },
      { value: 'IP-CIDR', label: 'IP-CIDR' },
      { value: 'IP-CIDR6', label: 'IP-CIDR6' },
      { value: 'GEOIP', label: 'GEOIP' },
      { value: 'MATCH', label: 'MATCH' }
    ],
    rules: 'required'
  },
  {
    field: 'token',
    title: '关键字',
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
    field: 'token'
  },
  {
    field: 'proxyGroup'
  },
  {
    field: 'enable',
    render: (scope: Record<string, any>, cb: () => void) => {
      return <ElSwitch v-model={scope.row.enable} onChange={cb} disabled={!scope.row.canDisable} />
    }
  }
]

const importHandle = (data: any) => {
  return (isArray(data) ? data : (data.rules || [])).map((item: any) => {
    const token = item.split(',')
    return {
      type: token[0],
      token: token[1],
      proxyGroup: token[2],
      enable: true
    }
  })
}
</script>
