package main

import (
	"fmt"
	"log"

	"github.com/atrem13/golang-api-universitas/handler"
	"github.com/atrem13/golang-api-universitas/prodi"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "golang_universitas"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	// migrate database
	db.AutoMigrate(&prodi.Prodi{})

	// initiate repository
	prodiRepository := prodi.NewRepository(db)

	// initiate service
	prodiService := prodi.NewService(prodiRepository)

	// initiate handler
	prodiHandler := handler.NewProdiHandler(prodiService)

	// initiate route
	router := gin.Default()

	api := router.Group("/api/v1")

	// prodi
	api.GET("/prodis", prodiHandler.GetProdis)
	api.GET("/prodis/:id", prodiHandler.GetProdi)
	api.POST("/prodis", prodiHandler.CreateProdi)
	api.PUT("/prodis/:id", prodiHandler.UpdateProdi)
	api.DELETE("/prodis/:id", prodiHandler.DeleteProdi)

	router.Run()

}
