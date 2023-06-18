# go-WebSocket
基于gin框架实现的WebSocket聊天系统

WebSocket请求：

1. 获取请求token `http://192.168.1.106:12223/User/GetInfo`

2. 发送WebSocket请求
    ```
    ws://192.168.1.106:12223/ws/
    ```
    请求头携带`X-Token`

    ![Img](https://raw.githubusercontent.com/liutongke/Image-Hosting/master/images/yank-note-picgo-img-20230613013752.png)

3. 请求体
    ```json
    {"id":123,"path":"/ping","ver":"1.0.0","data":""}
    ```

    请求体解析:
    ```
    id 客户端消息唯一id
    path 请求路由
    ver 客户端版本
    data 消息体
    ```
    
docker run时候使用`-e MY_IP=%myip%`将 IP 地址作为环境变量 MY_IP 传递给 Docker 容器

go语言获取环境变量:

```go
	ip := os.Getenv("MY_IP")
	fmt.Println("IP Address:", ip)
```

`-e DOCKER_IN=1`设置当前服务是否在容器内运行


**输出文件会自动创建个`pb`目录**
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


protoc -I ./protobuf ./protobuf/websocket.proto -go-grpc_out=:./protobuf/pb
protoc -I ./protobuf ./protobuf/websocket.proto --go_out=./protobuf --go-grpc_out=./protobuf
protoc -I ./protobuf ./protobuf/websocket.proto --go_out=./app\services\grpc --go-grpc_out=./app\services\grpc
```

- -I ./protobuf：指定了导入路径，告诉编译器在 ./protobuf 目录中查找导入的文件。
- ./protobuf/websocket.proto：指定要编译的 Protocol Buffers 文件路径。
- --go_out=./protobuf/pb：指定生成的 Go 代码的输出目录，这里是 ./protobuf/pb 目录。
- --go-grpc_out=./protobuf/pb：指定生成的 gRPC 相关的 Go 代码的输出目录，也是 ./protobuf/pb 目录。
所以，该命令的目的是将 websocket.proto 文件编译为 Go 代码，并将生成的代码输出到 ./protobuf/pb 目录中。生成的代码将包括 Protobuf 的基本消息类型以及与 gRPC 相关的服务和客户端代码。

请注意，您需要确保已经安装了 protoc 编译器，并且已经安装了 protoc-gen-go 和 protoc-gen-go-grpc 插件。这些插件负责生成 Go 代码和 gRPC 相关的代码。

go get go.etcd.io/etcd/client/v3