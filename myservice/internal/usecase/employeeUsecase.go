package usecase

import (
	entity "lambda-test/internal/entity"
	model "lambda-test/internal/repository/model"
	"time"
)

type EmployeeUsecase interface {
	GetEmployees() ([]model.Employee, error)
	GetEmployee() ([]model.Employee, error)
	CreateEmployee(employee model.Employee) error
	UpdateEmployee(id int, employee model.Employee) error
	DeleteEmployee(id int) error
}

func CreateEmployee(req entity.CreateEmployeeRequest) (entity.PostEmployeeResponse, error) {
	resp := entity.PostEmployeeResponse{
		RequestID:   req.RequestID,
		CreatedTime: time.Now(),
		Data:        req.Data,
	}
	resp.Status = "Complete"

	return resp, nil
}
