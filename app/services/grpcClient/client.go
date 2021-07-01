package grpcClient

import (
	"fmt"
	"github.com/spf13/cast"
	"go-websocket/protobuf/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//func InitGrpcClient() {
//	// 连接服务器
//	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
//	if err != nil {
//		fmt.Printf("faild to connect: %v", err)
//	}
//	defer conn.Close()
//
//	c := pb.NewGreeterClient(conn)
//	// 调用服务端的SayHello
//	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "q1mi"})
//	if err != nil {
//		fmt.Printf("could not greet: %v", err)
//	}
//	fmt.Printf("Greeting: %s !\n", r.Message)
//}

func grpcConn(addr string) *grpc.ClientConn {
	// 连接服务器
	conn, err := grpc.Dial(addr+":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	return conn
}

//给单个用户发送信息
func SendUserMsg(addr string, toUserId int, data []byte) {
	fmt.Println("客户端SendUserMsg准备开始发送信息:", addr, toUserId, string(data))
	conn := grpcConn(addr)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	c.SendUserMsgToLocal(context.Background(), &pb.SendUserMsgToLocalRequest{
		UserId: cast.ToInt64(toUserId),
		Data:   data,
	})
}

//给群发送消息
func SendGroupMsgToLocal(addr string, toGroupId int, data []byte) {
	fmt.Println("客户端SendGroupMsgToLocal准备开始发送信息:", addr, toGroupId, string(data))
	conn := grpcConn(addr)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	c.SendGroupMsgToLocal(context.Background(), &pb.SendGroupMsgToLocalRequest{
		GroupId: cast.ToInt64(toGroupId),
		Data:    data,
	})
}

func SendMsgALl(addr string, data []byte) {
	conn := grpcConn(addr)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	c.SendAllMsgToLocal(context.Background(), &pb.SendAllMsgToLocalRequest{
		Data: data,
	})
}
