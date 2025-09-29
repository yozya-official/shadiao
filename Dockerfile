# ----------------------
# 前端构建阶段 (Vue)
# ----------------------
FROM node:20-alpine AS frontend-builder
WORKDIR /frontend

# 安装 pnpm
RUN npm install -g pnpm

# 复制依赖文件
COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile && pnpm store prune

# 复制源码并构建
COPY frontend/ .
RUN pnpm build

# ----------------------
# 后端构建阶段 (Go)
# ----------------------
FROM golang:1.22-alpine AS backend-builder
WORKDIR /app

# 设置默认数据库路径
ENV DATABASE_URL=/app/index.db

# 安装必要工具
RUN apk add --no-cache git build-base

# 复制 Go 模块依赖并下载
COPY go.mod go.sum ./
RUN go mod download

# 复制后端源码
COPY . .

# 拷贝前端构建好的 dist
COPY --from=frontend-builder /frontend/dist ./frontend/dist

# 构建 Go 可执行文件
RUN go build -o server main.go

# ----------------------
# 运行阶段
# ----------------------
FROM alpine:latest

# 安装 SQLite 运行时
RUN apk add --no-cache sqlite

WORKDIR /app

# 复制后端可执行文件
COPY --from=backend-builder /app/server ./

# 复制前端静态资源
COPY --from=backend-builder /app/frontend/dist ./frontend/dist

# 设置数据库路径环境变量
ENV DATABASE_URL=/app/index.db

# 暴露端口
EXPOSE 8000

# 默认启动命令
CMD ["./server"]
