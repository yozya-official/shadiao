# -----------------------
# 1. 构建前端 (Vue)
# -----------------------
FROM node:20 AS frontend-build

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install --frozen-lockfile
COPY frontend/ .
RUN npm run build

# -----------------------
# 2. 构建后端 (Go)
# -----------------------
FROM golang:1.22 AS backend-build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# 拷贝前端构建好的 dist
COPY --from=frontend-build /app/frontend/dist ./frontend/dist

RUN go build -o server main.go

# -----------------------
# 3. 最终运行镜像
# -----------------------
FROM debian:bookworm-slim

WORKDIR /app

# 拷贝 Go 可执行文件
COPY --from=backend-build /app/server .
# 拷贝前端静态资源
COPY --from=backend-build /app/frontend/dist ./frontend/dist


# 运行端口（Go 默认 8000）
EXPOSE 8000

CMD ["./server"]
