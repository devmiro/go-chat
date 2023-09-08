package routes

import (
	"github.com/devmiro/go-chat/controllers"
	"github.com/devmiro/go-chat/http/middlewares"
	"github.com/devmiro/go-chat/services"
	"github.com/gorilla/mux"
)

var RegisterChatRoutes = func(router *mux.Router) {

	sb := router.PathPrefix("/v1/api/chat").Subrouter()
	sb.Use(middlewares.HeaderMiddleware)
	sb.Use(middlewares.Authenticated)

	var chat controllers.ChatController
	chat.RegisterService(services.NewChatService())

	sb.HandleFunc("/create", chat.Create).Methods("POST")
	sb.HandleFunc("/rooms", chat.ChatRooms).Methods("POST")
	sb.HandleFunc("/room-messages", chat.ChatRoomMessages).Methods("POST")
}
