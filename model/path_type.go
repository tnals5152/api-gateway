package model

type PathInfo struct {
	ID     string `json:"id" bson:"_id"`
	Level  int    `json:"level,omitempty" bson:"level"`   // path의 레벨로 1부터 시작해서 1, 2, 3, ... 값을 가진다.
	Path   string `json:"path,omitempty" bson:"path"`     // url의 각각의 path를 저장하며 path param일 경우 {{param 이름}} 으로 저장한다.
	Method string `json:"method,omitempty" bson:"method"` // GET, POST, PUT, DELETE, HEAD, PATCH, OPTION, ANY 값들이 들어가고, 기본 대문자, default 값은 ANY다.
	// PrePathIds []string `json:"pre_path_ids,omitempty" bson:"pre_path_ids"` // 바로 이전 path의 id를 배열로 저장한다. -> 보류
	SubPathIds []string `json:"sub_path_ids,omitempty" bson:"sub_path_ids"` // 바로 다음 path의 id를 배열로 저장한다.
	IsParam    bool     `json:"is_param,omitempty" bson:"is_param"`         // 현재 row가 path_param인지 아니면 그냥 path인지를 나타낸다. path_param이면 true, 아니면 false이며, default = false
	IsLastPath bool     `json:"is_last_path,omitempty" bson:"is_last_path"` // 현재 path가 마지막 path인지를 나타내는 것으로 true, false값이 저장되고 default = false
}
