package api

import (
	"database/sql"

	"mono.thienhang.com/pkg/database"
)

type Model struct {
	TableName string

	Conn database.Connection
	Tx   *sql.Tx
}

func (b Model) SetConn(con database.Connection) Model {
	b.Conn = con
	return b
}

func (b Model) Table(table string) *database.SQL {
	return database.Table(table).WithDriver(b.Conn)
}
