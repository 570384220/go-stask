package main

const (
	requestParamError = 10001
	operateSucess     = 200
)

/*
首字母必须大写，才能在包外访问或者序列化后才能有对应的字段，小写只能包内访问
*/
type ResObj struct {
	Code    int
	Message string
	Data    interface{}
}

func getRequestStatusTxt(status int) (int, string) {

	switch status {
	case requestParamError:

		return requestParamError, "Request Param Error"
	case operateSucess:
		return operateSucess, "operator Success"
	default:
		return -1, ""
	}
}

func ResSuccess() any {
	code, msg := getRequestStatusTxt(operateSucess)
	return ResObj{Code: code, Message: msg, Data: nil}
}
