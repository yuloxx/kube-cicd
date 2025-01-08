# 第一阶段：构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将当前目录的所有文件复制到容器中
COPY . .

# 下载依赖（可选：加快构建速度）
RUN go mod tidy

# 构建 Go 应用程序
RUN go build -o server .

# 第二阶段：运行阶段
FROM alpine:latest

# 安装必要的依赖（如果需要 TLS 支持，可安装 ca-certificates）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 从构建阶段复制应用程序
COPY --from=builder /app/server .

# 监听 8080 端口
EXPOSE 8080

# 运行应用程序
CMD ["./server"]