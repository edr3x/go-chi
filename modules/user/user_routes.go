package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edr3x/chi-explore/middlewares"
	"github.com/edr3x/chi-explore/utils"
)

func Router(r chi.Router) {
	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		contextValue := r.Context().Value(middlewares.UserDataKey).(string)

		if name == "err" {
			panic(utils.NewError(http.StatusBadRequest, "Panic was triggred"))
		}

		utils.SendJSONResponse(w, http.StatusOK, utils.SuccessResponse{
			Success: true,
			Payload: name + " " + contextValue,
		})
	})
}
