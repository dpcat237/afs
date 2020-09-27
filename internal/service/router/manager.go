package router

import (
	"context"
	"net/http"
	"time"

	"github.com/dpcat237/afs/internal/controller"
	"github.com/dpcat237/afs/internal/interfaces"
)

const (
	//systemAPI prefix for system endpoints.
	systemAPI = "system"
	//v1API prefix for version 1 endpoints.
	v1API = "apiv1"

	//rateLimit the limit of requests per second.
	rateLimit = 100
	//serviceReadTimeout the time limit for reading a request.
	serviceReadTimeout = 10 * time.Second
)

//Manager contains required details for router management.
type Manager struct {
	lgr  interfaces.Logger
	lmt  *Limiter
	port string
	rtr  *Router
	srv  *http.Server
}

//NewManager initializes required details for HTTP server.
func NewManager(cnts controller.Collector, lgr interfaces.Logger, port string) Manager {
	var mng Manager
	mng.lgr = lgr
	mng.lmt = NewLimiter(1, rateLimit)
	mng.port = port

	mng.createRouter(cnts)
	mng.defineHTTPServer()

	return mng
}

//Start initializes HTTP server's listener.
func (mng *Manager) Start() {
	go func() {
		if err := mng.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mng.lgr.Fatal("Error starting HTTP service: " + err.Error())
		}
	}()
}

//Shutdown processes graceful shutdown of router connections.
func (mng *Manager) Shutdown(ctx context.Context) {
	if err := mng.srv.Shutdown(ctx); err != nil {
		mng.lgr.Warn("Error stopping HTTP router " + err.Error())
	}
}

func (mng *Manager) createRouter(cnts controller.Collector) {
	mng.rtr = NewRouter()
	mng.rtr.Endpoint(systemAPI + "/health").Get(cnts.Controller.HealthCheck)
	mng.rtr.Endpoint(v1API + "/link/process").Post(cnts.Link.ProcessLinks)
}

func (mng *Manager) defineHTTPServer() {
	mng.srv = &http.Server{
		Addr:        ":" + mng.port,
		Handler:     mng.panicRecovery(mng.limit(mng.rtr)),
		ReadTimeout: serviceReadTimeout,
	}
}

func (mng Manager) limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mng.lmt.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (mng Manager) panicRecovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				mng.lgr.Infow("Panic recovered",
					// Structured context as loosely typed key-value pairs
					"method", r.Method,
					"path", r.URL.Path,
					"remote_addr", r.RemoteAddr,
					"user_agent", r.UserAgent(),
					"error", err,
				)

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
