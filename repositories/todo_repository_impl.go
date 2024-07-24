package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/mrizkisaputra/golang-database/entity"
	"github.com/mrizkisaputra/golang-database/utilities"
)

type todoRepositoryImpl struct {
	connDB *sql.DB
}

func NewTodoRepository(connDB *sql.DB) TodoRepository {
	return &todoRepositoryImpl{connDB: connDB}
}

func (repository *todoRepositoryImpl) Insert(ctx context.Context, todo entity.Todo) error {
	insertSQL := "INSERT INTO todo(id, title, description) values (?,?,?)"
	selectSQL := "SELECT id FROM todo WHERE id = ?"
	tx, err := repository.connDB.Begin() // start transactional
	if err != nil {
		return utilities.StartTransactionError
	}

	rows, err := tx.QueryContext(ctx, selectSQL, todo.Id)
	if rows.Next() {
		return errors.New("todo with the id already exists")
	}

	_, err = tx.ExecContext(ctx, insertSQL, todo.Id, todo.Title, todo.Description) // Tx 1
	if err != nil {
		return tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (repository todoRepositoryImpl) GetAll(ctx context.Context) ([]entity.Todo, error) {
	db := repository.connDB
	defer db.Close()

	selectSQL := "SELECT id,title, description, created_at FROM todo"
	rows, err := db.QueryContext(ctx, selectSQL)
	if err != nil {
		return nil, err
	}

	var todos []entity.Todo
	for rows.Next() {
		todo := entity.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
