package models

import "github.com/edgedb/edgedb-go"

type User struct {
	Id       edgedb.UUID `edgedb:"id"`
	Username string      `edgedb:"username"`
	Password string      `edgedb:"password"`
}
