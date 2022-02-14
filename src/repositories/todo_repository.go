package repositories

import (
	"context"

	"github.com/edgedb/edgedb-go"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/infrastructure"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/models"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/utils"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (repo *TodoRepository) GetAll(ctx context.Context) (todos []models.Todo, err error) {
	client := infrastructure.NewDBClient(ctx)
	defer client.Close()
	query := "SELECT Todo{id, title, status, deadline, user: {id, username}}"
	err = client.Query(ctx, query, &todos)
	if err != nil {
		return
	}
	return
}

func (repo *TodoRepository) CreateTodo(ctx context.Context, title string, description string, deadlineStr string, userIdStr string) error {
	client := infrastructure.NewDBClient(ctx)
	defer client.Close()
	var inserted struct{ id edgedb.UUID }
	deadline, err := utils.StringToTime(deadlineStr)
	if err != nil {
		return err
	}
	localdate := edgedb.NewLocalDate(deadline.Date())
	if err != nil {
		return err
	}
	userId, err := edgedb.ParseUUID(userIdStr)
	if err != nil {
		return err
	}
	query := `
		INSERT Todo {
			title := <str>$0, 
			description := <str>$1,
			deadline := <cal::local_date>$2,
			user := (SELECT User FILTER .id = <uuid>$3)
		}
	`
	err = client.QuerySingle(ctx, query, &inserted, title, description, localdate, userId)
	if err != nil {
		return err
	}
	return nil
}
