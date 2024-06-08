package simple_bank_db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatalf("Cannot connect to db: %s", err)
	}

	defer conn.Close(context.Background())

	testQueries = New(conn)

	os.Exit(m.Run())
}
