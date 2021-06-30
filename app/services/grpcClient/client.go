package grpcClient

import (
	"go-websocket/protobuf/pb"
	"fmt"
	"github.com/spf13/cast"
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

//发送用户给单个用户
func SendToUserMsg(addr string, toUserId, uid, zone int, data []byte) {
	fmt.Println("客户端准备开始发送信息:", addr, toUserId, uid, string(data))
	conn := grpcConn(addr)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	c.SendToUserMsgToLocal(context.Background(), &pb.SendToUserMsgToLocalRequest{
		UserId: cast.ToInt64(toUserId),
		Uid:    cast.ToInt64(uid),
		Zone:   cast.ToInt64(zone),
		Data:   data,
	})
}
