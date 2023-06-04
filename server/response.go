package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ServerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ServerResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func SendResponse(success bool, message string, code int, statusCode int, res http.ResponseWriter) {
	response := &ServerResponse{
		Success: success,
		Message: message,
		Code:    code}

	responseJSON, _ := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	fmt.Fprint(res, string(responseJSON))
}

func SendResponseData(success bool, message string, data interface{}, code int, statusCode int, res http.ResponseWriter) {
	response := &ServerResponseData{
		Success: success,
		Message: message,
		Code:    code,
		Data:    data,
	}

	responseJSON, _ := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	fmt.Fprint(res, string(responseJSON))
}

func CheckErr(err error, res http.ResponseWriter) {
	if err != nil {
		log.Printf("ERROR: %s", err)
		SendResponse(false, "Internal Server Error", 500, http.StatusInternalServerError, res)
	}
}