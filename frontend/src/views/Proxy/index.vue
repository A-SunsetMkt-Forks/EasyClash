<template>
  <proxy ref="p" />
  <el-divider />
  <proxy-provider ref="pp" />
  <el-divider />
  <proxy-group ref="pg" />
</template>
<script setup lang="ts">
import Proxy from './Proxy.vue'
import ProxyGroup from './ProxyGroup.vue'
import ProxyProvider from './ProxyProvider.vue'
</script>
