import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'
import visualizer from 'rollup-plugin-visualizer'
import vueJsx from '@vitejs/plugin-vue-jsx'
import copy from 'rollup-plugin-copy'
import svgLoader from 'vite-svg-loader'
const _export = process.env.EXPORT === 'true'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:9090',
        changeOrigin: true,
        rewrite: path => path.replace(RegExp('/api'), '')
      }
    }
  },
  resolve: {
    alias: {
      '@': resolve('src')
    }
  },
  plugins: [
    vue(),
    vueJsx({ transformOn: true }),
    svgLoader(),
    _export
      ? visualizer({
        open: true,
        gzipSize: true,
        brotliSize: true
      })
      : null
  ],
  build: {
    // sourcemap: true,
    rollupOptions: {
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`
      },
      // external: [new RegExp('^monaco-editor')],
      plugins: [
        copy({
          targets: [
            {
              src: 'node_modules/monaco-editor/min/vs/**/*',
              dest: 'dist/assets/monaco-editor/vs'
            }
          ],
          hook: 'writeBundle'
        })
      ]
    }
  }
})
