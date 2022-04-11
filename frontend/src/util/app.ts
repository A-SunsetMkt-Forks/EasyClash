// import { isDevMode } from './env'
import type { go } from '@/wailsjs/go'

const Go : go = window.go
const _app = () => Go?.app?.App

// const mock = {
//   app: {
//     App: {}
//   }
// }

const app = _app()

export default app
