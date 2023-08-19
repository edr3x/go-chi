package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/edr3x/chi-explore/middlewares"
	"github.com/edr3x/chi-explore/modules/auth"
	"github.com/edr3x/chi-explore/modules/user"
)

func MainRouter(r chi.Router) {
	r.Route("/auth", auth.Router)

	r.Route("/user", func(usr chi.Router) {
		usr.Use(middlewares.RequireAuth) // injecting middleware
		user.Router(usr)
	})
}
