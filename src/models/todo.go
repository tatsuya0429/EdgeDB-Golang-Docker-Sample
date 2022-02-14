package models

import (
	"github.com/edgedb/edgedb-go"
)

type Todo struct {
	ID          edgedb.UUID
	Title       string
	Description string
	Status      string
	Deadline    edgedb.LocalDate
	User        User
}
