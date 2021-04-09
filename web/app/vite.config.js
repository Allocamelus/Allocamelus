import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, '/src'),
    },
  },
  server: {
    https: true,
    proxy: {
      '/api': {
        target: 'https://allocamelus.localhost',
        changeOrigin: true,
        cookieDomainRewrite: "localhost",
        secure: true
      },
      '/media': {
        target: 'https://allocamelus.localhost',
        changeOrigin: true,
        cookieDomainRewrite: "localhost",
        secure: true
      }
    }
  }
})
