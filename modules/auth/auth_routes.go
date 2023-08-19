package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edr3x/chi-explore/utils"
)

func Router(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		name := "auth router"

		utils.SendJSONResponse(w, http.StatusOK, utils.SuccessResponse{
			Success: true,
			Payload: name,
		})
	})
}
