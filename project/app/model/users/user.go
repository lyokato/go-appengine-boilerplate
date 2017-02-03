package users

import "errors"

type User struct {
	UserId    int    `datastore:"-" goon:"id"`
	Username  string `datastore:"username"`
	Email     string `datastore:"email"`
	Disabled  bool   `datastore:"disabled,noindex"`
	CreatedAt int    `datastore:"created_at,noindex"`
	UpdatedAt int    `datastore:"updated_at,noindex"`
}

func FindById(userId int) (*User, error) {
	return nil, errors.New("not implemented yet")
}

func FindByUsername(username string) (*User, error) {
	return nil, errors.New("not implemented yet")
}
