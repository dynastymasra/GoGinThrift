package util

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

type Jsend struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type JsendObject struct {
	Jsend
	Object map[string]interface{} `json:"data"`
}

type JsendList struct {
	Jsend
	ObjectList []interface{} `json:"data"`
}

type JsendError struct {
	Status  string `json:"status"`
	Message error  `json:"message"`
}

func FailResponse(msg string) Jsend {
	return Jsend{"failed", msg}
}

func SuccessResponse(msg string) Jsend {
	return Jsend{Status: "success", Message: msg}
}

func ObjectResponse(status, message string, data map[string]interface{}) JsendObject {

	response := JsendObject{}
	response.Status = status
	response.Message = message
	response.Object = data

	return response
}

func ListResponse(status, message string, data []interface{}) JsendList {
	response := JsendList{}
	response.Status = status
	response.Message = message
	response.ObjectList = data

	return response
}
