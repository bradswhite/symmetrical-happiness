FROM node:22-alpine AS base

WORKDIR /web

COPY package*.json .
RUN npm install -g pnpm
RUN pnpm i

COPY . .
RUN pnpm build

ENV HOST=0.0.0.0
ENV PORT=${WEB_PORT}
EXPOSE ${WEB_PORT}
CMD node ./dist/server/entry.mjs
