package grpcService

import (
	"fmt"
	"go-websocket/config"
	"go-websocket/protobuf/pb"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SendToUserMsgToLocal(ctx context.Context, in *pb.SendToUserMsgToLocalRequest) (*pb.Send2SystemReply, error) {
	fmt.Println("服务端rpc调用成功")
	fmt.Println(in.Uid, in.UserId, in.Zone, string(in.Data))
	//websocket.SendToUserMsgLocal(cast.ToInt(in.UserId), cast.ToInt(in.Uid), cast.ToInt(in.Zone), in.Data)
	return &pb.Send2SystemReply{}, nil
}

func InitGrpcServer() {
	lis, err := net.Listen("tcp", ":"+config.GetConfClient().Server.RpcPort) // 监听本地的端口
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
		return
	}
}
