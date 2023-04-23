package query

import (
	"errors"

	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/model"
	"tnals5152.com/api-gateway/utils"
)

func CreateHost(host *model.Host) (id string, err error) {
	count, err := GetHostCountByName(host.Name)

	if err != nil {
		return
	}

	// 이미 동일한 이름의 host가 존재할 때
	if count != 0 {
		err = errors.New("이미 일치하는 이름이 존재합니다. 다른 이름을 입력하세요")
		return
	}

	id, err = SetHost(host)

	return
}

func SetHost(host *model.Host) (id string, err error) {
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()

	collection, err := GetHostDB()

	if err != nil {
		return
	}

	result, err := collection.InsertOne(ctx, host)

	if result != nil {
		id = utils.GetObjectId(result.InsertedID)
	}

	return
}

// proxy document add
func SetProxy(resource *model.Resource) (err error) {
	ctx, cancel := utils.GetContext(constant.DBTimeout)
	defer cancel()

	collection, err := GetCollection(constant.PATH_COLLECTION)

	if err != nil {
		return
	}

	_, err = collection.Collection.InsertOne(ctx, resource)

	return
}
