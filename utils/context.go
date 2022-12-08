package utils

import (
	"context"
	"time"

	"github.com/spf13/viper"
	constant "tnals5152.com/api-gateway/const"
)

func GetContext(key string) (context.Context, context.CancelFunc) {
	var timeout time.Duration
	timeout = GetTimeout(key)

	return context.WithTimeout(context.Background(), timeout*time.Second)
}

func GetTimeout(key string) time.Duration {
	if viper.InConfig(key) {
		return viper.GetDuration(key)
	}
	value, ok := constant.ContextTimeoutMap[key]

	if ok {
		return value
	}
	return 10

}
