package constant

import "time"

const (
	DBPort     = "mongodb.port"
	DBUsername = "mongodb.username"
	DBPassword = "mongodb.password"
	DBHost     = "mongodb.hostname"
	DBDatabase = "mongodb.database"
	DBTimeout  = "mongodb.timeout"
)

// 기본 타임아웃 세팅
var ContextTimeoutMap map[string]time.Duration = map[string]time.Duration{
	DBTimeout: 10,
}
