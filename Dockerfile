FROM golang:1.20.4-bullseye

EXPOSE 12223
EXPOSE 8972
EXPOSE 20000

WORKDIR /var/www/html

RUN sed -i 's/deb.debian.org/mirrors.tencent.com/g' /etc/apt/sources.list
#设置时区
ENV TZ=Asia/Shanghai

RUN apt-get update \
    && apt-get install libgl1-mesa-glx libglib2.0-0 wget -y \
    && apt-get install supervisor -y \
    && apt-get install python3 python3-dev python3-pip -y \
    && pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

COPY go_websocket.conf /etc/supervisor/conf.d/

# 设置环境变量 GOPROXY
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

#看需求是否将运行程序复制进容器中,不复制的话在创建时候需要-v绑定文件夹
COPY . .

# 下载所依赖的模块
#RUN go mod download

# 将我们的代码编译成二进制可执行文件main
#RUN go build -o main .

#-n 是 supervisord 命令的一个选项，它表示以非守护进程模式（non-daemon mode）运行 supervisord。
#CMD [ "supervisord" , "-n", "-c", "/etc/supervisor/supervisord.conf" ]