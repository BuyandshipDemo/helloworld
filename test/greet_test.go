package test

import (
	"context"
	"testing"

	"github.com/BuyandshipDemo/helloworld/kitex_gen/buyandship/marketplace/helloworld"
	"github.com/BuyandshipDemo/helloworld/kitex_gen/buyandship/marketplace/helloworld/greetservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func TestGreet(t *testing.T) {
	klog.Info("Testing...")
	cli, err := greetservice.NewClient("helloworld", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Error(err)
	}
	req := helloworld.NewGreetReq()
	req.SetKey("bravo")
	resp, err := cli.Greet(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	klog.Info(resp)
}
