package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRespository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateNewUser(ctx context.Context, u User) (*User, error) {

	fmt.Println(u)

	row := ur.db.QueryRow(ctx, "INSERT INTO users(id, name) VALUES($1, $2) RETURNING id, name", u.ID, u.Name)

	var user User

	row.Scan(&user.ID, &user.Name)

	return &user, nil

}

func (ur *UserRepository) GetUsers(ctx context.Context) (*[]User, error) {

	rows, err := ur.db.Query(ctx, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	return &users, nil
}
