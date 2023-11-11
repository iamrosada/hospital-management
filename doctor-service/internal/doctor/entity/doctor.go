package entity

import (
	"errors"

	"github.com/google/uuid"
)

type DoctorRepository interface {
	Create(doctor *Doctor) error
	FindAll() ([]*Doctor, error)
	Update(doctor *Doctor) error
	DeleteByID(id string) error
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

func (d *Doctor) Update(name, bi, specialty, email, phoneNumber string, experience int) {
	d.Name = name
	d.BI = bi
	d.Specialty = specialty
	d.Experience = experience
	d.Email = email
	d.PhoneNumber = phoneNumber
}

type InMemoryDoctorRepository struct {
	doctors map[string]*Doctor
}

func NewInMemoryDoctorRepository() *InMemoryDoctorRepository {
	return &InMemoryDoctorRepository{
		doctors: make(map[string]*Doctor),
	}
}

func (r *InMemoryDoctorRepository) DeleteByID(id string) error {
	if _, exists := r.doctors[id]; !exists {
		return errors.New("doctor not found")
	}

	delete(r.doctors, id)
	return nil
}

func (r *InMemoryDoctorRepository) FindAll() ([]*Doctor, error) {
	var allDoctors []*Doctor
	for _, doctor := range r.doctors {
		allDoctors = append(allDoctors, doctor)
	}
	return allDoctors, nil
}
