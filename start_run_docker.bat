@echo off

set "current_dir=%cd%"
set "folder=%current_dir%/runtime"

if not exist "%folder%" (
    echo Creating %folder% folder...
    mkdir "%folder%"
    echo %folder% folder created.
) else (
    echo %folder% folder already exists.
)

docker build -t go-websocket:v1 .

docker run -itd --name go-websocket-v1 -p 12223:12223 -p 8972:8972 -v %folder%:/var/www/html/runtime go-websocket:v1