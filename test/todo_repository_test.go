package test

import (
	"context"
	"fmt"
	"github.com/mrizkisaputra/golang-database/config"
	"github.com/mrizkisaputra/golang-database/entity"
	"github.com/mrizkisaputra/golang-database/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

var repository = repositories.NewTodoRepository(config.GetConnection())

func TestInsertNotAlreadyId(t *testing.T) {
	var todo = entity.Todo{
		Id:          2,
		Title:       "Belajar Golang Web",
		Description: "sedang belajar golang web",
	}
	err := repository.Insert(context.Background(), todo)
	assert.Nil(t, err)
}

func TestInsertWithIdAlreadyExist(t *testing.T) {
	var todo = entity.Todo{
		Id:          1,
		Title:       "Belajar Golang dasar",
		Description: "sedang belajar golang dasar",
	}
	err := repository.Insert(context.Background(), todo)
	assert.NotNil(t, err)
}

func TestGetAll(t *testing.T) {
	todos, err := repository.GetAll(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, todos)
	fmt.Println(todos)
}
