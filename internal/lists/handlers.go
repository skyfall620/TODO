package lists

import (
	"net/http"
	"todo/config"
	"todo/internal/middleware"
	"todo/pkg/req"
	"todo/pkg/res"
)

type ListsHandlerDeps struct {
	*config.Config
	*ListsService
}

type ListsHandler struct {
	*ListsService
}

func NewListsHandler(router *http.ServeMux, deps ListsHandlerDeps) {
	handler := &ListsHandler{
		ListsService: deps.ListsService,
	}

	router.Handle("POST /lists", middleware.AuthMiddleware(handler.CreateList(), deps.Config.Auth.Secret))
	router.Handle("GET /lists", middleware.AuthMiddleware(handler.AllList(), deps.Config.Auth.Secret))
}

func (lh *ListsHandler) CreateList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			res.JSONWrite(w, ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		text, err := req.HandleBody[BodyRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		list, err := lh.ListsService.Create(r.Context(), userId, text.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSONWrite(w, list, http.StatusCreated)
	}
}

func (lh *ListsHandler) AllList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		lists, err := lh.ListsService.GetAllByUserID(r.Context(), userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSONWrite(w, lists, http.StatusOK)
	}
}
