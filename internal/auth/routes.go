package auth

import "net/http"

func GetAuthRoutesMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/user", createUserHandler)
	mux.HandleFunc("POST /v1/auth/login", loginUserHandler)
	mux.HandleFunc("POST /v1/auth/logout", logoutUserHandler)

	return mux
}
