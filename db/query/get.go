package query

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/db"
	"tnals5152.com/api-gateway/model"
	"tnals5152.com/api-gateway/utils"
	util_error "tnals5152.com/api-gateway/utils/error"
)

type Collection struct {
	Collection *mongo.Collection
	result     any
	filter     any
	sort       any
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

	findOptions := make([]*options.FindOptions, 0)

	if !utils.IsNil(c.GetSort()) {
		sort, ok := c.GetSort().(*options.FindOptions)

		if ok {
			findOptions = append(findOptions, sort)
		}
	}

	cursor, err := c.Collection.Find(
		ctx,
		c.filter,
		findOptions...,
	)

	if err != nil {
		return
	}

	err = cursor.All(ctx, c.result)
	return
}

func (c *Collection) Exists() (exists bool, err error) {
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()

	countOptions := make([]*options.CountOptions, 0)

	if !utils.IsNil(c.GetSort()) {
		sort, ok := c.GetSort().(*options.CountOptions)

		if ok {
			countOptions = append(countOptions, sort)
		}
	}

	count, err := c.Collection.CountDocuments(
		ctx,
		c.filter,
		countOptions...,
	)

	if err != nil {
		return
	}

	if count != 0 {
		exists = true
	}
	return
}

func (c *Collection) GetOne() (err error) {
	defer util_error.DeferWrap(&err)
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

func (c *Collection) GetSort() any {
	return c.sort
}

func (c *Collection) SetResult(result any) *Collection {
	c.result = result
	return c
}

func (c *Collection) SetFilter(filter any) *Collection {
	c.filter = filter
	return c
}

func (c *Collection) SetSort(sort any) *Collection {
	c.sort = sort
	return c
}
