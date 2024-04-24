package mysql

import (
	"github.com/BuyandshipDemo/helloworld/internal/entity"
)

func Query(key string) (string, error) {
	db := DB.Model(entity.Helloworld{})
	var res *entity.Helloworld
	db = db.Where("`key` = ?", key).Debug().Find(&res)
	if db.Error != nil {
		return "", db.Error
	}
	return res.Val, nil
}