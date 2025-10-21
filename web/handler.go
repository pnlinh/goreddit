package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/pnlinh/goreddit"
	"github.com/pnlinh/goreddit/stores"
)

type Handler struct {
	*chi.Mux
	store stores.Store
}

func NewHandler(store stores.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
		r.Post("/", h.ThreadStore())
		r.Delete("/{id}", h.ThreadDestroy())
	})

	return h
}

func (h *Handler) ThreadsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := h.store.List()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func (h *Handler) ThreadStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")
		err := h.store.CreateThread(&goreddit.Thread{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *Handler) ThreadDestroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = h.store.DeleteThread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
