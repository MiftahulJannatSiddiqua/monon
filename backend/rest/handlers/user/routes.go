// package user

// import (
// 	//"ecommerce/rest/handlers"
// 	"ecommerce/rest/middleware"
// 	"net/http"
// )

// func (h Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
// 	mux.Handle(
// 		"POST /users",
// 		manager.With(
// 			http.HandlerFunc(h.CreateUser),
// 		),
// 	)

// 	mux.Handle(
// 		"POST /users/login",
// 		manager.With(
// 			http.HandlerFunc(h.Login),
// 		),
// 	)
// }

package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	auth := middleware.NewMiddlewares(h.cnf).AuthenticateJwt

	mux.Handle(
		"POST /users",
		manager.With(http.HandlerFunc(h.CreateUser)),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(http.HandlerFunc(h.Login)),
	)

	mux.Handle(
		"GET /users/me",
		manager.With(http.HandlerFunc(h.GetMe), auth),
	)

	mux.Handle(
		"PATCH /users/me",
		manager.With(http.HandlerFunc(h.UpdateMe), auth),
	)
}