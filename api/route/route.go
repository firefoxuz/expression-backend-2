package route

import (
	"expression-backend/api/handlers"
	"expression-backend/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type RouterMux struct {
	r *mux.Router
}

func NewRouterMux() *RouterMux {
	return &RouterMux{
		r: mux.NewRouter(),
	}
}

func (router *RouterMux) RegisterRoutes() {
	router.r.HandleFunc("/api/v1/register", handlers.RegisterUser)
	router.r.HandleFunc("/api/v1/login", handlers.LoginUser)

	router.r.HandleFunc("/api/v1/expressions/{id:[0-9]+}", middleware.AuthMiddleware(handlers.GetUserExpressions)).Methods(http.MethodGet)
	router.r.HandleFunc("/api/v1/expressions", middleware.AuthMiddleware(handlers.GetUserExpressions)).Methods(http.MethodGet)
	router.r.HandleFunc("/api/v1/expressions", middleware.AuthMiddleware(handlers.StoreUserExpression)).Methods(http.MethodPost)
}

func (router *RouterMux) GetRouter() *mux.Router {
	return router.r
}
