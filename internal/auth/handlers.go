package auth

import (
	"net/http"
	"todo/config"
	"todo/pkg/jwt"
	"todo/pkg/req"
	"todo/pkg/res"
)

type AuthHandlerDeps struct {
	*config.Config
	*AuthService
}

type AuthHandler struct {
	*config.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (ah *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}

		userId, err := ah.AuthService.Register(ctx, user.Name, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(ah.Auth.Secret).Create(jwt.JWTdata{
			UserId: userId,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegisterResponse{
			Token: token,
		}

		res.JSONWrite(w, data, http.StatusCreated)
	}
}

func (ah *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}

		userId, err := ah.AuthService.Login(ctx, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(ah.Auth.Secret).Create(jwt.JWTdata{
			UserId: userId,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponse{
			Token: token,
		}

		res.JSONWrite(w, data, http.StatusCreated)
	}
}
