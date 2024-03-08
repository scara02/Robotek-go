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
	userRepo := repository.NewUserRepo(db)
	userUseCase := usecase.NewUserUsecase(*userRepo)
	userDelivery := delivery.NewUserDelivery(userUseCase)

	router := mux.NewRouter()

	// Auth routes
	router.HandleFunc("/signin", userDelivery.SignIn).Methods("POST")

	// Teacher routes
	teacher := router.PathPrefix("/teacher").Subrouter()
	teacher.Use(userDelivery.IsAuthenticated)
	teacher.HandleFunc("/", teacherDelivery.CreateHandler).Methods("POST")
	teacher.HandleFunc("/", teacherDelivery.GetAllHandler).Methods("GET")
	teacher.HandleFunc("/{id}", teacherDelivery.GetOneHandler).Methods("GET")
	teacher.HandleFunc("/{id}", teacherDelivery.DeleteHandler).Methods("DELETE")
	teacher.HandleFunc("/{teacherID}/group/{groupID}", teacherDelivery.AddToGroupHandler).Methods("POST")
	teacher.HandleFunc("/{id}/groups", teacherDelivery.GetGroupsHandler).Methods("GET")
	teacher.HandleFunc("/{teacherID}/group/{groupID}", teacherDelivery.DeleteGroupHandler).Methods("DELETE")
	teacher.HandleFunc("/{id}", teacherDelivery.UpdateHandler).Methods("PUT")

	// Student routes
	student := router.PathPrefix("/student").Subrouter()
	student.Use(userDelivery.IsAuthenticated)
	student.HandleFunc("/", studentDelivery.CreateHandler).Methods("POST")
	student.HandleFunc("/", studentDelivery.GetAllHandler).Methods("GET")
	student.HandleFunc("/{id}", studentDelivery.GetOneHandler).Methods("GET")
	student.HandleFunc("/{id}/group", studentDelivery.GetGroupHandler).Methods("GET")
	student.HandleFunc("/{id}", studentDelivery.DeleteHandler).Methods("DELETE")
	student.HandleFunc("/{studentID}/group/{groupID}", studentDelivery.ChangeGroupHandler).Methods("PUT")
	student.HandleFunc("/{id}", studentDelivery.UpdateHandler).Methods("PUT")

	// Group routes
	group := router.PathPrefix("/group").Subrouter()
	group.Use(userDelivery.IsAuthenticated)
	group.HandleFunc("/", groupDelivery.CreateHandler).Methods("POST")
	group.HandleFunc("/", groupDelivery.GetAllHandler).Methods("GET")
	group.HandleFunc("/{id}", groupDelivery.GetOneHandler).Methods("GET")
	group.HandleFunc("/{id}", groupDelivery.DeleteHandler).Methods("DELETE")
	group.HandleFunc("/{id}", groupDelivery.UpdateHandler).Methods("PUT")

	log.Print("starting server on :4000")

	err = http.ListenAndServe(":4000", router)
	log.Fatal(err)

}
