package constant

import "time"

const (
	DBPort     = "mongodb.port"
	DBUsername = "mongodb.username"
	DBPassword = "mongodb.password"
	DBHost     = "mongodb.hostname"
	DBDatabase = "mongodb.database"
	DBTimeout  = "mongodb.timeout"

	ServerAppName      = "server.app_name"
	ServerPort         = "server.port"
	ServerReadTimeout  = "server.read_timeout"
	ServerWriteTimeout = "server.write_timeout"

	PATH_COLLECTION     = "path"
	ENDPOINT_COLLECTION = "endpoint"

	GatewayGrpcPort = "gateway.grpc.port"

	GrpcTimeout = "grpc.timeout"

	Host = "host"

	PLUS  = "+"
	SLASH = "/"

	Path        = "path"
	IsParam     = "is_param"
	RequestPath = "request_path"
	SubPath     = "sub_path"

	AND    = "$and"
	OR     = "$or"
	EXISTS = "$exists"
)

// 기본 타임아웃 세팅
var ContextTimeoutMap map[string]time.Duration = map[string]time.Duration{
	DBTimeout:          10,
	ServerReadTimeout:  10,
	ServerWriteTimeout: 10,
}
