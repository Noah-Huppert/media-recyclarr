package httpapi

import (
	"context"
	"fmt"
	"github.com/Noah-Huppert/media-recyclarr/trasher"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// HTTPAPI is the HTTP API server
type HTTPAPI struct {
	// logger is used to output runtime information
	logger *zap.Logger

	// router knows what endpoint to send requests to
	router chi.Router

	// api runs the HTTP server
	api huma.API

	// address on which HTTP API will listen
	address string

	// trasher finds content that needs to be recycled
	trasher *trasher.Trasher
}

// NewHTTPAPIOpts are options for NewHTTPAPI()
type NewHTTPAPIOpts struct {
	// Logger is used to output runtime information
	Logger *zap.Logger

	// Trasher to use
	Trasher *trasher.Trasher

	// Address on which HTTP API will listen
	Address string
}

// NewHTTPAPI creates a new HTTPAPI
func NewHTTPAPI(opts NewHTTPAPIOpts) HTTPAPI {
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("Media Recylarr", "0.0.0"))

	httpAPI := HTTPAPI{
		logger:  opts.Logger,
		router:  router,
		api:     api,
		address: opts.Address,
		trasher: opts.Trasher,
	}
	httpAPI.registerEndpoints()

	return httpAPI
}

// Run listens and serves the HTTP API
func (h HTTPAPI) Run(gracefulCtx context.Context, harshCtx context.Context) error {
	var wg errgroup.Group
	earlyExit := make(chan interface{})

	server := &http.Server{
		Addr:    h.address,
		Handler: h.router,
	}

	wg.Go(func() error {
		h.logger.Info("starting", zap.String("address", h.address))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			earlyExit <- nil
			return fmt.Errorf("failed to listen and serve HTTP API: %s", err)
		}

		return nil
	})

	wg.Go(func() error {
		select {
		case <-gracefulCtx.Done():

			h.logger.Info("Gracefully shutting down")

			if err := server.Shutdown(harshCtx); err != nil {
				return fmt.Errorf("failed to shut down HTTP API server: %s", err)
			}

			h.logger.Debug("shut down")

			return nil
		case <-earlyExit:
			return nil
		}
	})

	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}

// registerEndpoints registers the endpoint
func (h HTTPAPI) registerEndpoints() {
	huma.Register(h.api, huma.Operation{
		OperationID: "health",
		Method:      http.MethodGet,
		Path:        "/api/v0/health",
		Summary:     "Get HTTP API health",
		Tags:        []string{"health"},
	}, h.epHealth)
}

// epHealthResp is the response for epHealth()
type epHealthResp struct {
	// Body response
	Body struct {
		// Ok indicates if the server is okay
		Ok bool `json:"ok"`
	}
}

// epHealth gets the health of the server
func (h HTTPAPI) epHealth(ctx context.Context, req *struct{}) (*epHealthResp, error) {
	resp := &epHealthResp{}
	resp.Body.Ok = true

	return resp, nil
}
