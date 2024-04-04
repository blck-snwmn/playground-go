package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/blck-snwmn/playground-go/testcontainers/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

var q *db.Queries

func TestMain(m *testing.M) {
	ctx := context.Background()

	mysqlContainer, err := mysql.RunContainer(
		ctx,
		testcontainers.WithImage("mysql:5.7"),
		testcontainers.WithEnv(map[string]string{
			"MYSQL_ROOT": "secret",
		}),
	)
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}
	defer mysqlContainer.Terminate(ctx)

	connStr, err := mysqlContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("Could not get connection string: %s", err)
	}
	mdb, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Could not open connection: %s", err)
	}

	_, err = mdb.Exec(db.Schema)
	if err != nil {
		log.Fatalf("Could not create table: %s", err)
	}

	q = db.New(mdb)

	code := m.Run()

	os.Exit(code)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	_, err := q.CreateUser(
		context.Background(), db.CreateUserParams{
			ID:   uuid.NewString(),
			Name: uuid.NewString(),
		},
	)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	t.Parallel()

	id := uuid.NewString()
	name := uuid.NewString()
	_, err := q.CreateUser(
		context.Background(), db.CreateUserParams{
			ID:   id,
			Name: name,
		},
	)
	assert.NoError(t, err)

	u, err := q.GetUser(context.Background(), id)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	assert.Equal(t, id, u.ID)
	assert.Equal(t, name, u.Name)
}
