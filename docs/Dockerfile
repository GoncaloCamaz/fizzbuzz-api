FROM node:20-alpine

WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm ci --omit=dev || npm i --omit=dev

COPY server.js ./server.js
COPY openapi.yaml ./openapi.yaml

EXPOSE 3000
CMD ["node", "server.js"]