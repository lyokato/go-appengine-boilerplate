package users

import "github.com/mjibson/goon"

type User struct {
	UserId    int    `datastore:"-" goon:"id"`
	Username  string `datastore:"username"`
	Email     string `datastore:"email"`
	Disabled  bool   `datastore:"disabled,noindex"`
	CreatedAt int    `datastore:"created_at,noindex"`
	UpdatedAt int    `datastore:"updated_at,noindex"`
}

func FindById(g goon.Goon, userId int) (*User, error) {
	user := &User{UserId: userId}
	if err := g.Get(user); err != nil {
		return nil, err
	}
	return user, nil
}
