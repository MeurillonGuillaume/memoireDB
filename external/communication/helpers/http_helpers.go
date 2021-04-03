package helpers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route is an object used to create a HTTP routing object
type Route struct {
	Name    string
	Path    string
	Methods []string
	Handler http.HandlerFunc
}

// NewHTTPServer creates a new HTTP server with a router created from a slice of given HTTP routes
func NewHTTPServer(port int, routes []Route) *http.Server {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Name(route.Name).Path(route.Path).Methods(route.Methods...).HandlerFunc(route.Handler)
	}
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
}

// func AddRouteLogging(in http.HandlerFunc) (out http.HandlerFunc) {
// 	// TODO: add route logging overlay
// 	return
// }
