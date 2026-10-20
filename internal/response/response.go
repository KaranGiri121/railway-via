package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func SendByteWriteJson(w http.ResponseWriter, status int, data string) error {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
	return nil
}

func GenralError(err error) Response {
	return Response{
		Status: "Error",
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errorMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMsg = append(errorMsg, fmt.Sprintf("%s field is required", err.Field()))
		default:
			errorMsg = append(errorMsg, fmt.Sprintf("field %s is invalid", err.Field()))

		}
	}
	return Response{
		Status: "Error",
		Error:  strings.Join(errorMsg, ", "),
	}
}
