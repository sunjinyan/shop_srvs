package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "shop_srv/user_srv/proto/api/gen/v1"
	"testing"
)

func TestHandler_GetUserList(t *testing.T) {

	conn, err := grpc.Dial("127.0.0.1:8090",grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	cli := proto.NewUserClient(conn)
	list, err := cli.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(list)
}