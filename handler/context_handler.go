package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/db/query"
	grpc "tnals5152.com/api-gateway/grpc"
	client_grpc "tnals5152.com/api-gateway/grpc/client"
	"tnals5152.com/api-gateway/model"
	"tnals5152.com/api-gateway/utils"
	util_error "tnals5152.com/api-gateway/utils/error"
)

type ContextHandler struct {
	c             *fiber.Ctx
	requestParams []string
	resource      *model.Resource
	err           error

	queryString map[string][]string
	header      map[string][]string
	body        map[string]any
	Response    any
}

func (c *ContextHandler) SetCtx(ctx *fiber.Ctx) *ContextHandler {
	c.c = ctx

	return c
}

func (c *ContextHandler) SetReqeustParams(params []string) *ContextHandler {
	c.requestParams = append(c.requestParams, params...)

	return c
}

// 일치하는 resource DB에서 조회
func (c *ContextHandler) GetCorrectResource() (contextHandler *ContextHandler) {
	var err error
	contextHandler = c

	if c.err != nil {
		return
	}

	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&err)

	// 1. filter를 request_path에 맞게 세팅
	filter, sort := SetRequestPathFilterAndSort(c.requestParams, c.c.Method(), c.c.GetReqHeaders())
	var resources []*model.Resource

	collection, err := query.GetCollection(constant.PATH_COLLECTION)

	if err != nil {
		return
	}

	// 2. filter 적용하여 결과값 받기
	err = collection.
		SetResult(&resources).
		SetFilter(filter).
		SetSort(sort).
		GetAll()

	if err != nil {
		return
	}

	if len(resources) == 0 {
		err = errors.New("일치하는 경로가 존재하지 않습니다.")
		return
	}

	// 3. queryString, header, form_data, is_private 가 맞는지 확인한다.
	resource, err := c.getCorrectResource(resources)

	if err != nil {
		return
	}

	contextHandler.resource = resource

	return
}

func (c *ContextHandler) getCorrectResource(resources []*model.Resource) (correctResource *model.Resource, err error) {
	if c.err != nil {
		err = c.err
		return
	}

	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&err)

	for _, resource := range resources {
		// 3-1. 필수 querString이 존재하는지 체크, 없으면 다음 resource로 넘어감
		if !c.CheckQueryString(resource.QueryString) {
			continue
		}

		// 3-2. 필수 header가 존재하는지 체크, 없으면 다음 resource로 넘어감
		if !c.CheckHeader(resource.Header) {
			continue
		}

		// 3-3. 필수 formData가 존재하는지 체크. 없으면 다음 resource로 넘어감
		if !c.CheckFormData(resource.FormData) {
			continue
		}

		// 3-4. 필수 body key가 존재하는지 체크, 없으면 다음 resource로 넘어감
		if check, bodyErr := c.CheckBody(resource.Body); bodyErr != nil {
			err = bodyErr
			return
		} else if !check {
			continue
		}

		var exists bool
		// 3-4. isPrivate 체크(해당 resource가 isPrivate 일 시, approve collection에 저장된 host에서만 호출 가능)
		if exists, err = c.CheckIsPrivate(resource.IsPrivate); err != nil {
			return
		} else if !exists {
			continue
		}

		correctResource = resource
		return
	}

	err = errors.New("일치하는 경로가 존재하지 않습니다.")

	return
}

// queryString에서 넘어온 데이터와 c에 있는 데이터가 일치하는지 확인
func (c *ContextHandler) CheckQueryString(queryString model.StringMap) bool {

	var queryMap map[string][]string = make(map[string][]string)

	c.c.Context().QueryArgs().VisitAll(func(key, value []byte) {
		queryMap[string(key)] = append(queryMap[string(key)], string(value))
	})

	for key := range queryString {
		// queryArgs에 key가 있는지 체크
		// c.c.QueryArgs().PeekMulti(key)로 멀티 값 추출 가능
		if _, ok := queryMap[key]; !ok {
			return false
		}
	}

	return true
}

// header에서 넘어온 데이터와 c에 있는 데이터가 일치하는지 확인
func (c *ContextHandler) CheckHeader(header model.StringMap) bool {
	var headerMap map[string][]string = make(map[string][]string)

	c.c.Request().Header.VisitAll(func(key, value []byte) {
		headerMap[string(key)] = append(headerMap[string(key)], string(value))
	})

	for key := range header {
		if _, ok := headerMap[key]; !ok {
			return false
		}
	}

	c.header = headerMap

	return true
}

// formData에서 넘어온 데이터와 c에 있는 데이터가 일치하는지 확인
func (c *ContextHandler) CheckFormData(formData *model.FormData) bool {

	if formData == nil {
		return true
	}
	multipartForm, err := c.c.MultipartForm()

	if err != nil {
		return false
	}

	for key := range formData.File {
		file := multipartForm.File[key]

		if file == nil || len(file) == 0 {
			return false
		}
	}

	for key := range formData.Value {
		value := multipartForm.Value[key]

		if value == nil || len(value) == 0 {
			return false
		}
	}

	return true
}

// body에서 넘어온 key가 일치하는지 확인
func (c *ContextHandler) CheckBody(bodyString model.StringMap) (check bool, err error) {

	if len(bodyString) == 0 {
		return true, nil
	}
	c.body = map[string]any{}

	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&err)
	// api로 넘어온 body
	var body map[string]string
	err = c.c.BodyParser(&body)

	if err != nil {
		return
	}

	// db에 저장된 body
	for key, value := range bodyString {
		if bodyValue, ok := body[key]; !ok {
			return false, errors.New(key + "가 body에 존재하지 않습니다.")
		} else {
			c.body[value] = bodyValue
		}
	}

	return true, nil
}

func (c *ContextHandler) CheckIsPrivate(isPrivate bool) (isOk bool, err error) {
	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&err)

	// private api가 아니면 사용 가능
	if !isPrivate {
		isOk = true
		return
	}

	var approves []*model.Approve

	collection, err := query.GetCollection(constant.APPROVE_COLLECTION)

	if err != nil {
		return
	}

	exists, err := collection.
		SetResult(&approves).
		SetFilter(
			bson.D{{Key: constant.IP, Value: c.c.IP()}},
		).Exists()

	if err != nil {
		return
	}

	isOk = exists
	return
}

// 일치하는 resource를 찾은 경우 해당 api 혹은 gRPC를 호출하는 함수
// header, query, params 는 체크 완료
// 여기선
func (c *ContextHandler) Call() (response any, err error) {
	if c.err != nil {
		err = c.err
		return
	}
	response = &c.Response
	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&c.err)

	// 1. form_data가 있을 시 validate 체크

	// 2. grpc면 grpc 호출
	if strings.ToUpper(c.resource.Method) == constant.GRPC {
		response, err = c.CallGrpc()
		return
	}

	// 3. http 호출
	request := requests.
		URL(c.resource.Host.GetUrl()).
		ToHeaders(c.header).
		BodyJSON(c.body).
		Method(c.resource.Method).
		Path("/" + strings.Join(c.requestParams, "/"))

	for key, value := range c.queryString {
		request.Param(key, value...)
	}

	ctx, cancel := utils.GetContext(viper.GetString(constant.HttpTimeout))
	defer cancel()

	err = request.ToJSON(&c.Response).Fetch(ctx)

	return
}

func (c *ContextHandler) CallGrpc() (response any, err error) {
	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&c.err)

	conn, cancel, err := grpc.ConnectGrpcClient(c.resource.Host.Host, c.resource.Host.Port)

	if err != nil {
		return
	}

	defer cancel()

	client := client_grpc.NewGprcInitClient(conn)
	ctx, cancel := utils.GetContext(viper.GetString(constant.GrpcTimeout))
	defer cancel()

	grpcResponse, err := client.Connector(ctx, &client_grpc.HttpRequest{
		Method: c.c.Method(),
		Headers: func() []*client_grpc.Header {

			var headers []*client_grpc.Header

			for mapKey, mapValue := range c.header {
				headers = append(headers, &client_grpc.Header{
					Key:   mapKey,
					Value: mapValue,
				})
			}

			return headers
		}(),
		Params: c.requestParams,
		Queries: func() []*client_grpc.Query {

			var queries []*client_grpc.Query

			for mapKey, mapValue := range c.queryString {
				queries = append(queries, &client_grpc.Query{
					Key:   mapKey,
					Value: mapValue,
				})
			}
			return queries
		}(),
		Body: c.c.Body(),
	})

	if err != nil {
		return
	}

	c.Response = grpcResponse
	response = grpcResponse
	fmt.Println(grpcResponse.String())
	return

}

func (c *ContextHandler) DeferWrap(err *error) {
	if c.err != nil {
		return
	}
	util_error.DeferWrap(err, 1)

	c.err = *err
}
