package service

import (
	"context"

	_ "github.com/BuyandshipDemo/helloworld/internal/dao"
	"github.com/BuyandshipDemo/helloworld/internal/dao/mysql"
	"github.com/BuyandshipDemo/helloworld/internal/dao/redis"
	helloworld "github.com/BuyandshipDemo/helloworld/kitex_gen/buyandship/marketplace/helloworld"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Greet(ctx context.Context, req *helloworld.GreetReq, resp *helloworld.GreetResp) (err error) {
	resp.BaseResp = &helloworld.BaseResp{}
	val, err := mysql.Query(req.Key)
	if err != nil {
		resp.BaseResp.Msg = err.Error()
		resp.BaseResp.Code = "1"
		return
	}
	res := redis.RedisCli.Set(ctx, req.Key, val, 0)
	_, err = res.Result()
	if err != nil {
		resp.BaseResp.Msg = err.Error()
		resp.BaseResp.Code = "2"
		return
	}
	resp.BaseResp.Code = "0"
	resp.BaseResp.Msg = "ok"
	resp.Val = val
	klog.CtxDebugf(ctx, "Greet Resp: %+v", resp)
	return nil
}