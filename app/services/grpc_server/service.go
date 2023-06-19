package grpc_server

import (
	"fmt"
	"github.com/spf13/cast"
	"go-websocket/app/services/grpc/pb"
	"go-websocket/app/services/websocket"
	"go-websocket/config"
	"go-websocket/tools"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// 给本机单个用户发送消息
func (s *server) SendUserMsgToLocal(ctx context.Context, in *pb.SendUserMsgToLocalRequest) (*pb.SendUserResponse, error) {
	fmt.Printf("服务端SendToUserMsgToLocal rpc调用成功 - userId:%d,msg:%s", in.UserId, string(in.Data))
	websocket.NewRpcService().SendUserMsgLocal(cast.ToInt(in.UserId), in.Data)
	return &pb.SendUserResponse{}, nil
}

// 给本机分组用户发送消息
func (s *server) SendGroupMsgToLocal(ctx context.Context, in *pb.SendGroupMsgToLocalRequest) (*pb.SendGroupResponse, error) {
	fmt.Printf("服务端SendToGroupMsgToLocal rpc调用成功 - GroupId:%d,msg:%s", in.GroupId, string(in.Data))
	websocket.NewRpcService().SendToGroupMsgToLocal(cast.ToInt(in.GroupId), in.Data)
	return &pb.SendGroupResponse{}, nil
}

// 给全服用户发送消息
func (s *server) SendAllMsgToLocal(ctx context.Context, in *pb.SendAllMsgToLocalRequest) (*pb.SendAllResponse, error) {
	fmt.Printf("服务端SendAllMsgToLocal rpc调用成功 - msg:%s", string(in.Data))
	websocket.NewRpcService().SendMsgALlLocal(in.Data)
	return &pb.SendAllResponse{}, nil
}

func InitGrpcServer() {
	lis, err := net.Listen("tcp", ":"+config.GetConf().Grpc.RpcPort) // 监听本地的端口
	if err != nil {
		tools.EchoError(fmt.Sprintf("GRPC server failed to listen: %v", err))
	}
	//log.Printf("开启的RPC端口-------------->：%s \n", config.GetConf().Server.RpcPort)
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	//log.Printf("开启的RPC端口-------------->：%s \n", config.GetConf().Grpc.RpcPort, err)
	err = s.Serve(lis)

	if err != nil {
		tools.EchoError(fmt.Sprintf("GRPC server failed to serve: %v", err))
	}

}
