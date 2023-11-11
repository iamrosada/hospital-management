package usecase

import (
	"github.com/iamrosada/hospital-management/doctor-service/internal/doctor/entity"
)

type DeleteDoctorInputDto struct {
	ID string `json:"id"`
}

type DeleteDoctorOutputDto struct {
	ID string `json:"id"`
}

type DeleteDoctorUseCase struct {
	DoctorRepository entity.DoctorRepository
}

func NewDeleteDoctorUseCase(DoctorRepository entity.DoctorRepository) *DeleteDoctorUseCase {
	return &DeleteDoctorUseCase{DoctorRepository: DoctorRepository}
}

func (u *DeleteDoctorUseCase) Execute(input DeleteDoctorInputDto) (*DeleteDoctorOutputDto, error) {
	err := u.DoctorRepository.DeleteByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteDoctorOutputDto{ID: input.ID}, nil
}
