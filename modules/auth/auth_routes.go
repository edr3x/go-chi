package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edr3x/chi-explore/utils"
)

func Router(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Get query params
		queryparams := r.URL.Query()
		name := queryparams.Get("name") // reads from:  http://localhost:5050/api/v1/auth?name=obiwan

		val, err := testService(name)
		if err != nil {
			panic(err)
		}

		utils.SendJSONResponse(w, http.StatusOK, utils.SuccessResponse{
			Success: true,
			Payload: val,
		})
	})
}
