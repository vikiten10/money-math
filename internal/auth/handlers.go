package auth

import "net/http"

func createUserHandler(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("User created"))
}

func loginUserHandler(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("User logged in"))
}

func logoutUserHandler(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("User logged out"))
}
