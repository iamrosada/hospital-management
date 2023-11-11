package usecase

import (
	"github.com/iamrosada/hospital-management/doctor-service/internal/doctor/entity"
)

type UpdateDoctorInputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BI          string `json:"bi"`
	Specialty   string `json:"specialty"`
	Experience  int    `json:"experience"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateDoctorOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BI          string `json:"bi"`
	Specialty   string `json:"specialty"`
	Experience  int    `json:"experience"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateDoctorUseCase struct {
	DoctorRepository entity.DoctorRepository
}

func NewUpdateDoctorUseCase(DoctorRepository entity.DoctorRepository) *UpdateDoctorUseCase {
	return &UpdateDoctorUseCase{DoctorRepository: DoctorRepository}
}

func (u *UpdateDoctorUseCase) Execute(input UpdateDoctorInputDto) (*UpdateDoctorOutputDto, error) {
	existingDoctor, err := u.DoctorRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Update the existing doctor with the new information
	existingDoctor.Update(
		input.Name,
		input.BI,
		input.Specialty,
		input.Email,
		input.PhoneNumber,
		input.Experience,
	)

	err = u.DoctorRepository.Update(existingDoctor)
	if err != nil {
		return nil, err
	}

	return &UpdateDoctorOutputDto{
		ID:          existingDoctor.ID,
		Name:        existingDoctor.Name,
		BI:          existingDoctor.BI,
		Specialty:   existingDoctor.Specialty,
		Experience:  existingDoctor.Experience,
		Email:       existingDoctor.Email,
		PhoneNumber: existingDoctor.PhoneNumber,
	}, nil
}
