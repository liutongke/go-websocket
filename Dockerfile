FROM golang:1.19.1
#https://learnku.com/go/wikis/38122
EXPOSE 9500
EXPOSE 9501
EXPOSE 9502
EXPOSE 9503

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

COPY / /var/local

WORKDIR /var/local

#CMD ["php","./apiswoole.php"]

#docker build -t go:v1 ./