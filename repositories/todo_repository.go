package repositories

import (
	"context"
	"github.com/mrizkisaputra/golang-database/entity"
)

type TodoRepository interface {
	Insert(ctx context.Context, todo entity.Todo) error
	GetAll(ctx context.Context) ([]entity.Todo, error)
}
