package infrastructure

import (
	"context"
	"log"

	"github.com/edgedb/edgedb-go"
)

func NewDBClient(ctx context.Context) *edgedb.Client {
	config := NewDBConfig()
	opts := edgedb.Options{
		Database: config.DBName,
		Host: config.Host,
		User: config.User,
		Password: edgedb.NewOptionalStr(config.Password),
		Port: config.Port,
		TLSOptions: edgedb.TLSOptions{
			SecurityMode: edgedb.TLSModeInsecure,
		},
		Concurrency: 4,
	}
	client, err := edgedb.CreateClient(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

