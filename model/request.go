package model

import (
	"errors"
	"strings"

	"tnals5152.com/api-gateway/utils"
)

// type RequestPath struct {
// 	RequestPath *Path `json:"request_path,omitempty" bson:"request_path"`
// }

type RequestParam struct {
	RequestParam map[string]string `json:"request_params"`
}
type RequestResource struct {
	RequestPath   string    `json:"request_path"`
	RequestMethod string    `json:"request_method"`
	Host          *Host     `json:"host"` // name으로 지정할지 고민 중
	Method        string    `json:"method"`
	FunctionName  string    `json:"function_name"`
	QueryString   StringMap `json:"query_string,omitempty"`
	Header        StringMap `json:"header,omitempty"`
	FormData      *FormData `json:"form_data,omitempty"`
	Body          StringMap `json:"body,omitempty"`
	IsPrivate     bool      `json:"is_private" `
	CorsCheckApi  bool      `json:"cors_check_api" `
	Path          string    `json:"path"`
}

func (resource *RequestResource) Validate() (err error) {
	if resource.RequestPath == "" {
		err = errors.New("request path is required")
		return
	}
	if resource.RequestMethod == "" {
		err = errors.New("request method is required")
		return
	}
	if resource.Host == nil {
		err = errors.New("host is required")
		return
	}
	if resource.Method == "" {
		err = errors.New("method is required")
		return
	}
	if resource.FunctionName == "" {
		err = errors.New("function_name is required")
		return
	}
	if resource.Path == "" {
		err = errors.New("path is required")
		return
	}

	return
}

func (resource *RequestResource) ToResource() (result *Resource, err error) {

	result = &Resource{
		// RequestPath:  result.RequestPath,
		RequestMethod: resource.RequestMethod,
		Host:          resource.Host,
		Method:        resource.Method,
		FunctionName:  resource.FunctionName,
		QueryString:   resource.QueryString,
		Header:        resource.Header,
		FormData:      resource.FormData,
		Body:          resource.Body,
		IsPrivate:     resource.IsPrivate,
		CorsCheckApi:  resource.CorsCheckApi,
		Path:          resource.Path,
	}

	paths := strings.Split(resource.RequestPath, "/")

	path := &Path{}

	prevPath := path

	for i, subPath := range paths {
		if subPath == "" { // /로 시작되어 전달되는 경우 건너뛰기
			continue
		}
		isParam := utils.RePathParam.MatchString(subPath)

		prevPath.Path = subPath
		prevPath.IsParam = isParam

		if i != len(paths)-1 {
			prevPath.SubPath = &Path{}
			prevPath = prevPath.SubPath
		}

	}
	result.RequestPath = path

	return
}
