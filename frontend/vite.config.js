import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// В dev-режиме (npm run dev) запросы к /api проксируются на бэкенд напрямую.
export default defineConfig({
  plugins: [react()],
  // Только dev: npm run dev поднимает vite на 5173 и проксирует /api на бэкенд.
  // На проде статику раздаёт nginx (см. frontend/Dockerfile), vite здесь не участвует.
  server: {
    host: true,
    port: 5173,
    proxy: {
      "/api": "http://localhost:8080",
    },
  },
});
