package txdb

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/blck-snwmn/playground-go/testcontainers/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

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
	txdb.Register("txdb", "mysql", connStr)

	_, err = mdb.Exec(db.Schema)
	if err != nil {
		log.Fatalf("Could not create table: %s", err)
	}

	code := m.Run()

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
