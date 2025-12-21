package v1

import (
	"net/http"

	"github.com/HeyReyHR/twitch-clone/streaming/internal/model"
	"github.com/go-faster/errors"
)

func (a *api) OnPublish(w http.ResponseWriter, r *http.Request) {
	streamKey := r.FormValue("name")
	err := a.service.ValidateStream(r.Context(), streamKey)
	if err != nil {
		if errors.Is(err, model.ErrInvalidStreamKey) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
}
