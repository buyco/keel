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

// Response if an interface to get HTTP code
type Response interface {
	GetHTTPCode() int
}

// ErrorAPIResponse is an Error response struct
type ErrorAPIResponse struct {
	*Error
	HTTPResponse *HTTPResponse `json:"http_response"`
}

// GetHTTPCode return HTTP code
func (ear *ErrorAPIResponse) GetHTTPCode() int {
	return ear.HTTPResponse.Code
}

// SuccessAPIResponse is a Success response struct
type SuccessAPIResponse struct {
	Data         interface{}   `json:"data"`
	HTTPResponse *HTTPResponse `json:"http_response"`
}

// GetHTTPCode return HTTP code
func (sar *SuccessAPIResponse) GetHTTPCode() int {
	return sar.HTTPResponse.Code
}

// HTTPResponse is a light HTTP response struct
type HTTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorJSONHandler handle error for Gorilla mux router
func ErrorJSONHandler(message *ErrorAPIResponse) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		HandleJSONResponse(w, message, nil)
	}
}

// HandleJSONResponse handle JSON API response
func HandleJSONResponse(w http.ResponseWriter, data Response, refFields []string, filters ...string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(data.GetHTTPCode())
	output, _ := json.Marshal(data)
	if len(filters) > 0 {
		output = FilterResponseJSONData(output, refFields, filters)
	}
	log.Debugf("Sent response: %s", output)
	w.Write(output)
}

// FormatAPIError is a helper function to generate error response
func FormatAPIError(err *Error, httpCode int) *ErrorAPIResponse {
	return &ErrorAPIResponse{
		Error: err,
		HTTPResponse: &HTTPResponse{
			Message: http.StatusText(httpCode),
			Code:    httpCode,
		},
	}
}

// FormatAPISuccess is a helper function to generate success response
func FormatAPISuccess(data interface{}, httpCode int) *SuccessAPIResponse {
	return &SuccessAPIResponse{
		Data: data,
		HTTPResponse: &HTTPResponse{
			Message: http.StatusText(httpCode),
			Code:    httpCode,
		},
	}
}

// FilterResponseJSONData filters fields from JSON
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
