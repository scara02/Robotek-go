package main

import (
	"log"
	"net/http"

	"go-web-robotek/pkg/store/postgres"
	"go-web-robotek/services/internal/delivery"
	"go-web-robotek/services/internal/repository"
	"go-web-robotek/services/internal/usecase"

	"github.com/gorilla/mux"
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
	studentRepo := repository.NewStudentRepo(db)
	studentUseCase := usecase.NewStudentUsecase(*studentRepo)
	studentDelivery := delivery.NewStudentDelivery(studentUseCase)
	groupRepo := repository.NewGroupRepo(db)
	groupUseCase := usecase.NewGroupUsecase(*groupRepo)
	groupDelivery := delivery.NewGroupDelivery(groupUseCase)

	mux := mux.NewRouter()

	// Teacher routes
	mux.HandleFunc("/teacher", teacherDelivery.CreateHandler).Methods("POST")
	mux.HandleFunc("/teacher", teacherDelivery.GetAllHandler).Methods("GET")
	mux.HandleFunc("/teacher/{id}", teacherDelivery.GetOneHandler).Methods("GET")

	// Student routes
	mux.HandleFunc("/student", studentDelivery.CreateHandler).Methods("POST")
	mux.HandleFunc("/student/{id}", studentDelivery.GetOneHandler).Methods("GET")
	mux.HandleFunc("/student/{id}/group", studentDelivery.GetGroupHandler).Methods("GET")
	mux.HandleFunc("/student", studentDelivery.GetAllHandler).Methods("GET")

	// Group routes
	mux.HandleFunc("/group", groupDelivery.CreateHandler).Methods("POST")
	mux.HandleFunc("/group/{id}", groupDelivery.GetOneHandler).Methods("GET")
	mux.HandleFunc("/group", groupDelivery.GetAllHandler).Methods("GET")

	log.Print("starting server on :4000")

	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
