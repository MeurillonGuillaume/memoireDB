package helpers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
)

func TestHTTPRouter(t *testing.T) {
	routes := []Route{
		{
			Name:    "status",
			Path:    "/",
			Methods: []string{http.MethodGet},
			Handler: testingHTTPHandler,
		},
		{
			Name:    "route-1",
			Path:    "/route-1",
			Methods: []string{http.MethodGet},
			Handler: testingHTTPHandler,
		},
	}

	port, err := freeport.GetFreePort()
	assert.NoError(t, err)
	server := NewHTTPServer(port, routes)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil && err != context.Canceled {
			assert.NoError(t, err)
		}
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			assert.NoError(t, err)
		}
	}()
	awaitOnline(fmt.Sprintf("http://localhost:%d", port))
}

func testingHTTPHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("Ok"))
	if err != nil {
		http.Error(rw, "Could not write status", http.StatusInternalServerError)
	}
}

func awaitOnline(serverAddr string) {
	var (
		err = errors.New("i could fail for a while")
	)

	for err != nil {
		_, err = http.Get(serverAddr)
	}
}
