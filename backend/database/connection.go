package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct{
	Connection *pgxpool.Pool
}

func (db *Database) Connect() error{
	connectionPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil{
		return err
	}

	db.Connection = connectionPool

	return nil
}
