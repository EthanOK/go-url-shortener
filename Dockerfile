# 第一阶段：编译 Go 项目
FROM golang:1.22-alpine

# 设置工作目录
WORKDIR /app

# 复制 Go 项目文件到容器中
COPY . .

ENV DOCKER_ENV=true

# 安装依赖 redis
RUN apk add --no-cache redis

RUN go mod download

RUN go build -o app .


# 暴露端口
EXPOSE 9808

# 运行应用
CMD ["./app"]
