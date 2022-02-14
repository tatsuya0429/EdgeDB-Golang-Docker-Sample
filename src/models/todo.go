package models

import (
	"github.com/edgedb/edgedb-go"
)

type Todo struct {
	ID          edgedb.UUID              `edgedb:"id"`
	Title       string                   `edgedb:"title"`
	Description string                   `edgedb:"description"`
	Status      edgedb.OptionalStr       `edgedb:"status"`
	Deadline    edgedb.OptionalLocalDate `edgedb:"deadline"`
	User        User                     `edgedb:"user"`
}
