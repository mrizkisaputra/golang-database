package test

import (
	"context"
	"github.com/mrizkisaputra/golang-database/config"
	"github.com/mrizkisaputra/golang-database/entity"
	"github.com/mrizkisaputra/golang-database/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

var repository = repositories.NewTodoRepository(config.GetConnection())

func TestInsertNotAlreadyData(t *testing.T) {
	var todo = entity.Todo{
		Id:          1,
		Title:       "Belajar Golang dasar",
		Description: "sedang belajar golang dasar",
	}
	err := repository.Insert(context.Background(), todo)
	assert.Nil(t, err)
}

func TestWithIdAlreadyExist(t *testing.T) {
	var todo = entity.Todo{
		Id:          1,
		Title:       "Belajar Golang dasar",
		Description: "sedang belajar golang dasar",
	}
	err := repository.Insert(context.Background(), todo)
	assert.NotNil(t, err)
}
