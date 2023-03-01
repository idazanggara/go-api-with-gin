package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarList = []Car{}

// CreateCar
func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarList)+1)
	CarList = append(CarList, newCar)

	ctx.JSON(http.StatusOK, gin.H{
		"message": newCar,
	})
}

// GetAllCar
func GetAllCar(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"cars": CarList,
	})
}

// UpdateCar
func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")

	condition := false

	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarList {
		if carID == car.CarID {
			condition = true
			CarList[i] = updateCar
			CarList[i].CarID = carID
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

	var carData Car

	for i, car := range CarList {
		if carID == car.CarID {
			condition = true
			carData = CarList[i]
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

	for i, car := range CarList {
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

	CarList[carIndex] = CarList[len(CarList)-1]
	CarList[len(CarList)-1] = Car{}
	CarList = CarList[:len(CarList)-1] // Truncate slice

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v has been successfully delete", carID),
	})

}

func main() {
	routerEngine := gin.Default()

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})

	routerEngine.POST("/cars", CreateCar)
	routerEngine.GET("/cars", GetAllCar)
	routerEngine.PUT("/cars/:carID", UpdateCar)
	routerEngine.GET("/cars/:carID", GetCarById)
	routerEngine.DELETE("/cars/:carID", DeleteCar)

	if err := routerEngine.Run(); err != nil {
		panic(err)
	}
	// secara default menggunakan port :8080
}
