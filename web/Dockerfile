FROM node:22-alpine AS base

WORKDIR /web

COPY .env .

COPY web/package*.json web/pnpm-lock.yaml .
RUN npm install -g pnpm
RUN pnpm i

COPY web/. .

RUN pnpm build

ENV HOST=0.0.0.0
ENV PORT=${WEB_PORT}
EXPOSE ${WEB_PORT}
CMD node ./dist/server/entry.mjs
