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