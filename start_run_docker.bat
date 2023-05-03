SET dirPath=%~dp0

docker run -i -t --name go-websocket-1 -p 12223:12223 -p 8972:8972 -v %dirPath%:/var/local -d go-websocket:v1