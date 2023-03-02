package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idazanggara/go-api-with-gin/model"
)

// CreateCar
func CreateCar(ctx *gin.Context) {
	var newCar model.Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(model.CarList)+1)
	model.CarList = append(model.CarList, newCar)

	ctx.JSON(http.StatusOK, gin.H{
		"message": newCar,
	})
}

// GetAllCar
func GetAllCar(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"cars": model.CarList,
	})
}

// UpdateCar
func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")

	condition := false

	var updateCar model.Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range model.CarList {
		if carID == car.CarID {
			condition = true
			model.CarList[i] = updateCar
			model.CarList[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v has been successfully updated", carID),
	})

}

// GetCar
func GetCarById(ctx *gin.Context) {
	carID := ctx.Param("carID")

	condition := false

	var carData model.Car

	for i, car := range model.CarList {
		if carID == car.CarID {
			condition = true
			carData = model.CarList[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})

}

// DeleteCar
func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")

	condition := false

	var carIndex int

	for i, car := range model.CarList {
		if carID == car.CarID {
			condition = true
			carIndex = i
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	model.CarList[carIndex] = model.CarList[len(model.CarList)-1]
	model.CarList[len(model.CarList)-1] = model.Car{}
	model.CarList = model.CarList[:len(model.CarList)-1] // Truncate slice

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v has been successfully delete", carID),
	})

}
