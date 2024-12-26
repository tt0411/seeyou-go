   # 使用官方 Go 镜像作为基础镜像
   FROM golang:1.23 AS builder

   # 设置工作目录
   WORKDIR /app

   # 复制 go.mod 和 go.sum 文件
   COPY go.mod go.sum ./
   # 下载依赖
   RUN go mod download

   # 复制源代码
   COPY . .

   # 构建可执行文件
   RUN go build -o seeyou main.go

   # 使用轻量级的镜像来运行应用
   FROM alpine:latest

   # 设置工作目录
   WORKDIR /root/

   # 复制构建好的可执行文件
   COPY --from=builder /app/seeyou .

   # 暴露服务端口
   EXPOSE 8080

   # 运行可执行文件
   CMD ["./seeyou"]