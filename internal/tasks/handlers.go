package tasks

import (
	"net/http"
	"strconv"
	"todo/config"
	"todo/internal/middleware"
	"todo/pkg/req"
	"todo/pkg/res"
)

type TaskHandlersDeps struct {
	*config.Config
	*TasksService
}

type TaskHandlers struct {
	*TasksService
}

func NewTaskHandlers(router *http.ServeMux, deps TaskHandlersDeps) {
	handler := &TaskHandlers{
		TasksService: deps.TasksService,
	}

	router.Handle("POST /lists/{id}/tasks", middleware.AuthMiddleware(handler.CreateTask(), deps.Config.Auth.Secret))
	router.Handle("GET /lists/{id}/tasks", middleware.AuthMiddleware(handler.AllTasks(), deps.Config.Auth.Secret))
}

func (th *TaskHandlers) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		idStr := r.PathValue("id")
		listId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := req.HandleBody[TaskRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, err := th.TasksService.Create(r.Context(), userId, listId, body.Title, body.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSONWrite(w, task, http.StatusCreated)
	}
}

func (th *TaskHandlers) AllTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		idStr := r.PathValue("id")
		listId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tasks, err := th.TasksService.GetAllByListID(r.Context(), userId, listId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSONWrite(w, tasks, http.StatusOK)
	}
}
