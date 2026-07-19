# =========================================================================
#  Фронт-контейнер = веб-сервер, раздающий собранную статику React.
#  Стадия 1: Node собирает dist/.
#  Стадия 2: nginx отдаёт эту статику на :80 (замена dev-сервера vite preview).
#  /api сюда НЕ приходит — его проксирует edge-nginx на бэкенд.
# =========================================================================

# --- Стадия 1: сборка ---
FROM node:20-alpine AS build

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm ci

COPY . .
# Относительный /api вшивается в бандл: браузер шлёт на /api, edge-nginx проксирует.
ARG VITE_API_URL=/api
ENV VITE_API_URL=$VITE_API_URL
RUN npm run build

# --- Стадия 2: раздача статики через nginx ---
FROM nginx:1.27-alpine

# Конфиг внутреннего nginx (раздача dist со SPA-fallback).
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Только собранная статика (Node в финальный образ не тянем).
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
# nginx:alpine сам запускает "nginx -g 'daemon off;'" — CMD не нужен.
