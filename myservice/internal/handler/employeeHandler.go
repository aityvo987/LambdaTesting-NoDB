package handler

import (
	"fmt"
	"net/http"

	entity "lambda-test/internal/entity"
	usecase "lambda-test/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	GetEmployees(c *gin.Context)
}

func CreateEmployee(c *gin.Context) {
	body := entity.CreateEmployeeRequest{}
	if err := c.BindJSON(&body); err != nil {
		log.Printf("Create Employee - ERROR Reading Body: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	resp, err := usecase.CreateEmployee(body)
	if err != nil {
		log.Printf("Create Employee - ERROR calling to Usecase: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot create new employee: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, resp)
}
