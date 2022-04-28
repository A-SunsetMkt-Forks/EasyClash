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
    proxyDownload: {},
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
      this.proxyDownload = {}
    },
    clashConnInfoWatcher() {
      console.log('clashConnInfoWatcher')
      wsMessage({
        ws: 'ws://' + window.BaseApi + '/connections',
        onmessage: (msg) => {
          const currentIds = [];
          (msg?.connections || []).forEach(item => {
            this.clashConnMsgId[item.id] = item
            currentIds.push(item.id)
          })

          const proxyDownload = {}
          Object.values(this.clashConnMsgId).forEach(item => {
            const proxy = item.chains[item.chains.length - 1]
            if (currentIds.indexOf(item.id) === -1) {
              if (this.proxyDownload[proxy]) {
                this.proxyDownload[proxy] += item.download
              } else {
                this.proxyDownload[proxy] = item.download
              }
              delete this.clashConnMsgId[item.id]
            } else {
              if (proxyDownload[proxy]) {
                proxyDownload[proxy] += item.download
              } else {
                proxyDownload[proxy] = item.download
              }
            }
          })

          Object.keys(this.proxyDownload).forEach(item => {
            const num = this.proxyDownload[item]
            if (proxyDownload[item]) {
              proxyDownload[item] += num
            } else {
              proxyDownload[item] = num
            }
          })

          this.clashConnInfo = {
            uploadTotal: msg.uploadTotal,
            downloadTotal: msg.downloadTotal,
            connNum: Object.keys(this.clashConnMsgId).length,
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
