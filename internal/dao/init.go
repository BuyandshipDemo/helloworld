package dao

import (
	"github.com/BuyandshipDemo/helloworld/internal/dao/mysql"
	"github.com/BuyandshipDemo/helloworld/internal/dao/redis"
)

func init() {
	mysql.Init()
	redis.Init()
}