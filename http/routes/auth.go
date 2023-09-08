package routes

import (
	"github.com/devmiro/go-chat/controllers"
	"github.com/devmiro/go-chat/http/middlewares"
	"github.com/devmiro/go-chat/services"
	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router) {

	sb := router.PathPrefix("/v1/api/auth").Subrouter()
	sb.Use(middlewares.HeaderMiddleware)

	var auth controllers.AuthController
	auth.RegisterService(services.NewAuthService())

	sb.HandleFunc("/login", auth.Login).Methods("POST")
	sb.HandleFunc("/signup", auth.SignUp).Methods("POST")
}
