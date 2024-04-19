package route

import (
	"expression-backend/api/handlers"
	"github.com/gorilla/mux"
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

	//router.r.HandleFunc("/expressions/{id:[0-9]+}", handlers.GetExpression).Methods(http.MethodGet)
	//router.r.HandleFunc("/expressions", handlers.GetExpressions).Methods(http.MethodGet)
	//router.r.HandleFunc("/expressions", handlers.StoreExpression).Methods(http.MethodPost)
	//router.r.HandleFunc("/agents", handlers.GetAgents).Methods(http.MethodGet)
}

func (router *RouterMux) GetRouter() *mux.Router {
	return router.r
}
