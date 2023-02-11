package query

import (
	"go.mongodb.org/mongo-driver/mongo"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/db"
	"tnals5152.com/api-gateway/model"
	"tnals5152.com/api-gateway/utils"
)

type Collection struct {
	Collection *mongo.Collection
	result     any
	filter     any
}

func GetHostCountByName(name string) (count int64, err error) {
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()

	collection, err := GetHostDB()

	if err != nil {
		return
	}

	count, err = collection.
		CountDocuments(
			ctx,
			model.Host{
				Name: name,
			},
		)

	return
}

func GetHostDB() (collection *mongo.Collection, err error) {
	collection = db.Mongo.DB.Collection(constant.Host)

	if collection != nil {
		return
	}
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()
	err = db.Mongo.DB.CreateCollection(ctx, constant.Host)
	if err != nil {
		return
	}
	collection = db.Mongo.DB.Collection(constant.Host)
	return
}

func GetCollection(collectionName string) (collection *Collection, err error) {
	coll := db.Mongo.DB.Collection(collectionName)

	if coll != nil {
		collection = &Collection{
			Collection: coll,
		}
		return
	}
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()
	err = db.Mongo.DB.CreateCollection(ctx, collectionName)
	if err != nil {
		return
	}
	coll = db.Mongo.DB.Collection(collectionName)
	collection = &Collection{
		Collection: coll,
	}
	return
}

func (c *Collection) GetAll() (err error) {
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()
	cursor, err := c.Collection.Find(
		ctx,
		c.filter,
	)

	if err != nil {
		return
	}

	err = cursor.All(ctx, c.result)
	return
}

func (c *Collection) GetOne() (err error) {
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()
	singleResult := c.Collection.FindOne(
		ctx,
		c.filter,
	)

	err = singleResult.Decode(c.result)
	return
}

func (c *Collection) GetResult() any {
	return c.result
}

func (c *Collection) GetFilter() any {
	return c.filter
}

func (c *Collection) SetResult(result any) *Collection {
	c.result = result
	return c
}

func (c *Collection) SetFilter(filter any) *Collection {
	c.filter = filter
	return c
}
