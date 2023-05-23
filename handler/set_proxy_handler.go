package handler

import (
	"tnals5152.com/api-gateway/db/query"
	"tnals5152.com/api-gateway/model"
	util_error "tnals5152.com/api-gateway/utils/error"
)

func SetProxyData(resource *model.Resource) (err error) {
	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer util_error.DeferWrap(err, 1)

	err = query.SetProxy(resource)
	return
}
