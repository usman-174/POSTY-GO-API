package app

import (
	"github.com/gorilla/mux"
	"github.com/usman-174/controller"
)

func Router() *mux.Router {
	// database.ConnectDataBase()

	router := mux.NewRouter()
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/getposts", controller.GetAllPosts).Methods("GET")
	router.HandleFunc("/getpost", controller.GetPost).Methods("GET")
	protectedRoutes := router.PathPrefix("/x").Subrouter()
	protectedRoutes.Use(controller.AuthMiddleware)
	protectedRoutes.HandleFunc("/user", controller.GetUser).Methods("GET")
	protectedRoutes.HandleFunc("/post", controller.Post).Methods("POST")
	protectedRoutes.HandleFunc("/logout", controller.Logout).Methods("POST")
	protectedRoutes.HandleFunc("/deletepost", controller.DeletePost).Methods("DELETE")
	protectedRoutes.HandleFunc("/updatepost", controller.UpdatePost).Methods("PUT")
	protectedRoutes.HandleFunc("/likepost", controller.Likepost).Methods("PUT")
	return router
}
