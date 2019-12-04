package app

import (
	"encoding/json"
	"fmt"
	"github.com/buyco/keel/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"net/http"
	"strconv"
)

type Response interface {
	GetHttpCode() int
}

// API Error response structure
type ErrorApiResponse struct {
	*Error
	HttpResponse *HttpResponse `json:"http_response"`
}

func (ear *ErrorApiResponse) GetHttpCode() int {
	return ear.HttpResponse.Code
}

// API Success response structure
type SuccessApiResponse struct {
	Data         interface{}   `json:"data"`
	HttpResponse *HttpResponse `json:"http_response"`
}

func (sar *SuccessApiResponse) GetHttpCode() int {
	return sar.HttpResponse.Code
}

type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error handler for Gorilla mux router
func ErrorJSONHandler(message *ErrorApiResponse) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		HandleJSONResponse(w, message, nil)
	}
}

// Use response interface to format response to JSON
func HandleJSONResponse(w http.ResponseWriter, data Response, refFields []string, filters ...string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(data.GetHttpCode())
	output, _ := json.Marshal(data)
	if len(filters) > 0 {
		output = FilterResponseJSONData(output, refFields, filters)
	}
	log.Debugf("Sent response: %s", output)
	w.Write(output)
}

// Helper function to generate error response
func FormatApiError(err *Error, httpCode int) *ErrorApiResponse {
	return &ErrorApiResponse{
		Error: err,
		HttpResponse: &HttpResponse{
			Message: http.StatusText(httpCode),
			Code:    httpCode,
		},
	}
}

// Helper function to generate success response
func FormatApiSuccess(data interface{}, httpCode int) *SuccessApiResponse {
	return &SuccessApiResponse{
		Data: data,
		HttpResponse: &HttpResponse{
			Message: http.StatusText(httpCode),
			Code:    httpCode,
		},
	}
}

func FilterResponseJSONData(data []byte, refFields, fields []string) []byte {
	filteredFields := utils.SliceDiff(refFields, fields)
	for _, field := range filteredFields {
		result := gjson.GetBytes(data, fmt.Sprintf("data.#.%s", field))
		if len(result.Array()) > 0 {
			for index := range result.Array() {
				data = deleteFieldJSONData(data, fmt.Sprintf("data.%s.%s", strconv.Itoa(index), field))
			}
		} else {
			data = deleteFieldJSONData(data, fmt.Sprintf("data.%s", field))
		}
	}
	return data
}

func deleteFieldJSONData(data []byte, path string) []byte {
	data, err := sjson.DeleteBytes(data, path)
	if err != nil {
		log.Error(err)
	}
	return data
}
