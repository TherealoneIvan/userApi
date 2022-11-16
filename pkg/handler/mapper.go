package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"refactoring/pkg/handler/requests"
)

func (h *Handler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	UserStore, err := h.userApi.SearchUsers()
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, UserStore.List)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := requests.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	id, err := h.userApi.CreateUser(request.DisplayName, request.Email)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})

}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.userApi.GetUser(id)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	id := chi.URLParam(r, "id")

	err := h.userApi.UpdateUser(id, request)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusNoContent)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.userApi.DeleteUser(id)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusNoContent)
}
