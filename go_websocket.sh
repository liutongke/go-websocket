#!/bin/bash

current_dir=$(pwd)
data_dir_log="$current_dir/log"

# 检查目录是否存在
if [ -d "$data_dir_log" ]; then
  echo "directory already exists: $data_dir_log"
else
  # 创建目录
  mkdir "$data_dir_log"
  echo "Create a directory: $data_dir_log"
fi

# 获取本机IP
ip=$(hostname -I | awk '{print $1}')
echo "本机的 IP 地址是：$ip"
docker build -t go-websocket:v1 .

docker run -itd --name go-websocket-v1 --restart=always -e MY_IP="$ip" -e DOCKER_IN=1 -p 12223:12223 -p 8972:8972 -v "$data_dir_log":/var/www/html/log go-websocket:v1
