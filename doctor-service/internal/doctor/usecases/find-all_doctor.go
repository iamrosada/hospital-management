package usecase

import (
	"github.com/iamrosada/hospital-management/doctor-service/internal/doctor/entity"
)

type GetAllDoctorsOutputDto struct {
	Doctors []*entity.Doctor `json:"doctors"`
}

type GetAllDoctorsUseCase struct {
	DoctorRepository entity.DoctorRepository
}

func NewGetAllDoctorsUseCase(DoctorRepository entity.DoctorRepository) *GetAllDoctorsUseCase {
	return &GetAllDoctorsUseCase{DoctorRepository: DoctorRepository}
}

func (u *GetAllDoctorsUseCase) Execute() (*GetAllDoctorsOutputDto, error) {
	doctors, err := u.DoctorRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &GetAllDoctorsOutputDto{Doctors: doctors}, nil
}
