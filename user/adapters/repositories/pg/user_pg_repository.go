package pg

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pozelim/product-review-app/common"
	"github.com/pozelim/product-review-app/user/domain"
)

type UserPgRepository struct {
	conn *pgx.Conn
}

func NewUserPgRepository(connString string) *UserPgRepository {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &UserPgRepository{
		conn: conn,
	}
}

func (r *UserPgRepository) Save(user domain.User) error {
	_, err := r.conn.Exec(
		context.Background(),
		"insert into \"user\" (username, password) values ($1, $2) returning username",
		user.Username,
		user.Password,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.ErrResourceAlreadyExists
			}
		}
		return err
	}
	return nil
}

func (r *UserPgRepository) Get(username string) (domain.User, error) {
	var user domain.User
	err := r.conn.QueryRow(context.Background(), "select * from \"user\" where username = $1", username).
		Scan(&user.Username, &user.Password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserPgRepository) Close() error {
	return r.conn.Close(context.Background())
}
