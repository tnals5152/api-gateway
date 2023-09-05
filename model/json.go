package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"tnals5152.com/api-gateway/utils"
	util_error "tnals5152.com/api-gateway/utils/error"
)

type ErrorData struct {
	InputData  string `json:"input_tata"`
	Err        error  `json:"-"`
	TargetType string `json:"target_type"`
}

func (e *ErrorData) UnmarshalJSON(data []byte) (err error) {
	defer util_error.DeferWrap(&err)
	type E ErrorData
	errorDataInner := &struct {
		*E
		ErrorString *string `json:"error"`
	}{
		E: (*E)(e),
	}
	err = json.Unmarshal(data, errorDataInner)

	if err != nil {
		return
	}

	if errorDataInner.ErrorString != nil {
		e.Err = errors.New(*errorDataInner.ErrorString)
	}

	return
}

func (e ErrorData) MarshalJSON() ([]byte, error) {
	var err error
	defer util_error.DeferWrap(&err)
	type E ErrorData
	errorDataInner := &struct {
		*E
		ErrorString *string `json:"error"`
	}{
		E: (*E)(&e),
	}

	if e.Err != nil {
		errString := e.Err.Error()
		errorDataInner.ErrorString = &errString
	}

	data, err := json.Marshal(errorDataInner)

	return data, err
}

func (e *ErrorData) IsError() bool {
	return e.Err != nil
}

func (e *ErrorData) Error() string {

	return utils.Join("error: ", e.Err.Error(), ", inputData: ", e.InputData, ", TargetType: ", e.TargetType)
}

func Unmarshal(data []byte, v any) (errorData error) {

	err := json.Unmarshal(data, v)

	if err == nil {
		return
	}

	TargetType := reflect.TypeOf(v)

	errorData = &ErrorData{
		InputData:  string(data),
		Err:        err,
		TargetType: TargetType.String(),
	}

	return
}
