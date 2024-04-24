package main

import (
	"context"

	"github.com/BuyandshipDemo/helloworld/internal/service"
	helloworld "github.com/BuyandshipDemo/helloworld/kitex_gen/buyandship/marketplace/helloworld"

	"github.com/cloudwego/kitex/pkg/klog"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// Greet implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) Greet(ctx context.Context, req *helloworld.GreetReq) (resp *helloworld.GreetResp, err error) {
	klog.CtxInfof(ctx, "req: %+v", req)
	resp = helloworld.NewGreetResp()
	err = service.Greet(ctx, req, resp)
	return
}
