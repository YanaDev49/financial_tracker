package routes

import (
	"collab/handlers"
	"collab/middleware"
	"net/http"
    "github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, TransactionHandler *handlers.TransactionHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register/user", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
	r.HandleFunc("/return/user", userHandler.GetUser).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.Handle("/make/transaction", middleware.AuthMiddleware(http.HandlerFunc(TransactionHandler.CreateTransaction))).Methods("POST")
	protected.Handle("/update/transaction", middleware.AuthMiddleware(http.HandlerFunc(TransactionHandler.UpdateTransaction))).Methods("PUT")
	protected.Handle("/delete/transaction", middleware.AuthMiddleware(http.HandlerFunc(TransactionHandler.DeleteTransaction))).Methods("POST")
	protected.Handle("/return/transactions", middleware.AuthMiddleware(http.HandlerFunc(TransactionHandler.GetTransaction))).Methods("GET")

	return r


}