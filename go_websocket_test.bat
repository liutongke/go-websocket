@echo off

set "current_dir=%cd%"
set "data_dir_log=%current_dir%\log"

REM 检查目录是否存在
if exist "%data_dir_log%" (
  echo directory already exists: %data_dir_log%
) else (
  REM 创建目录
  mkdir "%data_dir_log%"
  echo Create a directory: %data_dir_log%
)

for /f "tokens=16" %%i in ('ipconfig ^|find /i "ipv4"') do (
set myip=%%i
:: 正常情况下find查询只有一行结果，如果主机安装了虚拟机则会有多个适配器有ip地址。第一个才是本机IP，故使用goto保证for只执行一次就跳出循环，防止后续myip的值被覆盖
goto out
)

:: 标签
:out


docker build -t go-websocket:v1 .

docker run -it --name go-websocket-v1 -e MY_IP=%myip% -e DOCKER_IN=1 -p 12223:12223 -p 8972:8972 -v %current_dir%:/var/www/html go-websocket:v1