import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': { // 匹配所有以 /api 开头的请求
        target: 'http://localhost:8080', // 目标后端服务器
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '') // 重写路径，去掉 /api
      },
    },
  },
})
