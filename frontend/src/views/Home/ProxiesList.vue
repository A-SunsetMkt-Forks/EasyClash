<template>
  <div>
    <!-- <div class="current-proxy">当前使用: <el-tag>{{ selected }}</el-tag></div> -->
    <div class="box">
      <div class="part">
        <div class="part-name">
          <span style="padding-right: 15px;">{{ proxies.name }}</span>
          <el-button
            type="primary"
            :icon="() => SvgIcon({name: 'flash'})"
            circle
            @click="() => speedTest(proxies.proxies, 1)"
          />
        </div>
        <div class="proxies">
          <div
            v-for="(each, index) in proxies.proxies"
            :key="index"
            :class="['cell' , 'disable-select', bg(each.history[each.history.length - 1]?.delay || 0), each.name === proxies.now ? 'selected' : '']"
            @dblclick="changeProxy('GLOBAL', each.name)"
          >
            <div class="cell-name">{{ each.name }}</div>
            <div class="cell-info">
              <div class="cell-type">{{ each.type }}</div>
              <div class="cell-type">{{ bytesToSize(proxyDownload[each.name] || 0) }}</div>
              <div class="cell-delay">{{ getDelay(each.history || []) || '0 ms' }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <template v-if="proxiders[proxies.now]">
    <div class="box">
      <div class="part">
        <div class="part-name">
          <span style="padding-right: 15px;">Services of <span style="color: #606266">{{ proxiders[proxies.now].name }}</span></span>
          <el-button
            type="primary"
            :icon="() => SvgIcon({name: 'flash'})"
            circle
            @click="() => speedTest(proxiders[proxies.now].use, 2)"
          />
        </div>
        <div class="proxies">
          <div
            v-for="(each, index) in proxiders[proxies.now].proxies"
            :key="index"
            :class="['cell', bg(each.history[each.history.length - 1]?.delay || 0), each.name === proxiders[proxies.now].now ? 'selected' : '']"
          >
            <div class="cell-name">{{ each.name }}</div>
            <div class="cell-info">
              <div class="cell-type">{{ each.type }}</div>
              <div class="cell-delay">{{ getDelay(each.history || []) }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { SvgIcon } from '@/components'
import { sortBy } from 'lodash'
import { get, bytesToSize } from '@/util'
import app from '@/util/app'
import { useAppStore } from '@/store/app'

const appStore = useAppStore()

const proxyDownload = ref({})
watch(() => appStore.getClashConnInfo(), (newVal) => {
  proxyDownload.value = newVal.proxyDownload
})

const proxies = ref<Record<string, any>>([])
const proxiders = ref<Record<string, any>>([])

const _sort = (proxies: Array<Record<string, any>>) => {
  return sortBy(proxies, (each: Record<any, any>) => {
    return each.history[each.history.length - 1]?.delay || 99999999999999
  })
}

const httpApi = `http://${window.BaseApi}`

function getAll() {
  Promise.all([
    get(`${httpApi}/proxies`),
    get(`${httpApi}/providers/proxies`),
    app.GetProxyGroup()
  ]).then(res => {
    proxies.value = []
    const tmp: Array<Record<string, any>> = []
    const _proxies = res[0].proxies
    Object.keys(_proxies).forEach(k => {
      if (['GLOBAL'].indexOf(k) === -1) {
        tmp.push(_proxies[k])
      }
    })
    proxies.value = {
      name: 'Proxies',
      now: res[0].proxies.GLOBAL.now,
      proxies: _sort(tmp)
    }

    const _providers = res[1].providers

    res[2].forEach(item => {
      let tmp = []
      item.use.forEach(k => {
        tmp = tmp.concat(_providers[k].proxies)
      })
      proxiders.value[item.name] = {
        name: item.name,
        now: _proxies[item.name].now,
        proxies: _sort(tmp),
        use: item.use
      }
    })
  })
}

onMounted(() => {
  getAll()
  setTimeout(() => {
    speedAll()
  }, 100)
})

const speedAll = () => {
  speedTest(proxies.value.proxies, 1)
  if (proxiders.value[proxies.value.now]) {
    speedTest(proxiders.value[proxies.value.now].use, 2)
  }
}

const bg = (delay: number) => {
  if (delay === 0) {
    return 'info'
  }
  if (delay < 100) {
    return 'success'
  }
  if (delay < 400) {
    return 'warning'
  }

  return 'error'
}

const getDelay = (history: Array<Record<string, any>>): string => {
  const delay = history[history.length - 1]?.delay
  if (delay) {
    return delay + ' ms'
  }
  return ''
}

const speedTest = (proxy, type) => {
  const list: Array<Record<string, any>> = []
  if (type === 1) {
    (proxy || []).forEach((item: Record<string, any>) => {
      list.push(get(`${httpApi}/proxies/${item.name}/delay?timeout=5000&url=http://www.gstatic.com/generate_204`))
    })
  }
  if (type === 2) {
    (proxy || []).forEach((item: Record<string, any>) => {
      list.push(get(`${httpApi}/providers/proxies/${item}/healthcheck`))
    })
  }
  Promise.all(list).then(() => {
    getAll()
  })
}

const changeProxy = (selector, proxy) => {
  fetch(`${httpApi}/proxies/${selector}`, {
    body: JSON.stringify({ name: proxy }),
    headers: {
      'Content-Type': 'application/json'
    },
    method: 'PUT'
  }).then(res => {
    // if (proxiders.value[proxies.value.now]) {
    //   speedTest(proxiders.value[proxies.value.now].name, 2)
    // }
    getAll()
    // setTimeout(() => {
    // }, 100)
  }).catch(err => {
    console.log(err)
  })
}
</script>

<style scoped>
.current-proxy {
  margin: 10px;
}
.box {
  display: flex;
}
.part {
  display: flex;
  flex-direction: column;
  width: 100%;
}
.part-name {
  margin: 10px;
  height: 30px;
  display: flex;
  align-items: center;
}
.proxies {
  display: flex;
  flex-wrap: wrap;
}
.cell {
  display: flex;
  width: 160px;
  flex-direction: column;
  justify-content: space-between;
  margin: 10px;
  padding: 15px;
  border: 1px solid var(--el-border-color-base);
  border-radius: 5px;
  color: var(--el-color-white);
}
.cell-info {
  display: flex;
  justify-content: space-between;
  font-size: 10px;
}
.cell-name {
  font-size: 16px;
  overflow: hidden;
  /*文本不会换行*/
  white-space: nowrap;
  /*当文本溢出包含元素时，以省略号表示超出的文本*/
  text-overflow: ellipsis;
}
.cell-type {
}
.success {
  background: #67c23a;
}
.warning {
  background: #e6a23c;
}
.error {
  background: #f56c6c;
}
.info {
  background: #909399;
}

.selected {
  position: relative;
}
.selected:before {
  content: "";
  position: absolute;
  right: 0;
  bottom: 0;
  border: 17px solid green;
  /*border: 17px solid #4ABE84;*/
  border-top-color: transparent;
  border-left-color: transparent;
}
.selected:after {
  content: "";
  width: 5px;
  height: 12px;
  position: absolute;
  right: 6px;
  bottom: 6px;
  border: 2px solid #fff;
  border-top-color: transparent;
  border-left-color: transparent;
  transform: rotate(45deg);
}
.disable-select{
  -moz-user-select:none; /*火狐*/
  -webkit-user-select:none; /*webkit浏览器*/
  -ms-user-select:none; /*IE10*/
  -khtml-user-select:none; /*早期浏览器*/
  user-select:none;
}
</style>
