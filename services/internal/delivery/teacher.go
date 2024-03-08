package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/usecase"

	"github.com/gorilla/mux"
)

type TeacherDelivery struct {
	teacherUseCase usecase.Teacher
}

func NewTeacherDelivery(teacherUseCase usecase.Teacher) *TeacherDelivery {
	return &TeacherDelivery{
		teacherUseCase: teacherUseCase,
	}
}

func (d *TeacherDelivery) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var t domain.Teacher

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := d.teacherUseCase.Create(t.FullName, t.Email, t.Password, t.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (d *TeacherDelivery) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teacher, err := d.teacherUseCase.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	teacherJSON, err := json.Marshal(teacher)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(teacherJSON)))
}

func (d *TeacherDelivery) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	teachers, err := d.teacherUseCase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	teachersJSON, err := json.Marshal(teachers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(teachersJSON)))
}

func (d *TeacherDelivery) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedID, err := d.teacherUseCase.Delete(id)
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

func (d *TeacherDelivery) AddToGroupHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	teacherID, err := strconv.Atoi(params["teacherID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupID, err := strconv.Atoi(params["groupID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = d.teacherUseCase.AddToGroup(teacherID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *TeacherDelivery) GetGroupsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groups, err := d.teacherUseCase.GetGroups(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	groupsJSON, err := json.Marshal(groups)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(groupsJSON)
}

func (d *TeacherDelivery) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	teacherID, err := strconv.Atoi(params["teacherID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupID, err := strconv.Atoi(params["groupID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = d.teacherUseCase.DeleteGroup(teacherID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *TeacherDelivery) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedTeacher domain.Teacher
	if err := json.NewDecoder(r.Body).Decode(&updatedTeacher); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = d.teacherUseCase.Update(id, updatedTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}