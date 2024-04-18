package route

import (
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
	//router.r.HandleFunc("/expressions/{id:[0-9]+}", handlers.GetExpression).Methods(http.MethodGet)
	//router.r.HandleFunc("/expressions", handlers.GetExpressions).Methods(http.MethodGet)
	//router.r.HandleFunc("/expressions", handlers.StoreExpression).Methods(http.MethodPost)
	//router.r.HandleFunc("/agents", handlers.GetAgents).Methods(http.MethodGet)
	router.r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
		writer.WriteHeader(http.StatusOK)
	})
}

func (router *RouterMux) GetRouter() *mux.Router {
	return router.r
}
