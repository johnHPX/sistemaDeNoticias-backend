package util

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// =====================================================================
// Error response
// =====================================================================
type errorResponse struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"message"`
	MID    string `json:"_mid"`
}

func (e *errorResponse) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, mid: %s",
		e.Code, e.Msg, e.MID,
	)
}

func ErrorEncoder() httptransport.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		// check if http error type
		rErr, ok := err.(*errorResponse)
		if !ok {
			rErr = &errorResponse{
				Status: 500,
				Code:   0,
				Msg:    err.Error(),
				MID:    "na",
			}
		}
		// write status
		w.WriteHeader(rErr.Status)
		// encode and write error response
		encoder := json.NewEncoder(w)
		encoder.Encode(rErr)
	}
}

// =====================================================================
// Error response creator
// =====================================================================
func CreateHttpErrorResponse(status, code int, err error, mid string) *errorResponse {
	return &errorResponse{
		Status: status,
		Code:   code,
		Msg:    err.Error(),
		MID:    mid,
	}
}

// ==============================================================
// Response json
// ==============================================================

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
