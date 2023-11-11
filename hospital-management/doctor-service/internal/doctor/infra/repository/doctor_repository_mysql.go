// repository/doctor_repository_mysql.go

package repository

import (
	"database/sql"

	"github.com/iamrosada/hospital-management/doctor-service/internal/doctor/entity"
)

type DoctorRepositoryMysql struct {
	DB *sql.DB
}

func NewDoctorRepositoryMysql(db *sql.DB) *DoctorRepositoryMysql {
	return &DoctorRepositoryMysql{DB: db}
}

func (r *DoctorRepositoryMysql) Create(doctor *entity.Doctor) error {
	_, err := r.DB.Exec("INSERT INTO doctors (id, name, bi, specialty, experience, email, phone_number) VALUES (?, ?, ?, ?, ?, ?, ?)",
		doctor.ID, doctor.Name, doctor.BI, doctor.Specialty, doctor.Experience, doctor.Email, doctor.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *DoctorRepositoryMysql) FindAll() ([]*entity.Doctor, error) {
	rows, err := r.DB.Query("SELECT id, name, bi, specialty, experience, email, phone_number FROM doctors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []*entity.Doctor
	for rows.Next() {
		var doctor entity.Doctor
		err = rows.Scan(&doctor.ID, &doctor.Name, &doctor.BI, &doctor.Specialty, &doctor.Experience, &doctor.Email, &doctor.PhoneNumber)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, &doctor)
	}
	return doctors, nil
}

func (r *DoctorRepositoryMysql) Update(doctor *entity.Doctor) error {
	_, err := r.DB.Exec("UPDATE doctors SET name=?, bi=?, specialty=?, experience=?, email=?, phone_number=? WHERE id=?",
		doctor.Name, doctor.BI, doctor.Specialty, doctor.Experience, doctor.Email, doctor.PhoneNumber, doctor.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *DoctorRepositoryMysql) DeleteByID(id string) error {
	_, err := r.DB.Exec("DELETE FROM doctors WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *DoctorRepositoryMysql) GetByID(id string) (*entity.Doctor, error) {
	var doctor entity.Doctor
	err := r.DB.QueryRow("SELECT id, name, bi, specialty, experience, email, phone_number FROM doctors WHERE id=?", id).
		Scan(&doctor.ID, &doctor.Name, &doctor.BI, &doctor.Specialty, &doctor.Experience, &doctor.Email, &doctor.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}
