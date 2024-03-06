package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/usecase"

	"github.com/gorilla/mux"
)

type GroupDelivery struct {
	groupUseCase usecase.Group
}

func NewGroupDelivery(groupUseCase usecase.Group) *GroupDelivery {
	return &GroupDelivery{
		groupUseCase: groupUseCase,
	}
}

func (d *GroupDelivery) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var g domain.Group

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := d.groupUseCase.Create(g.GroupName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (d *GroupDelivery) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group, err := d.groupUseCase.GetOne(id)
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

func (d *GroupDelivery) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	groups, err := d.groupUseCase.GetAll()
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
	w.Write([]byte(string(groupsJSON)))
}
