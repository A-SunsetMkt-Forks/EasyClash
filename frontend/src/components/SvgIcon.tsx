// @ts-nocheck
import { ElIcon } from 'element-plus'
import Refresh from '../assets/images/refresh.svg?component'
import Flash from '../assets/images/flash.svg?component'

import { Component } from 'vue'

const iconMap : Record<string, Component> = {
  refresh: Refresh,
  flash: Flash
}
export const SvgIcon = (props: {name: string}) => {
  const icon = iconMap[props.name]
  return <ElIcon><icon/></ElIcon>
}
