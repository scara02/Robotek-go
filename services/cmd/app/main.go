package main

import (
	"log"
	"net/http"

	"go-web-robotek/pkg/store/postgres"
	"go-web-robotek/services/internal/delivery"
	"go-web-robotek/services/internal/repository"
	"go-web-robotek/services/internal/usecase"
)

func main() {
	db, err := postgres.DBconnection("postgres", "1234", "localhost", "5432", "robotek")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer db.Close()

	teacherRepo := repository.NewTeacherRepo(db)
	teacherUseCase := usecase.NewTeacherUsecase(*teacherRepo)
	teacherDelivery := delivery.NewTeacherDelivery(teacherUseCase)

	mux := http.NewServeMux()

	mux.HandleFunc("/teachers/create", teacherDelivery.CreateHandler)

	log.Print("starting server on :4000")
	
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
