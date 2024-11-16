package routes

import (
    "database/sql"
    "net/http"

    "backend/middleware"
    "backend/services"

    "github.com/gorilla/mux"
)

func InitRoutes(db *sql.DB) *mux.Router {
    router := mux.NewRouter()

    // User routes
    userService := services.NewUserService(db)
    router.HandleFunc("/register", userService.RegisterUser).Methods("POST")
    router.HandleFunc("/login", userService.LoginUser).Methods("POST")

    // Authenticated routes
    protected := router.PathPrefix("/api").Subrouter()
    protected.Use(middleware.JWTMiddleware("your_jwt_secret", http.DefaultServeMux))
    protected.HandleFunc("/profile", userService.GetProfile).Methods("GET")

    return router
}
