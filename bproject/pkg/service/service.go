package service

import (
	"context"

	"bproject/entity"
	"bproject/repo"
	"github.com/pkg/errors"
)

// BProjectService describes the service.
type BProjectService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	AddUser(ctx context.Context, username string, fraction string) (err error)
	GetUser(ctx context.Context, userID int64) (res *entity.User, err error)
}

type basicBProjectService struct {
	repo repo.Repository
}

func (b *basicBProjectService) AddUser(ctx context.Context, username string, fraction string) (err error) {
	if err := b.repo.AddUserRepo(ctx, username, fraction); err != nil {
		return err
	}
	return nil
}
func (b *basicBProjectService) GetUser(ctx context.Context, userID int64) (res *entity.User, err error) {
	//err = foo(1, 2)
	//if err != nil {
	//	return nil, err
	//}
	res, err = b.repo.GetUserRepo(ctx, userID)
	if err != nil {
		return nil, errors.WithMessage(err, "调用repo中的函数")
	}
	return res, nil
}

func foo(i int, i2 int) error {
	return errors.Wrap(errors.New("cesji "), "ceshi")
}

// NewBasicBProjectService returns a naive, stateless implementation of BProjectService.
func NewBasicBProjectService() BProjectService {
	return &basicBProjectService{
		repo: repo.NewGaussRepo(),
	}
}

// New returns a BProjectService with all of the expected middleware wired in.
func New(middleware []Middleware) BProjectService {
	var svc BProjectService = NewBasicBProjectService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
