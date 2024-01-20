# go-WebSocket
基于gin框架实现的WebSocket聊天系统

WebSocket请求：

1. 直接GET请求不需要携带任何参数获取请求token `http://192.168.1.106:12223/User/GetInfo`

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

### 要查看 Docker 容器中的环境变量（ENV 变量）
```
docker exec <容器名称或容器ID> env
```

go语言获取环境变量:

```go
ip := os.Getenv("MY_IP")
fmt.Println("IP Address:", ip)
```

`-e DOCKER_IN=1`设置当前服务是否在容器内运行


go get go.etcd.io/etcd/client/v3


### Linux系统安装protobuf

如果你想要安装特定版本的 Protocol Buffers 编译器（protoc），而不是使用系统的包管理器提供的版本，你可以从 [Protocol Buffers GitHub Releases](https://github.com/protocolbuffers/protobuf/releases) 页面下载预编译的二进制文件。

以下是在 Linux 系统上安装 protoc 3.17.3 版本的步骤。请根据你的操作系统选择合适的二进制文件。

1. 打开终端或命令行窗口。

2. 使用以下命令下载 protoc 3.17.3 版本的二进制文件（假设你的系统是 64 位 Linux）：

   ```bash
   wget https://github.com/protocolbuffers/protobuf/releases/download/v3.17.3/protoc-3.17.3-linux-x86_64.zip
   ```

   如果你的系统不是 64 位 Linux，可以从 [GitHub Releases](https://github.com/protocolbuffers/protobuf/releases) 页面下载其他适用于你系统的文件。[v3.17.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3)

3. 解压下载的 zip 文件：

   ```bash
   unzip protoc-3.17.3-linux-x86_64.zip
   ```

4. 将 protoc 可执行文件移动到你的 PATH 中，以便全局使用：

   ```bash
    mv bin/protoc /usr/local/bin/
   ```

5. 验证 protoc 安装是否成功：

   ```bash
   protoc --version
   ```

   这应该显示安装的 protoc 的版本信息（3.17.3）。

通过这些步骤，你就可以在 Linux 系统上安装并使用 protoc 3.17.3 版本。请注意，如果你使用的是其他操作系统，请下载相应的二进制文件并按照类似的步骤进行安装。


### protobuf使用

**输出文件会自动创建个`pb`目录**
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

protoc -I ./protobuf ./protobuf/websocket.proto --go_out=./app/services/grpc --go-grpc_out=./app/services/grpc
```

- -I ./protobuf：指定了导入路径，告诉编译器在 ./protobuf 目录中查找导入的文件。
- ./protobuf/websocket.proto：指定要编译的 Protocol Buffers 文件路径。
- --go_out=./app\services\grpc：指定生成的 Go 代码的输出目录，这里是`./app\services\grpc`目录。
- --go-grpc_out=./app\services\grpc：指定生成的 gRPC 相关的 Go 代码的输出目录，也是`./protobuf/pb`目录。
所以，该命令的目的是将 websocket.proto 文件编译为 Go 代码，并将生成的代码输出到`./app\services\grpc`目录中。生成的代码将包括 Protobuf 的基本消息类型以及与 gRPC 相关的服务和客户端代码。

请注意，您需要确保已经安装了 protoc 编译器，并且已经安装了 protoc-gen-go 和 protoc-gen-go-grpc 插件。这些插件负责生成 Go 代码和 gRPC 相关的代码。