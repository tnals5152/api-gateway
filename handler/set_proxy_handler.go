package handler

import (
	"errors"

	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/db/query"
	"tnals5152.com/api-gateway/model"
	util_error "tnals5152.com/api-gateway/utils/error"
)

func SetProxyData(requestResource *model.RequestResource) (err error) {
	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer util_error.DeferWrap(&err, 1)

	// requestResource to resource
	resource, err := requestResource.ToResource()

	if err != nil {
		return
	}

	collection, err := query.GetCollection(constant.PATH_COLLECTION)

	if err != nil {
		return
	}
	// 1. 중복되는 api가 있는지 체크
	// 중복 여부 체크하는 filter 생성

	// 1. host 가 같아야 하며, method가 같아야 하고, path가 같을 때
	filter := CheckDuplicateHostFilter(resource)

	var hostResult []any

	exists, err := collection.SetFilter(filter).SetResult(&hostResult).Exists()

	if err != nil {
		return
	}

	if exists {
		err = errors.New("이미 해당 path로 api가 존재합니다. 기존 api를 수정하세요.")
		return
	}

	// 2. (request_path && request_method) || function_name 이 같을 때 중복
	filter = CheckDuplicateRequestFilter(resource)

	var requestResult []any

	exists, err = collection.SetFilter(filter).SetResult(&requestResult).Exists()

	if err != nil {
		return
	}

	if exists {
		err = errors.New("이미 해당 request 혹은 functionName으로 api가 존재합니다. request 요청 정보를 변경하세요.")
		return
	}
	// request_method가 나오지 않음!!!
	// resource 생성
	err = query.SetProxy(resource)
	return
}
