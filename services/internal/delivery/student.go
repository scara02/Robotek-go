package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/usecase"

	"github.com/gorilla/mux"
)

type StudentDelivery struct {
	studentUseCase usecase.Student
}

func NewStudentDelivery(studentUseCase usecase.Student) *StudentDelivery {
	return &StudentDelivery{
		studentUseCase: studentUseCase,
	}
}

func (d *StudentDelivery) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var s domain.Student

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := d.studentUseCase.Create(s.FullName, s.Email, s.Password, s.PhoneNumber, s.GroupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (d *StudentDelivery) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student, err := d.studentUseCase.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentJSON, err := json.Marshal(student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(studentJSON)))
}

func (d *StudentDelivery) GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group, err := d.studentUseCase.GetGroup(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	groupJSON, err := json.Marshal(group)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(groupJSON)))
}

func (d *StudentDelivery) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	students, err := d.studentUseCase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentsJSON, err := json.Marshal(students)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(studentsJSON)))
}


func (d *StudentDelivery) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedID, err := d.studentUseCase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idJSON, err := json.Marshal(deletedID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(idJSON)))
}

func (d *StudentDelivery) ChangeGroupHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	studentID, err := strconv.Atoi(params["studentID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupID, err := strconv.Atoi(params["groupID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = d.studentUseCase.ChangeGroup(studentID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *StudentDelivery) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedStudent domain.Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = d.studentUseCase.Update(id, updatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}