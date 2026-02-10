package helper

import (
	"encoding/json"
	"net/http"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
)

func ReadFromRequestBody(req *http.Request, result any) error {
	return json.NewDecoder(req.Body).Decode(result)
}

func WriteToResponseBody(writer http.ResponseWriter, data any) error {
	writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(data)
}

func ResponseJsonSuccess(writer http.ResponseWriter, data any) error {
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(webResponse)
}

func ResponseJSONError(writer http.ResponseWriter, statusCode int, messageStatus string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	webResponse := web.WebResponse{
		Code:   statusCode,
		Status: messageStatus,
		Data:   nil,
	}
	WriteToResponseBody(writer, webResponse)
}