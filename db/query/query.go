package query

import (
	"go.mongodb.org/mongo-driver/bson"
	constant "tnals5152.com/api-gateway/const"
)

func And(values ...any) bson.D {
	returnValue := bson.A{}

	for _, value := range values {
		returnValue = append(returnValue, value)
	}
	return bson.D{{Key: constant.AND, Value: returnValue}}
}

func Or(values ...any) bson.D {
	returnValue := bson.A{}

	for _, value := range values {
		returnValue = append(returnValue, value)
	}
	return bson.D{{Key: constant.OR, Value: returnValue}}
}

func Not(value any) bson.E {
	return bson.E{Key: constant.NOT, Value: value}
}

func AddFilter(key string, value any) bson.D {
	return bson.D{{Key: key, Value: value}}
}

func ElemMatch(value any) bson.D {
	return bson.D{{Key: constant.ELEMMATCH, Value: value}}
}

func Exists(key string, exists bool) bson.D {
	return bson.D{{Key: key, Value: bson.D{{Key: constant.EXISTS, Value: exists}}}}
}

func List(values ...any) bson.A {
	returnValue := bson.A{}

	for _, value := range values {
		returnValue = append(returnValue, value)
	}
	return returnValue
}

func AddOption(key string, value any) bson.E {
	return bson.E{Key: key, Value: value}
}
