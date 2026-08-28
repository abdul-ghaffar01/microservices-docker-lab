package main

import (
	"backend/db"
	"backend/users"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Connecting with database
	dbPgx, err := db.ConnectDB()

	db.CreateSchema(dbPgx)
	
	if err != nil {
		panic(err)
	}

	ur := users.NewUserRespository(dbPgx)

	us := users.NewUserService(ur)

	uh := users.NewUserHandler(us)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", uh.CreateNewUser)

	mux.HandleFunc("GET /users", uh.GetUsers)

	PORT := os.Getenv("BACKEND_PORT")

	fmt.Println("Server is listerning")
	http.ListenAndServe(":"+PORT, mux)
}
