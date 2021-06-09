package repo

import (
	"context"
	"fmt"

	"bproject/entity"
	"bproject/newDb"
	"bproject/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// NewGaussRepo return a new Gauss repo
func NewGaussRepo() Repository {
	fmt.Println("sss", newDb.ConnectDatabase)
	return &GaussRepo{
		db: newDb.ConnectDatabase,
	}
}

// GaussRepo Repository 接口类
type GaussRepo struct {
	db *gorm.DB
}

func (g GaussRepo) GetUserRepo(ctx context.Context, userID int64) (*entity.User, error) {
	res := &entity.User{}
	err := g.db.WithContext(ctx).Where(map[string]interface{}{"id": userID}).Find(res).Error

	if err != nil {
		return nil, utils.WrapError(err, "取user失败")
	}
	return nil, utils.WrapError(errors.New("自定义错误"), "cccc")
	//return res, nil
}

func (g GaussRepo) AddUserRepo(ctx context.Context, username string, fraction string) error {
	res := &entity.User{
		ID:       utils.GenID(),
		Username: username,
		Fraction: fraction}
	if err := g.db.WithContext(ctx).Create(
		res).Error; err != nil {
		return utils.WrapError(err, "创建user失败")
	}

	return nil
}
