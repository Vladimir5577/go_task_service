package helper

import (
	"encoding/json"
	"net/http"
	"task_service/internal/model"
)

type JsonResponseType struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func JsonResponse(w http.ResponseWriter, data any, statusCode int) {
	resp := JsonResponseType{}
	if statusCode == http.StatusOK ||
		statusCode == http.StatusCreated {
		resp.Success = true
	} else {
		resp.Success = false
	}

	resp.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func ParseErrorResponse(w http.ResponseWriter, data any, statusCode int) {
	if serviceError, ok := data.(*model.ServiceError); ok {
		JsonResponse(w, serviceError.Message, serviceError.StatusCode)
	} else {
		err, _ := data.(error)
		JsonResponse(w, err.Error(), statusCode)
	}
}
