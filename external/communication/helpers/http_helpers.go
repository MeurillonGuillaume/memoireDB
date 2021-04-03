package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	_contentType = "Content-Type"
	_appJson     = "application/json"
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

// HTTPReplyJSON attempts to reply to a HTTP ResponseWriter using an interface as JSON body, or sends an error response
func HTTPReplyJSON(rw http.ResponseWriter, statusCode int, body interface{}) {
	rawReply, err := json.Marshal(body)
	if err != nil {
		logrus.WithError(err).Error("Could not marshal reply body")
		http.Error(rw, "Response could not be properly encoded", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(statusCode)
	rw.Header().Add(_contentType, _appJson)
	_, err = rw.Write(rawReply)
	if err != nil {
		logrus.WithError(err).Error("Could not write response to client")
		http.Error(rw, "Actual response could not be written", http.StatusInternalServerError)
		return
	}
}
