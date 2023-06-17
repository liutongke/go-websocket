@echo off

set "current_dir=%cd%"
set "folder=%current_dir%/runtime"
set "data_dir_log=%current_dir%\log"

if not exist "%folder%" (
    echo Creating %folder% folder...
    mkdir "%folder%"
    echo %folder% folder created.
) else (
    echo %folder% folder already exists.
)

REM 检查目录是否存在
if exist "%data_dir_log%" (
  echo directory already exists: %data_dir_log%
) else (
  REM 创建目录
  mkdir "%data_dir_log%"
  echo Create a directory: %data_dir_log%
)

docker build -t go-websocket:v1 .

docker run -it --name go-websocket-v1 -p 12223:12223 -p 8972:8972 -v %current_dir%:/var/www/html go-websocket:v1