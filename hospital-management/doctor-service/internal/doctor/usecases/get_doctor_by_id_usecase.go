// usecase/get_doctor_by_id_usecase.go

package usecase

import (
	"github.com/iamrosada/hospital-management/doctor-service/internal/doctor/entity"
)

type GetDoctorByIDInputDto struct {
	ID string `json:"id"`
}

type GetDoctorByIDOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BI          string `json:"bi"`
	Specialty   string `json:"specialty"`
	Experience  int    `json:"experience"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type GetDoctorByIDUseCase struct {
	DoctorRepository entity.DoctorRepository
}

func NewGetDoctorByIDUseCase(doctorRepository entity.DoctorRepository) *GetDoctorByIDUseCase {
	return &GetDoctorByIDUseCase{DoctorRepository: doctorRepository}
}

func (u *GetDoctorByIDUseCase) Execute(input GetDoctorByIDInputDto) (*GetDoctorByIDOutputDto, error) {
	doctor, err := u.DoctorRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetDoctorByIDOutputDto{
		ID:          doctor.ID,
		Name:        doctor.Name,
		BI:          doctor.BI,
		Specialty:   doctor.Specialty,
		Experience:  doctor.Experience,
		Email:       doctor.Email,
		PhoneNumber: doctor.PhoneNumber,
	}, nil
}
