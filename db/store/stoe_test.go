package store

import (
	"database/sql"
	"testing"
)

const (
	drivername   = "postgresq12"
	driverSource = "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable"
)

func TestTransferTx(t testing.T) {
	dbcon, err := sql.Open(drivername, driverSource)
	if err == nil {
		s := newStore(dbcon)
		account1 := gen.

	}

}
