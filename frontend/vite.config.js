import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// В dev-режиме (npm run dev) запросы к /api проксируются на бэкенд напрямую.
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 5173,
    proxy: {
      "/api": "http://localhost:8080",
    },
  },
  preview: {
    host: true,
    port: 4173,
  },
});
