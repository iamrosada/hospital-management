// hospital-management/doctor-service/api/doctor.go

package api

import (
	"encoding/json"
	"net/http"

	usecase "github.com/iamrosada/hospital-management/doctor-service/internal/doctor/usecases"
)

type DoctorHandlers struct {
	CreateDoctorUseCase  *usecase.CreateDoctorUseCase
	ListDoctorsUseCase   *usecase.GetAllDoctorsUseCase
	DeleteDoctorUseCase  *usecase.DeleteDoctorUseCase
	GetDoctorByIDUseCase *usecase.GetDoctorByIDUseCase
	UpdateDoctorUseCase  *usecase.UpdateDoctorUseCase
}

func NewDoctorHandlers(
	createDoctorUseCase *usecase.CreateDoctorUseCase,
	listDoctorsUseCase *usecase.GetAllDoctorsUseCase,
	deleteDoctorUseCase *usecase.DeleteDoctorUseCase,
	getDoctorByIDUseCase *usecase.GetDoctorByIDUseCase,
	updateDoctorUseCase *usecase.UpdateDoctorUseCase,
) *DoctorHandlers {
	return &DoctorHandlers{
		CreateDoctorUseCase:  createDoctorUseCase,
		ListDoctorsUseCase:   listDoctorsUseCase,
		DeleteDoctorUseCase:  deleteDoctorUseCase,
		GetDoctorByIDUseCase: getDoctorByIDUseCase,
		UpdateDoctorUseCase:  updateDoctorUseCase,
	}
}

func (p *DoctorHandlers) CreateDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateDoctorInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateDoctorUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *DoctorHandlers) ListDoctorsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListDoctorsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *DoctorHandlers) DeleteDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.DeleteDoctorInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.DeleteDoctorUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *DoctorHandlers) GetDoctorByIDHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.GetDoctorByIDInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.GetDoctorByIDUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *DoctorHandlers) UpdateDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateDoctorInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.UpdateDoctorUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
