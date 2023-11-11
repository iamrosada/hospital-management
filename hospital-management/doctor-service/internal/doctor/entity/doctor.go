package entity

import (
	"github.com/google/uuid"
)

type DoctorRepository interface {
	Create(doctor *Doctor) error
	FindAll() ([]*Doctor, error)
	Update(doctor *Doctor) Doctor
	Delete(id string) error
	GetByID(id string) (*Doctor, error)
}

// Doctor represents a doctor entity.
type Doctor struct {
	ID          string
	Name        string
	BI          string
	Specialty   string
	Experience  int
	Email       string
	PhoneNumber string
}

func NewDoctor(name, bi, specialty, email, phoneNumber string, experience int) *Doctor {
	return &Doctor{
		ID:          uuid.New().String(),
		Name:        name,
		BI:          bi,
		Specialty:   specialty,
		Experience:  experience,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}
