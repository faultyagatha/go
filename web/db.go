package main

import (
	"time"
)

type User struct {
	ID int64 `json:"id"`
	// can not be empty
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//using pg package
// func (u *User) Create(ctx context.Context, db *sql.DB) (pg.User, error) {
// return pg.New(db).CreateUser(ctx, "Someone")
// it'l return an empty struct in case of an err as if: return pg.User{}, err
// }
