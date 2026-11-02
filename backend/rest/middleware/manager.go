package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	GlobalMiddleWares []Middleware //nil slice
}

func NewMiddleware() *Manager {
	return &Manager{
		GlobalMiddleWares: make([]Middleware, 0), //empty slice
	}
}

func (mngr *Manager) Use(globalMiddleware ...Middleware) {
	mngr.GlobalMiddleWares = append(mngr.GlobalMiddleWares, globalMiddleware...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	// for i := len(middlewares) - 1; i>=0; i--{
	// 	middleware:= middlewares[i]
	// 	n= middleware(n)
	// }

	// for _, globalMiddleware := range mngr.GlobalMiddleWares {
	// 	n = globalMiddleware(n)
	// }

	return n

}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
	n := next

	for _, globalMiddleware := range mngr.GlobalMiddleWares {
		n = globalMiddleware(n)
	}

	return n

}