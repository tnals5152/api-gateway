package utils

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var RePathParam *regexp.Regexp

func init() {
	re, err := regexp.Compile(`{{.*}}`) // main/{{num}}/post

	if err != nil {
		panic(err)
	}

	RePathParam = re
}

func CreateUUID() string {
	return uuid.New().String()
}

func GetObjectId(objectId any) (oid string) {
	if IsNil(objectId) {
		return
	}

	if oid, ok := objectId.(primitive.ObjectID); ok {
		return oid.Hex()
	}
	return
}

func JoinWithDot(str ...string) string {
	return strings.Join(str, ".")
}

func IsNil(i interface{}) bool {

	if i == nil {
		return true
	}

	reflectTypeOf := reflect.TypeOf(i)
	kind := reflectTypeOf.Kind()
	reflectValueOf := reflect.ValueOf(i)

	// 기본 타입으로 되면 아래로
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Func, reflect.Chan, reflect.Slice, reflect.Interface, reflect.UnsafePointer:
		return reflectValueOf.IsNil()
	}

	reflectType := reflectValueOf.Type().String()

	switch reflectType {
	case "primitive.ObjectID":
		return i.(primitive.ObjectID).IsZero()
	}

	return false
}
