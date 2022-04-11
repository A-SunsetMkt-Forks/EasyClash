import { defineStore } from 'pinia'
import { wsMessage } from '../util'

export const useAppStore = defineStore({
  id: 'app',
  state: () => ({
    notify: false,
    notifyEl: undefined,
    version: 0,
    clashRunning: false,
    clashConnInfo: {
      uploadTotal: 0,
      downloadTotal: 0,
      connNum: 0,
      proxyDownload: {}
    },
    clashConnMsgId: {},
    baseConf: {}
  }),
  actions: {
    setChange() {
      this.version++
    },
    getVersion() {
      return this.version
    },
    setNotify(s: boolean) {
      this.notify = s
    },
    getNotify(): boolean {
      return this.notify
    },
    setClashRunning(s: boolean) {
      this.clashRunning = s
    },
    resetClashConnInfo() {
      this.clashConnInfo = {
        uploadTotal: 0,
        downloadTotal: 0,
        connNum: 0,
        proxyDownload: {}
      }
      this.clashConnMsgId = {}
    },
    clashConnInfoWatcher() {
      console.log('clashConnInfoWatcher')
      wsMessage({
        ws: 'ws://' + window.BaseApi + '/connections',
        onmessage: (msg) => {
          const proxyDownload = {};
          (msg?.connections || []).forEach(item => {
            this.clashConnMsgId[item.id] = item
          })

          Object.values(this.clashConnMsgId).forEach(item => {
            const proxy = item.chains[item.chains.length - 1]
            if (proxyDownload[proxy]) {
              proxyDownload[proxy] += item.download
            } else {
              proxyDownload[proxy] = item.download
            }
          })

          this.clashConnInfo = {
            uploadTotal: msg.uploadTotal,
            downloadTotal: msg.downloadTotal,
            connNum: msg?.connections?.length,
            proxyDownload: proxyDownload
          }
        }
      })
    },
    getClashConnInfo() {
      return this.clashConnInfo
    },
    getClashConns() : Record<string, any> {
      return this.clashConnMsgId
    },
    setBaseConf(conf: Record<string, any>) {
      this.baseConf = conf
    }
  }
})
