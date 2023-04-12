package handler

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/db/query"
	"tnals5152.com/api-gateway/model"
	util_error "tnals5152.com/api-gateway/utils/error"
)

type ContextHandler struct {
	c             *fiber.Ctx
	requestParams []string
	resource      *model.Resource
	err           error

	queryString map[string][]string
	header      map[string][]string
}

func (c *ContextHandler) SetCtx(ctx *fiber.Ctx) *ContextHandler {
	c.c = ctx

	return c
}

func (c *ContextHandler) SetReqeustParams(params []string) *ContextHandler {
	c.requestParams = append(c.requestParams, params...)

	return c
}

// TODO: params 지우기
func (c *ContextHandler) GetCorrectResource(params []string) (contextHandler *ContextHandler) {
	var err error
	contextHandler = c

	if c.err != nil {
		return
	}

	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&err)

	// 1. filter를 request_path에 맞게 세팅
	filter, sort := SetRequestPathFilterAndSort(params, c.c.Method())
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
	for key := range queryString {
		// queryArgs에 key가 있는지 체크
		// c.c.QueryArgs().PeekMulti(key)로 멀티 값 추출 가능
		if !c.c.Context().QueryArgs().Has(key) {
			return false
		}
	}
	return true
}

// header에서 넘어온 데이터와 c에 있는 데이터가 일치하는지 확인
func (c *ContextHandler) CheckHeader(header model.StringMap) bool {
	headerMap := c.c.GetReqHeaders()
	for key := range header {
		if _, ok := headerMap[key]; !ok {
			return false
		}
	}

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
func (c *ContextHandler) Call() (err error) {
	if c.err != nil {
		return
	}
	// 0. err에 wrap을 사용하여 에러가 발생한 위치를 저장
	defer c.DeferWrap(&c.err)

	// grpc면 grpc 호출
	if strings.ToUpper(c.resource.Method) == constant.GRPC {

	}

	// TODO: carlmjohnson 사용할 예정(header value를 []string으로 세팅할 수 있음)
	// // TODO: timeout setting하기
	// client := &http.Client{}

	// // TODO: c.c.Request().Body() -> io.Reader
	// req, err := http.NewRequest(c.resource.Method, c.resource.Host.Host+":"+c.resource.Host.Port+"/"+c.resource.Path, nil)

	// if err != nil {
	// 	return
	// }

	// req.Header.Add()

	return
}

func (c *ContextHandler) CallGrpc() (err error) {
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

	response, err := client.Connector(ctx, &client_grpc.HttpRequest{
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

	fmt.Println(response.String())
	return

}

func (c *ContextHandler) DeferWrap(err *error) {
	if c.err != nil {
		return
	}
	util_error.DeferWrap(err, 1)

	c.err = *err
}
