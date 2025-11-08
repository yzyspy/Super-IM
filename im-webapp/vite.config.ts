import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    host: '0.0.0.0', // 监听所有网络接口
    port: 80, // 指定端口，可选
    proxy: {
      // 代理规则一：转发以 /api 开头的请求
      '/api': {
        target: 'http://localhost:8080', // 你的后端服务器地址
        changeOrigin: true, // 建议开启
        rewrite: (path) => path.replace(/^\/api/, 'api') // 可选，重写路径
      },
    }
  }
})
