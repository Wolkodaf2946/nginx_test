FROM node:20-alpine

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm install

COPY . .

# URL бэкенда вшивается в бандл на этапе сборки (Vite читает VITE_*).
ARG VITE_API_URL=/api
ENV VITE_API_URL=$VITE_API_URL
RUN npm run build

EXPOSE 4173

# vite preview раздаёт собранную статику из dist/.
CMD ["npm", "run", "preview", "--", "--host", "0.0.0.0", "--port", "4173"]
