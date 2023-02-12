package handler

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/utils"
)

// request_path에 맞게 쿼리 세팅
func SetRequestPathFilterAndSort(params []string, requestMethod string) (filter any, sort any) {
	sortOption := bson.D{}
	filterValue := bson.A{}
	key := constant.RequestPath

	filterValue = append(filterValue, bson.D{
		{
			Key:   constant.RequestMethod,
			Value: requestMethod,
		},
	})

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

		// is_param이 false인 데이터 먼저 반환
		sortOption = append(sortOption,
			primitive.E{
				Key: utils.JoinWithDot(
					key,
					constant.IsParam,
				),
				Value: 1,
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

	filter = bson.D{{
		Key:   constant.AND,
		Value: filterValue,
	}}

	sort = options.Find().SetSort(sortOption)

	return
}
