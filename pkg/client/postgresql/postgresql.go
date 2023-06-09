package postgresql

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/while-act/hackathon-backend/ent"
	_ "github.com/while-act/hackathon-backend/ent/runtime"
	"github.com/while-act/hackathon-backend/pkg/log"
)

// Open postgres connection, check it and create tables (if not exist)
func Open(username, password, host string, port int, DBName string) *ent.Client {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		username, password, host, port, DBName)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.WithErr(err).Fatal("error occurred while opening PostgreSQL connection")
	}

	if err = db.Ping(); err != nil {
		log.WithErr(err).Fatal("unable to connect to the postgres database")
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err = client.Schema.Create(context.Background()); err != nil {
		log.WithErr(err).Fatal("tables initialization failed")
	}

	return client
}
