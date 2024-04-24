// Code generated by Kitex v0.9.1. DO NOT EDIT.

package greetservice

import (
	"context"
	"errors"
	helloworld "github.com/BuyandshipDemo/helloworld/kitex_gen/buyandship/marketplace/helloworld"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Greet": kitex.NewMethodInfo(
		greetHandler,
		newGreetServiceGreetArgs,
		newGreetServiceGreetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	greetServiceServiceInfo                = NewServiceInfo()
	greetServiceServiceInfoForClient       = NewServiceInfoForClient()
	greetServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return greetServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return greetServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return greetServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "GreetService"
	handlerType := (*helloworld.GreetService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "helloworld",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func greetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*helloworld.GreetServiceGreetArgs)
	realResult := result.(*helloworld.GreetServiceGreetResult)
	success, err := handler.(helloworld.GreetService).Greet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGreetServiceGreetArgs() interface{} {
	return helloworld.NewGreetServiceGreetArgs()
}

func newGreetServiceGreetResult() interface{} {
	return helloworld.NewGreetServiceGreetResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Greet(ctx context.Context, req *helloworld.GreetReq) (r *helloworld.GreetResp, err error) {
	var _args helloworld.GreetServiceGreetArgs
	_args.Req = req
	var _result helloworld.GreetServiceGreetResult
	if err = p.c.Call(ctx, "Greet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
