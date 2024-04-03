# 第一阶段：构建阶段
FROM golang:1.22-bullseye as builder

WORKDIR /var/www/html

RUN sed -i 's/deb.debian.org/mirrors.tencent.com/g' /etc/apt/sources.list

# 设置时区
#ENV TZ=Asia/Shanghai

# 设置环境变量 GOPROXY
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0\
    DEBUG=0

COPY . .

# 下载所依赖的模块
RUN go mod download
# 将我们的代码编译成二进制可执行文件 main
RUN CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' -o main .


# 第二阶段：运行阶段
FROM scratch

WORKDIR /app

# 从构建阶段复制二进制可执行文件 main 到运行阶段
COPY --from=builder /var/www/html/main /app/main

# 从构建阶段复制 config 目录内容到运行阶段的 /app/config
COPY --from=builder /var/www/html/config/config_line.toml /app/config/config_line.toml
COPY --from=builder /var/www/html/config/config.toml /app/config/config.toml
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo


ENTRYPOINT ["/app/main"]
