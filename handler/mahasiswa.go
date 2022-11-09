package handler

import (
	"net/http"

	"github.com/atrem13/golang-api-universitas/helper"
	"github.com/atrem13/golang-api-universitas/mahasiswa"
	"github.com/gin-gonic/gin"
)

type mahasiswaHandler struct {
	service mahasiswa.Service
}

func NewMahasiswaHandler(service mahasiswa.Service) *mahasiswaHandler {
	return &mahasiswaHandler{service: service}
}

func (h *mahasiswaHandler) GetMahasiswas(c *gin.Context) {
	mahasiswas, err := h.service.GetMahasiswas()
	if err != nil {
		response := helper.APIResponse("Error, table mahasiswa is empty", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	message := "List of mahasiswa"
	if len(mahasiswas) < 1 {
		message = "Table mahasiswa is empty"
	}
	formatter := mahasiswa.FormatMahasiswas(mahasiswas)
	response := helper.APIResponse(message, http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *mahasiswaHandler) GetMahasiswa(c *gin.Context) {
	var input mahasiswa.GetMahasiswaDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error, fail to get params from uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mahasiswaDetail, err := h.service.GetMahasiswaByID(input)
	if err != nil {
		response := helper.APIResponse("Error, fail to get data with ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := mahasiswa.FormatMahasiswa(mahasiswaDetail)
	response := helper.APIResponse("Detail of mahasiswa", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *mahasiswaHandler) CreateMahasiswa(c *gin.Context) {
	var input mahasiswa.CreateMahasiswaInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Error, fail create mahasiswax", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newMahasiswa, err := h.service.CreateMahasiswa(input)
	if err != nil {
		response := helper.APIResponse("Error, fail to create mahasiswax", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := mahasiswa.FormatMahasiswa(newMahasiswa)
	response := helper.APIResponse("Success to create new mahasiswa", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *mahasiswaHandler) UpdateMahasiswa(c *gin.Context) {
	var inputID mahasiswa.GetMahasiswaDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData mahasiswa.CreateMahasiswaInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update mahasiswa", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateMahasiswa, err := h.service.UpdateMahasiswa(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := mahasiswa.FormatMahasiswa(updateMahasiswa)
	response := helper.APIResponse("Success to update mahasiswa", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *mahasiswaHandler) DeleteMahasiswa(c *gin.Context) {
	var inputID mahasiswa.GetMahasiswaDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteMahasiswa, err := h.service.DeleteMahasiswa(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := mahasiswa.FormatMahasiswa(deleteMahasiswa)
	response := helper.APIResponse("Success to delete mahasiswa", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
