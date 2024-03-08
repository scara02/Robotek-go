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

	mux := mux.NewRouter()

	// Auth routes
	mux.HandleFunc("/signin", userDelivery.SignIn).Methods("POST")

	// Teacher routes
	mux.Handle("/teacher", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.CreateHandler))).Methods("POST")
	mux.Handle("/teacher", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.GetAllHandler))).Methods("GET")
	mux.Handle("/teacher/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.GetOneHandler))).Methods("GET")
	mux.Handle("/teacher/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.DeleteHandler))).Methods("DELETE")
	mux.Handle("/teacher/{teacherID}/group/{groupID}", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.AddToGroupHandler))).Methods("POST")
	mux.Handle("/teacher/{id}/groups", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.GetGroupsHandler))).Methods("GET")
	mux.Handle("/teacher/{teacherID}/group/{groupID}", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.DeleteGroupHandler))).Methods("DELETE")
	mux.Handle("/teacher/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(teacherDelivery.UpdateHandler))).Methods("PUT")

	// Student routes
	mux.Handle("/student", userDelivery.IsAuthenticated(http.HandlerFunc(studentDelivery.CreateHandler))).Methods("POST")
	mux.Handle("/student", userDelivery.IsAuthenticated(http.HandlerFunc(studentDelivery.GetAllHandler))).Methods("GET")
	mux.Handle("/student/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(studentDelivery.GetOneHandler))).Methods("GET")
	mux.Handle("/student/{id}/group", userDelivery.IsAuthenticated(http.HandlerFunc(studentDelivery.GetGroupHandler))).Methods("GET")
	mux.Handle("/student/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(studentDelivery.DeleteHandler))).Methods("DELETE")
	mux.Handle("/student/{studentID}/group/{groupID}", userDelivery.IsAuthenticated(http.HandlerFunc(studentDelivery.ChangeGroupHandler))).Methods("PUT")
	mux.Handle("/student/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(studentDelivery.UpdateHandler))).Methods("PUT")

	// Group routes
	mux.Handle("/group", userDelivery.IsAuthenticated(http.HandlerFunc(groupDelivery.CreateHandler))).Methods("POST")
	mux.Handle("/group", userDelivery.IsAuthenticated(http.HandlerFunc(groupDelivery.GetAllHandler))).Methods("GET")
	mux.Handle("/group/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(groupDelivery.GetOneHandler))).Methods("GET")
	mux.Handle("/group/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(groupDelivery.DeleteHandler))).Methods("DELETE")
	mux.Handle("/group/{id}", userDelivery.IsAuthenticated(http.HandlerFunc(groupDelivery.UpdateHandler))).Methods("PUT")

	log.Print("starting server on :4000")

	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
