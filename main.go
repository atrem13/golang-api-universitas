package main

import (
	"fmt"
	"log"

	"github.com/atrem13/golang-api-universitas/handler"
	"github.com/atrem13/golang-api-universitas/mahasiswa"
	"github.com/atrem13/golang-api-universitas/prodi"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "golang_universitas"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// migrate database
	db.AutoMigrate(&prodi.Prodi{})
	db.AutoMigrate(&mahasiswa.Mahasiswa{})

	// initiate repository
	prodiRepository := prodi.NewRepository(db)
	mahasiswaRepository := mahasiswa.NewRepository(db)

	// initiate service
	prodiService := prodi.NewService(prodiRepository)
	mahasiswaService := mahasiswa.NewService(mahasiswaRepository)

	// initiate handler
	prodiHandler := handler.NewProdiHandler(prodiService)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaService)

	// initiate route
	router := gin.Default()

	api := router.Group("/api/v1")

	// prodi
	api.GET("/prodis", prodiHandler.GetProdis)
	api.GET("/prodis/:id", prodiHandler.GetProdi)
	api.POST("/prodis", prodiHandler.CreateProdi)
	api.PUT("/prodis/:id", prodiHandler.UpdateProdi)
	api.DELETE("/prodis/:id", prodiHandler.DeleteProdi)

	// mahasiswa
	api.GET("/mahasiswas", mahasiswaHandler.GetMahasiswas)
	api.GET("/mahasiswas/:id", mahasiswaHandler.GetMahasiswa)
	api.POST("/mahasiswas", mahasiswaHandler.CreateMahasiswa)
	api.PUT("/mahasiswas/:id", mahasiswaHandler.UpdateMahasiswa)
	api.DELETE("/mahasiswas/:id", mahasiswaHandler.DeleteMahasiswa)

	router.Run()

}
