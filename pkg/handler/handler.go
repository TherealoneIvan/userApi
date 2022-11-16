package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"refactoring/api"
	"time"
)

type Handler struct {
	userApi *api.UserApi
}

func NewHandler(userApi *api.UserApi) *Handler {
	return &Handler{userApi: userApi}
}

const minute = 60

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(minute * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", h.SearchUsers)
				r.Post("/", h.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.GetUser)
					r.Patch("/", h.UpdateUser)
					r.Delete("/", h.DeleteUser)
				})
			})
		})
	})
	return r
}
