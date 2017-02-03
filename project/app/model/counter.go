package model

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

func CountDatastoreModel(ctx context.Context, model string) (int, error) {
	query := datastore.NewQuery(model)
	n, err := query.Count(ctx)
	if err != nil {
		return -1, err
	}
	return n, nil
}
