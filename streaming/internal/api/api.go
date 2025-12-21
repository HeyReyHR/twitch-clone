package api

import "net/http"

type StreamingHandlerV1 interface {
	OnPublish(w http.ResponseWriter, r *http.Request)
}
