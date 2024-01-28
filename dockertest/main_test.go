package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/blck-snwmn/playground-go/dockertest/dockertest/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

func TestMain(m *testing.M) {

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	txdb.Register("txdb", "mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))

	var mdb *sql.DB
	if err := pool.Retry(func() error {
		var err error
		mdb, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return mdb.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	_, err = mdb.Exec(db.Schema)
	if err != nil {
		log.Fatalf("Could not create table: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCreate1(t *testing.T) {
	t.Parallel()
	txdb, err := sql.Open("txdb", "xxx")
	if err != nil {
		t.Fatal(err)
	}
	defer txdb.Close()

	q := db.New(txdb)
	_, err = q.CreateUser(context.Background(), db.CreateUserParams{
		ID:   "411c34d8-0110-4b2a-85d6-104603562b83",
		Name: "xxx_name_1",
		Bio:  sql.NullString{String: "xxx1", Valid: true},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate2(t *testing.T) {
	t.Parallel()
	txdb, err := sql.Open("txdb", "yyy") // not xxx
	if err != nil {
		t.Fatal(err)
	}
	defer txdb.Close()

	q := db.New(txdb)
	_, err = q.CreateUser(context.Background(), db.CreateUserParams{
		ID:   "411c34d8-0110-4b2a-85d6-104603562b83",
		Name: "xxx_name_2",
		Bio:  sql.NullString{String: "xxx2", Valid: true},
	})
	if err != nil {
		t.Fatal(err)
	}
}
