package main

import (
	"log"
	"net/http"

	"clean-architecture/api"
	"clean-architecture/infra"
	"clean-architecture/repository"
	"clean-architecture/service"
)

func main() {
	db, err := infra.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepoSQL(db)
	svc := service.NewUserService(repo)
	handler := api.NewUserHandler(svc)

	http.HandleFunc("/register", handler.Register)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
