package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type information struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Completed   bool   `json:"completed"`
}

var informations = []information{
	{Name: "dilay", Email: "dilay@zeelso.com", PhoneNumber: "12345", Completed: false},
	{Name: "kadir", Email: "kadir@zeelso.com", PhoneNumber: "23456", Completed: false},
	{Name: "yasir", Email: "yasir@zeelso.com", PhoneNumber: "34567", Completed: false},
	{Name: "yusuf", Email: "yusuf@zeelso.com", PhoneNumber: "45678", Completed: false},
}

func getInformations(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, informations)
}

func addInformation(context *gin.Context) {
	var newInformation information

	if err := context.BindJSON(&newInformation); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	informations = append(informations, newInformation)

	context.IndentedJSON(http.StatusCreated, newInformation)
}

func getInformation(context *gin.Context) {
	id := context.Param("id")
	information, err := getInformationByName(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Information not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, information)

}

func toggleInformationStatus(context *gin.Context) {
	id := context.Param("id")
	information, err := getInformationByName(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Information not found"})
		return
	}

	information.Completed = !information.Completed

	context.IndentedJSON(http.StatusOK, information)

}

func getInformationByName(name string) (*information, error) {
	for i, info := range informations {
		if info.Name == name {
			return &informations[i], nil
		}
	}

	return nil, errors.New("Information not found")
}

func main() {
	router := gin.Default()
	router.GET("/informations", getInformations)
	router.GET("/informations/:id", getInformation)
	router.PATCH("/informations/:id", toggleInformationStatus)
	router.POST("/informations", addInformation)
	router.Run("localhost:9090")
}
