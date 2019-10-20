package service

import (
	"context"
	"github.com/pojntfx/rpcx-math/math/proto"
)

type Math int

func (t *Math) Add(ctx context.Context, args *proto.MathAddArgs, reply *proto.MathAddReply) error {
	reply.Result = args.First + args.Second
	return nil
}

func (t *Math) Subtract(ctx context.Context, args *proto.MathSubtractArgs, reply *proto.MathSubtractReply) error {
	reply.Result = args.Second - args.First
	return nil
}
