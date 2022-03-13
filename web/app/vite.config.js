import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import svgLoader from "vite-svg-loader";
import { visualizer } from "rollup-plugin-visualizer";

import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    svgLoader({
      svgoConfig: {
        multipass: true,
      },
    }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  css: {
    modules: {
      generateScopedName: "[hash:base64:5]",
    },
    preprocessorOptions: {
      sass: { charset: false },
      scss: { charset: false },
    },
  },
  build: {
    rollupOptions: {
      plugins: [visualizer()],
    },
  },
  server: {
    https: true,
    proxy: {
      "/api": {
        target: "https://allocamelus.localhost",
        changeOrigin: true,
        cookieDomainRewrite: "localhost",
        secure: false,
      },
      "/media": {
        target: "https://allocamelus.localhost",
        changeOrigin: true,
        cookieDomainRewrite: "localhost",
        secure: false,
      },
    },
  },
});
