# ----------------------
# 前端构建阶段 (Vue)
# ----------------------
FROM node:20-alpine AS frontend-builder
WORKDIR /frontend

# 接收构建参数
ARG GH_TOKEN

# 安装 pnpm
RUN npm install -g pnpm

# 复制依赖文件
COPY frontend/package.json frontend/pnpm-lock.yaml ./

# 配置 GitHub Packages 认证
RUN echo "//npm.pkg.github.com/:_authToken=${GH_TOKEN}" > .npmrc && \
  echo "@yuelioi:registry=https://npm.pkg.github.com" >> .npmrc

# 安装依赖
RUN pnpm install --frozen-lockfile && pnpm store prune

# 清理 .npmrc（安全起见）
RUN rm -f .npmrc

# 复制源码并构建
COPY frontend/ .
RUN pnpm build

# ----------------------
# 后端构建阶段 (Go)
# ----------------------
FROM golang:1.25.1-alpine AS backend-builder
WORKDIR /app

ENV DATABASE_URL=/app/index.db

RUN apk add --no-cache git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=frontend-builder /frontend/dist ./frontend/dist

RUN go build -o server main.go

# ----------------------
# 运行阶段
# ----------------------
FROM alpine:latest

RUN apk add --no-cache sqlite

WORKDIR /app

COPY --from=backend-builder /app/server ./
COPY --from=backend-builder /app/frontend/dist ./frontend/dist

ENV DATABASE_URL=/app/index.db

EXPOSE 8000

CMD ["./server"]