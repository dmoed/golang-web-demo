package server

import (
	"apptastic/dashboard/auth"
	"apptastic/dashboard/env"
	"apptastic/dashboard/handler"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type key int

const (
	requestIDKey key = 0
)

//ServerConfig is the server config struct
type ServerConfig struct {
	Host      string
	HTTPPort  string
	HTTPSPort string
	CertFile  string
	KeyFile   string
}

type Server struct {
	Config ServerConfig
	env    *env.Env
}

//NewServer creates a new Server struct
func NewServer(c ServerConfig, e *env.Env) *Server {
	return &Server{
		Config: c,
		env:    e,
	}
}

//Run the server
func (s *Server) Run(port string) {

	s.env.Logger.Println("Server is starting...")

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	router := mux.NewRouter()

	//Static files
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./frontend/dist/"))))

	//Auth
	router.PathPrefix("/login").Handler(handler.LoginHandler(s.env.DB))
	router.PathPrefix("/logout").Handler(handler.LogoutHandler())
	//todo: request password reset

	//Ajax
	router.PathPrefix(`/ajax/profile`).Handler(auth.MustAuthorize(handler.ProfileHandler(s.env.DB)))
	router.PathPrefix(`/ajax/products`).Handler(auth.MustAuthorize(handler.ProductHandler(s.env.DB)))
	router.PathPrefix(`/ajax/inventory/total`).Handler(auth.MustAuthorize(handler.InventoryTotalSummaryHandler(s.env.DB)))
	router.PathPrefix(`/ajax/inventory`).Handler(auth.MustAuthorize(handler.InventoryHandler(s.env.DB)))

	//Dashboard
	router.PathPrefix(`/dashboard/`).Handler(auth.MustAuthorize(handler.DashboardHandler(s.env.DB, s.env.View)))
	router.PathPrefix(`/dashboard/{rest:[a-zA-Z0-9=\-\/]+}`).Handler(auth.MustAuthorize(handler.DashboardHandler(s.env.DB, s.env.View)))

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      tracing(nextRequestID)(logging(s.env.Logger)(recoverWrapper(router))),
		ErrorLog:     s.env.Logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	s.env.Logger.Println(fmt.Sprintf("Server ready at :%s", port))

	if err := httpServer.ListenAndServe(); err != nil {
		s.env.Logger.Fatal(err)
	}

	// go func() {

	// s.env.Logger.Println(fmt.Sprintf("HTTP server ready at :%s", s.Config.HTTPPort))

	// if err := http.ListenAndServe(fmt.Sprintf(":%s", s.Config.HTTPPort), http.HandlerFunc(redirectTLS)); err != nil {
	// 	s.env.Logger.Fatal(err)
	// }
	// }()

	// s.env.Logger.Println(fmt.Sprintf("HTPPS server ready at :%s", s.Config.HTTPSPort))

	// if err := httpServer.ListenAndServeTLS(s.Config.CertFile, s.Config.KeyFile); err != nil {
	// 	s.env.Logger.Fatal(err)
	// }
}

// func redirectTLS(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "https://localhost:443"+r.RequestURI, http.StatusMovedPermanently)
// }

//tracing
func tracing(nextRequestID func() string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			requestID := r.Header.Get("X-Request-Id")

			if requestID == "" {
				requestID = nextRequestID()
			}

			ctx := context.WithValue(r.Context(), requestIDKey, requestID)

			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//logging
func logging(logger *log.Logger) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//loggin
			defer func() {

				requestID, ok := r.Context().Value(requestIDKey).(string)

				if !ok {
					requestID = "unknown"
				}

				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

			}()

			next.ServeHTTP(w, r)
		})
	}
}

//recover
func recoverWrapper(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		defer func() {

			r := recover()

			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
