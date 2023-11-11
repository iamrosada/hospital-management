package usecase

import (
	"github.com/iamrosada/hospital-management/doctor-service/internal/doctor/entity"
)

type CreateDoctorInputDto struct {
	Name        string `json:"name"`
	BI          string `json:"bi"`
	Specialty   string `json:"specialty"`
	Experience  int    `json:"experience"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CreateDoctorOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BI          string `json:"bi"`
	Specialty   string `json:"specialty"`
	Experience  int    `json:"experience"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CreateDoctorUseCase struct {
	DoctorRepository entity.DoctorRepository
}

func NewCreateDoctorUseCase(DoctorRepository entity.DoctorRepository) *CreateDoctorUseCase {
	return &CreateDoctorUseCase{DoctorRepository: DoctorRepository}
}

func (u *CreateDoctorUseCase) Execute(input CreateDoctorInputDto) (*CreateDoctorOutputDto, error) {
	Doctor := entity.NewDoctor(
		input.Name,
		input.BI,
		input.Specialty,
		input.Email,
		input.PhoneNumber,
		input.Experience,
	)

	err := u.DoctorRepository.Create(Doctor)
	if err != nil {
		return nil, err
	}

	return &CreateDoctorOutputDto{
		ID:          Doctor.ID,
		Name:        Doctor.Name,
		BI:          Doctor.BI,
		Specialty:   Doctor.Specialty,
		Experience:  Doctor.Experience,
		Email:       Doctor.Email,
		PhoneNumber: Doctor.PhoneNumber,
	}, nil
}
