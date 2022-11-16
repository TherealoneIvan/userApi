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
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(minute * time.Second))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})
	router.Route("/api",
		func(router chi.Router) {
			router.Route("/v1",
				func(router chi.Router) {
					router.Route("/users",
						func(router chi.Router) {
							router.Get("/", h.SearchUsers)
							router.Post("/", h.CreateUser)
							router.Route("/{id}",
								func(router chi.Router) {
									router.Get("/", h.GetUser)
									router.Patch("/", h.UpdateUser)
									router.Delete("/", h.DeleteUser)
								})
						})
				})
		})
	return router
}
