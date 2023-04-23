package handler

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/db/query"
	"tnals5152.com/api-gateway/utils"
)

// request_path에 맞게 쿼리 세팅
func SetRequestPathFilterAndSort(params []string, requestMethod string, header map[string]string) (
	filter any, sort any) {

	sortOption := bson.D{}
	filterValue := bson.A{}
	filterValue = append(filterValue, query.AddFilter(constant.RequestMethod, requestMethod))
	key := constant.RequestPath

	for _, param := range params {
		filterValue = append(filterValue,
			query.Or(
				query.AddFilter(
					utils.JoinWithDot(key, constant.Path),
					param,
				),
				query.AddFilter(
					utils.JoinWithDot(key, constant.Path),
					true,
				),
			),
		)

		sortOption = append(sortOption,
			query.AddOption(
				utils.JoinWithDot(key, constant.IsParam),
				1,
			),
		)

		key = utils.JoinWithDot(key, constant.SubPath)
	}

	// 마지막엔 sub_path가 존재하지 않아야 함 -> 안 그러면 앞에만 일치하는 값도 나옴
	filterValue = append(filterValue, query.Exists(key, false))

	// TODO: 필수 header가 아닌 경우도 있을 수 있으니 체크 보류
	// for key := range header {

	// 	filterValue = append(filterValue, query.Exists(key, true))

	// }

	filter = query.And(filterValue...)

	sort = options.Find().SetSort(sortOption)

	return
}
