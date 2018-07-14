package api

import (
	"net/http"
	"encoding/json"
)

// any response body (or nil) can be passed as type-casted to ApiResponse
type ApiResponse interface{}

// an API handler processes a http request and returns either a body, or error
type ApiHandler func(r *http.Request) (ApiResponse, Error)

type Handler struct {
	apiHandler ApiHandler
}

func NewHandler(apiHandler ApiHandler) *Handler {
	return &Handler {
		apiHandler: apiHandler,
	}
}

// boiler plate wrapper for API Handler implementations
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if resp, err := h.apiHandler(r); err != nil {
		w.WriteHeader(err.ErrCode())
		json.NewEncoder(w).Encode(err)
	} else {
		switch method {
			case "GET":
				w.WriteHeader(200)
				if resp != nil {
					json.NewEncoder(w).Encode(resp)
				}
				break
			case "POST":
				if resp != nil {
					w.WriteHeader(201)
					json.NewEncoder(w).Encode(resp)
				} else {
					w.WriteHeader(202)
				}
				break
			case "PUT":
				w.WriteHeader(202)
				if resp != nil {
					json.NewEncoder(w).Encode(resp)
				}
				break
			case "DELETE":
				if resp != nil {
					w.WriteHeader(200)
					json.NewEncoder(w).Encode(resp)					
				} else {
					w.WriteHeader(204)					
				}
				break
			default:
				w.WriteHeader(405)
				json.NewEncoder(w).Encode(ApiError(ERR_METHOD_NOT_ALLOWED, "method not allowed"))
		}
	}
}