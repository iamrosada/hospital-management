// hospital-management/doctor-service/cmd/doctor-service/main.go

package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iamrosada/hospital-management/doctor-service/api"
	"github.com/iamrosada/hospital-management/doctor-service/internal/doctor/infra/repository"
	usecase "github.com/iamrosada/hospital-management/doctor-service/internal/doctor/usecases"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/doctors")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	doctorRepository := repository.NewDoctorRepositoryMysql(db)
	createDoctorUsecase := usecase.NewCreateDoctorUseCase(doctorRepository)
	listDoctorsUsecase := usecase.NewGetAllDoctorsUseCase(doctorRepository)
	deleteDoctorUsecase := usecase.NewDeleteDoctorUseCase(doctorRepository)
	getDoctorByIDUsecase := usecase.NewGetDoctorByIDUseCase(doctorRepository)
	updateDoctorUsecase := usecase.NewUpdateDoctorUseCase(doctorRepository)

	doctorHandlers := api.NewDoctorHandlers(createDoctorUsecase, listDoctorsUsecase, deleteDoctorUsecase, getDoctorByIDUsecase, updateDoctorUsecase)

	r := chi.NewRouter()
	r.Post("/doctors", doctorHandlers.CreateDoctorHandler)
	r.Get("/doctors", doctorHandlers.ListDoctorsHandler)
	r.Get("/doctors/{id}", doctorHandlers.GetDoctorByIDHandler)
	r.Put("/doctors/{id}", doctorHandlers.UpdateDoctorHandler)
	r.Delete("/doctors/{id}", doctorHandlers.DeleteDoctorHandler)

	http.ListenAndServe(":8000", r)
}
