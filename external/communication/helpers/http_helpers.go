package helpers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
		router.Name(route.Name).Path(route.Path).Methods(route.Methods...).HandlerFunc(addRouteLogging(route.Handler))
	}
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
}

// addRouteLogging adds a route logline to a HTTP HandlerFunc
func addRouteLogging(in http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		logrus.Infof("Received HTTP %s request at route %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		in(rw, r)
	}
}
