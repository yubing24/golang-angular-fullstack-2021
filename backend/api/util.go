package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func respond(ctx context.Context, w http.ResponseWriter, status int, err error, data interface{}) {
	result := struct{
		Status  int `json:"status"`
		Message string `json:"message"`
		Data interface{} `json:"data"`
	} {
		Status: status,
		Data: data,
	}

	if err != nil {
		result.Message = err.Error()
	}

	output, _ := json.Marshal(result)
	w.WriteHeader(status)
	w.Write(output)
}