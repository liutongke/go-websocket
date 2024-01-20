#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import platform
import re
import subprocess


def get_os():
    os_name = platform.system()
    if os_name == "Windows":
        return True
    elif os_name == "Linux":
        return False


def windows():
    # # 使用subprocess模块执行ipconfig命令，并使用findstr筛选包含"IPv4"的行
    # result = subprocess.check_output("ipconfig | findstr /i \"IPv4\"", shell=True)
    #
    # # 将输出结果按换行符分割成列表
    # lines = result.decode("gbk").split("\r\n")
    #
    # # 遍历列表，找到包含"IPv4"的行，并提取第16个字段
    # for line in lines:
    #     if "IPv4" in line:
    #         myip = line.split()[15]
    #         break
    # return myip
    # 执行 ipconfig 命令，并以 GBK 编码获取输出（适用于中文 Windows）
    result = subprocess.run(['ipconfig'], capture_output=True, text=True, encoding='gbk')
    ipconfig_output = result.stdout

    # 使用正则表达式匹配 WLAN 网络适配器的部分
    wlan_match = re.search(r'WLAN[^\n]*\n(.*?IPv4 地址[^\n]*)', ipconfig_output, re.DOTALL)
    if wlan_match:
        # 从匹配结果中提取 IPv4 地址
        ipv4_match = re.search(r'IPv4[^\d]*(\d+\.\d+\.\d+\.\d+)', wlan_match.group(1))
        if ipv4_match:
            return ipv4_match.group(1)


def linux():
    # 获取本机IP
    ip = subprocess.check_output(["hostname", "-I"]).decode().strip().split()[0]
    return ip


# 获取当前目录
current_dir = os.getcwd()
data_dir = os.path.join(current_dir, 'log')

# 检查目录是否存在
if os.path.exists(data_dir):
    print(f"Directory already exists: {data_dir}")
else:
    # 创建目录
    os.mkdir(data_dir)
    print(f"Created directory: {data_dir}")

if get_os():
    myip = windows()
else:
    myip = linux()

print(f"本机的 IP 地址是：{myip}")

# 构建 Docker 镜像名称
image_name = 'go-websocket:v1'

# 构建 Docker 镜像
subprocess.run(['docker', 'build', '-f', './Dockerfile', '-t', image_name, "."], check=True)

# 运行 Docker 容器
subprocess.run(
    [
        'docker',
        'run',
        '--name', 'go-websocket-v1',
        '--restart=always',
        '-e', f'MY_IP={myip}',
        '-e', 'DOCKER_IN=1',
        '-itd',
        '-p', '12223:12223',
        '-p', '8972:8972',
        '-v', f'{current_dir}/log:/var/www/html/log',
        image_name], check=True)
