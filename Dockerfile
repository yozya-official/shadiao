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
RUN echo "@yuelioi:registry=https://npm.pkg.github.com" > .npmrc && \
  echo "//npm.pkg.github.com/:_authToken=${GH_TOKEN}" >> .npmrc

# 安装依赖
RUN pnpm install --frozen-lockfile && pnpm store prune

# 清理 .npmrc（有 token）
RUN rm -f .npmrc

# 复制源码并构建
COPY frontend/ .
RUN pnpm build

# ----------------------
# 后端构建阶段 (Go)
# ----------------------
FROM golang:1.25.1-alpine AS backend-builder
WORKDIR /app

# 设置数据库路径
ENV DATABASE_URL=/app/index.db

# 安装依赖工具
RUN apk add --no-cache git build-base

# 复制 go.mod / go.sum 并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制后端源码
COPY . .

# 复制前端构建产物
COPY --from=frontend-builder /frontend/dist ./frontend/dist


# 编译 Go 后端
RUN go build -o server .

# ----------------------
# 运行阶段
# ----------------------
FROM alpine:latest

# 安装 sqlite
RUN apk add --no-cache sqlite

WORKDIR /app

# 复制后端编译产物
COPY --from=backend-builder /app/server ./
COPY --from=backend-builder /app/frontend/dist ./frontend/dist

# 设置数据库环境变量
ENV DATABASE_URL=/app/index.db

# 暴露端口
EXPOSE 9000

# 启动服务
CMD ["./server"]
