package model

type StringMap map[string]string

// Path의 {{}}의 개수와 EndpointPath의 {{}}개수가 같아야 한다.(grpc일 때 제외)
// grpc일 때는 Path의 param을 배열로 전달한다.
type Resource struct {
	ID           string    `json:"id" bson:"_id"`
	RequestPath  *Path     `json:"request_path" bson:"request_path,omitempty"`
	Host         *Host     `json:"host" bson:"host,omitempty"`
	Method       string    `json:"method" bson:"method,omitempty"`
	FunctionName string    `json:"function_name" bson:"function_name,omitempty"`
	QueryString  StringMap `json:"query_string" bson:"query_string,omitempty"`
	Header       StringMap `json:"header" bson:"header,omitempty"`
	FormData     *FormData `json:"form_data" bson:"form_data,omitempty"`
	IsPrivate    bool      `json:"is_private" bson:"is_private,omitempty"`
}

type Path struct {
	Path    string `json:"path" bson:"path,omitempty"`
	IsParam bool   `json:"is_param" bson:"is_param,omitempty"`
	SubPath *Path  `json:"sub_path" bson:"sub_path,omitempty"`
}

type Host struct {
	Name string `json:"name" bson:"name,omitempty"`
	Host string `json:"host" bson:"host,omitempty"`
	Port string `json:"port" bson:"port,omitempty"`
}

type FormData struct {
	Value StringMap `json:"value" bson:"value,omitempty"`
	File  StringMap `json:"file" bson:"file,omitempty"`
}
