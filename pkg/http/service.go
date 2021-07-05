package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sergionunezgo/go-reuse/pkg/logger"
	"github.com/urfave/negroni"
)

// Handler is responsible for defining a HTTP request route and corresponding handler.
type Handler interface {
	// ServeHTTP should handle HTTP requests.
	ServeHTTP(w http.ResponseWriter, r *http.Request)

	// AddRoute should allow the handler to configure itself accepting a router.
	AddRoute(r *mux.Router)
}

// HttpService is the struct that will hold references to all necessary data for
// running an http server.
type HttpService struct {
	server *http.Server
	router *mux.Router
}

// Start will begin listening on the host:port for requests.
// Blocking call.
func (h *HttpService) Start() error {
	logger.Log.Infof("service listening on address: %v", h.server.Addr)
	return h.server.ListenAndServe()
}

// Close will teardown/close any resources used by this http service.
func (h *HttpService) Close() error {
	logger.Log.Info("closing http service")
	return h.server.Close()
}

// AddRoutes can be used to let each Handler setup their service routing.
func (h *HttpService) AddRoutes(handlers ...Handler) {
	for _, handler := range handlers {
		handler.AddRoute(h.router)
	}
}

// NewHttpService will run the setup process and create a Service that can be
// used to run a http api.
func NewHttpService(port int) *HttpService {
	r := mux.NewRouter().StrictSlash(true)

	NewNotFoundHandler().AddRoute(r)

	n := negroni.New()
	n.UseHandler(r)

	srv := &http.Server{
		Handler:      n,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &HttpService{
		server: srv,
		router: r,
	}
}
