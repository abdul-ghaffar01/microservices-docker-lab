package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)


func ConnectDB() (*pgxpool.Pool, error) {
	// will read username, password host from environment variables
	connStr := "postgres://abdul:secret@localhost:5432/mydb"

	db, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db, nil

}

func CreateSchema(db *pgxpool.Pool){
	_, err := db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users(id uuid primary key not null, name text not null)")
	if err != nil {
		panic(err)
	}
}