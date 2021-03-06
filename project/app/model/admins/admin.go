package admins

import (
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Admin struct {
	AdminId  int64  `datastore:"-" goon:"id"`
	Email    string `datastore:"email"`
	IsSuper  bool   `datastore:"is_super"`
	Disabled bool   `datastore:"disabled,noindex"`
}

func All(c context.Context) ([]Admin, error) {
	q := datastore.NewQuery("Admin")
	// .Filter().Order()
	var list []Admin
	if _, err := q.GetAll(c, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func FindById(g *goon.Goon, adminId int64) (*Admin, error) {
	admin := &Admin{AdminId: adminId}
	if err := g.Get(admin); err != nil {
		return nil, err
	}
	return admin, nil
}

func FindByEmail(g *goon.Goon, email string) (*Admin, error) {
	admin := &Admin{Email: email}
	if err := g.Get(admin); err != nil {
		return nil, err
	}
	return admin, nil
}

func Create(g *goon.Goon, email string, isSuper bool) (*Admin, error) {
	admin := &Admin{Email: email, IsSuper: isSuper}
	if _, err := g.Put(admin); err != nil {
		return nil, err
	}
	return admin, nil
}
