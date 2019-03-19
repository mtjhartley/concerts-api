package main

import (
	"fmt"

	"github.com/mtjhartley/concerts-api/internal/app/concerts"
	"github.com/mtjhartley/concerts-api/internal/pkg/auth"
	"github.com/mtjhartley/concerts-api/internal/pkg/controllers"

	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	concerts.DoWork()
	router := mux.NewRouter()
	router.Use(auth.JwtAuthentication) //attach JWT auth middleware

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
