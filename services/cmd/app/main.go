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
	mux.HandleFunc("/teacher/{id}", teacherDelivery.DeleteHandler).Methods("DELETE")
	mux.HandleFunc("/teacher/{teacherID}/group/{groupID}", teacherDelivery.AddToGroupHandler).Methods("POST")
	mux.HandleFunc("/teacher/{id}/groups", teacherDelivery.GetGroupsHandler).Methods("GET")
	mux.HandleFunc("/teacher/{teacherID}/group/{groupID}", teacherDelivery.DeleteGroupHandler).Methods("DELETE")
	mux.HandleFunc("/teacher/{id}", teacherDelivery.UpdateHandler).Methods("PUT")

	// Student routes
	mux.HandleFunc("/student", studentDelivery.CreateHandler).Methods("POST")
	mux.HandleFunc("/student", studentDelivery.GetAllHandler).Methods("GET")
	mux.HandleFunc("/student/{id}", studentDelivery.GetOneHandler).Methods("GET")
	mux.HandleFunc("/student/{id}/group", studentDelivery.GetGroupHandler).Methods("GET")
	mux.HandleFunc("/student/{id}", studentDelivery.DeleteHandler).Methods("DELETE")
	mux.HandleFunc("/student/{studentID}/group/{groupID}", studentDelivery.ChangeGroupHandler).Methods("PUT")
	mux.HandleFunc("/student/{id}", studentDelivery.UpdateHandler).Methods("PUT")

	// Group routes
	mux.HandleFunc("/group", groupDelivery.CreateHandler).Methods("POST")
	mux.HandleFunc("/group", groupDelivery.GetAllHandler).Methods("GET")
	mux.HandleFunc("/group/{id}", groupDelivery.GetOneHandler).Methods("GET")
	mux.HandleFunc("/group/{id}", groupDelivery.DeleteHandler).Methods("DELETE")
	mux.HandleFunc("/group/{id}", groupDelivery.UpdateHandler).Methods("PUT")

	log.Print("starting server on :4000")

	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
