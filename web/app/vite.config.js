import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import svgLoader from "vite-svg-loader";
import { visualizer } from "rollup-plugin-visualizer";
import tailwindcss from "@tailwindcss/vite";

import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    tailwindcss(),
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
