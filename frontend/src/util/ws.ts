import ReconnectingWebSocket from 'reconnecting-websocket'

interface wsMessageOpt {
    ws: string,
    options?: Record<string, any>
    onmessage: (msg: any) => void,
    onerror?: (err: any) => void,
}

export function wsMessage({ ws, options, onmessage, onerror }: wsMessageOpt) {
  options = Object.assign({
    connectionTimeout: 1000,
    maxRetries: 1000
  }, options)
  console.log(ws)
  const traffic = new ReconnectingWebSocket(ws, [], options)
  traffic.onmessage = msg => onmessage(JSON.parse(msg.data))
  traffic.onerror = err => onerror ? onerror(err) : ''
}
