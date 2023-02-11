package handler

import (
	"go.mongodb.org/mongo-driver/bson"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/db/query"
	"tnals5152.com/api-gateway/model"
	"tnals5152.com/api-gateway/utils"
)

func GetEndpointPath(params []string) {
	// 1. filter를 request_path에 맞게 세팅
	filter := SetRequestPathFilter(params)
	var result *model.Resource

	collection, err := query.GetCollection("soominTest")

	if err != nil {
		return
	}

	err = collection.SetResult(&result).SetFilter(filter).GetOne()

	return
}

// request_path에 맞게 쿼리 세팅
func SetRequestPathFilter(params []string) any {
	filterValue := bson.A{}
	key := constant.RequestPath

	for _, param := range params {
		filterValue = append(filterValue,
			bson.D{
				{
					Key: constant.OR,
					Value: bson.A{
						bson.D{
							{
								Key: utils.JoinWithDot(
									key,
									constant.Path,
								),
								Value: param,
							},
						},
						bson.D{
							{
								Key: utils.JoinWithDot(
									key,
									constant.IsParam,
								),
								Value: true,
							},
						},
					},
				},
			},
		)

		key = utils.JoinWithDot(key, constant.SubPath)
	}

	// 마지막엔 sub_path가 존재하지 않아야 함 -> 안 그러면 앞에만 일치하는 값도 나옴
	filterValue = append(filterValue,
		bson.D{
			{
				Key: key,
				Value: bson.D{
					{
						Key:   constant.EXISTS,
						Value: false,
					},
				},
			},
		},
	)

	filter := bson.D{{
		Key:   constant.AND,
		Value: filterValue,
	}}

	return filter
}
