# 多阶段构建
# 阶段1: 构建
FROM golang:1.22-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的构建工具
RUN apk add --no-cache git

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fairytale-creator .

# 阶段2: 运行
FROM alpine:latest

# 安装ca证书（用于HTTPS请求）
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/fairytale-creator .

# 创建必要的目录
RUN mkdir -p /app/voices /app/stories /app/images

# 暴露端口
EXPOSE 9700

# 启动应用
CMD ["./fairytale-creator"]

