package mongo

import (
	"context"
	"project/github/comics/dbhandler/models"
	"project/github/comics/utils"
)

func (m *Mongo) Create(ctx context.Context, col string, req *models.CreateRequest) (int64, error) {
	// Create a collection object
	collection := m.getClient().Database(m.dbName).Collection(col)

	switch req.Operation {
	case utils.One:
		// Insert single document
		_, err := collection.InsertOne(ctx, req.Document)
		if err != nil {
			return 0, err
		}

		return 1, nil

	case utils.All:
		// Insert multiple documents
		objs, ok := req.Document.([]interface{})
		if !ok {
			return 0, utils.ErrInvalidParams
		}

		res, err := collection.InsertMany(ctx, objs)
		if err != nil {
			return 0, err
		}

		return int64(len(res.InsertedIDs)), nil

	default:
		return 0, utils.ErrInvalidParams
	}
}
