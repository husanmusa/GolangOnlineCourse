package main

import (
	"expensity/handler"
	"expensity/postgres"
	"expensity/repository"
	"expensity/service"
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
