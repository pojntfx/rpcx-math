package client

import (
	"context"
	"github.com/pojntfx/rpcx-math/math/proto"
	"github.com/smallnest/rpcx/client"
)

func Add(client client.XClient, first, second int) (int, error) {
	args := &proto.MathAddArgs{
		First:  int64(first),
		Second: int64(second),
	}
	reply := &proto.MathAddReply{}

	call, err := client.Go(context.Background(), "Add", args, reply, nil)

	if err != nil {
		return 0, err
	}

	replyCall := <-call.Done

	if replyCall.Error != nil {
		return 0, replyCall.Error
	}
	return int(reply.Result), nil
}

func Subtract(client client.XClient, first, second int) (int, error) {
	args := &proto.MathSubtractArgs{
		First:  int64(first),
		Second: int64(second),
	}
	reply := &proto.MathSubtractReply{}

	call, err := client.Go(context.Background(), "Subtract", args, reply, nil)

	if err != nil {
		return 0, err
	}

	replyCall := <-call.Done

	if replyCall.Error != nil {
		return 0, replyCall.Error
	}
	return int(reply.Result), nil
}
