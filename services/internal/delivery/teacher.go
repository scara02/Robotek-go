package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/usecase"
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
