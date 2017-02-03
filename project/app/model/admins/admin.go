package admins

import "github.com/mjibson/goon"

type Admin struct {
	AdminId   int    `datastore:"-" goon:"id"`
	Username  string `datastore:"username"`
	Email     string `datastore:"email"`
	Disabled  bool   `datastore:"disabled,noindex"`
	CreatedAt int    `datastore:"created_at,noindex"`
	UpdatedAt int    `datastore:"updated_at,noindex"`
}

func FindById(g goon.Goon, adminId int) (*Admin, error) {
	admin := &Admin{AdminId: adminId}
	if err := g.Get(admin); err != nil {
		return nil, err
	}
	return admin, nil
}
