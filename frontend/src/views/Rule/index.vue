<template>
  <RuleSet ref="ruleSet" />
  <el-divider />
  <Rule ref="rule" />
</template>
<script setup lang="ts">
import Rule from './Rule.vue'
import RuleSet from './RuleSet.vue'

// const importConf = () => {
//   yaml.loadAll(source.value, (data: any) => {
//     const rules = (data.rules || []).map((item: any) => {
//       const token = item.split(',')
//       return {
//         type: token[0],
//         token: token[1],
//         proxyGroup: token[2],
//         enable: true,
//       }
//     })
//     //     rule.value.addEle(rules)
//     dialog.value = false
//   })
// }

</script>
