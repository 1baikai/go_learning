package repo

import (
	"context"

	"bproject/entity"
)

type Repository interface {
	AddUserRepo(ctx context.Context, username string, fraction string) error
	GetUserRepo(ctx context.Context, userID int64) (*entity.User, error)
}
