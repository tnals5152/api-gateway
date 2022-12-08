package utils

import (
	"context"
	"time"

	"github.com/spf13/viper"
	constant "tnals5152.com/api-gateway/const"
)

func GetContext(key string) (context.Context, context.CancelFunc) {
	var timeout time.Duration
	if viper.InConfig(key) {
		timeout = viper.GetDuration(key)
	} else {
		value, ok := constant.ContextTimeoutMap[key]

		if ok {
			timeout = value
		} else {
			timeout = 10
		}
	}

	return context.WithTimeout(context.Background(), timeout*time.Second)
}
