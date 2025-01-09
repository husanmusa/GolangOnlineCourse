package main

import (
	"lesson28/handler"
	"lesson28/postgres"
	"lesson28/repository"
	"lesson28/service"
)

func main() {

	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(service.NewUserService(repository.NewUserRepo(db)))

	r := handler.Run(h)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
