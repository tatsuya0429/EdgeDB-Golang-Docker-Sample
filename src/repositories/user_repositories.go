package repositories

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/infrastructure"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) CreateUser(ctx context.Context, username string, password string) error {
	client := infrastructure.NewDBClient(ctx)
	defer client.Close()
	var inserted struct{ id edgedb.UUID }
	query := `INSERT User {
		username := <str>$0, 
		password := <str>$1
	}`
	err := client.QuerySingle(ctx, query, &inserted, username, password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, userId string) (user models.User, err error) {
	client := infrastructure.NewDBClient(ctx)
	query := "SELECT User filter .id=<uuid>$0"
	err = client.QuerySingle(ctx, query, &user, userId)
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUserByUsername(ctx context.Context, username string) (user models.User, err error) {
	client := infrastructure.NewDBClient(ctx)
	query := "SELECT User{id, username, password} FILTER .username = <str>$0"
	err = client.QuerySingle(ctx, query, &user, username)
	fmt.Print(user)
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetAll(ctx context.Context) (users []models.User, err error) {
	client := infrastructure.NewDBClient(ctx)
	query := "SELECT User{id, username, password}"
	err = client.Query(ctx, query, &users)
	if err != nil {
		return
	}
	return
}
