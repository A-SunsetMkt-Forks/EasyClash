import { isFunction } from 'lodash'
import { Component } from 'vue'
export * from './SvgIcon'

export const Render = (props: { content : Component | Function}) => {
  if (isFunction(props.content)) {
    return (props.content as Function)()
  }
  const tmp = props.content
  return <tmp/>
}

