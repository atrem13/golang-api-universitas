package handler

import (
	"net/http"

	"github.com/atrem13/golang-api-universitas/helper"
	"github.com/atrem13/golang-api-universitas/prodi"
	"github.com/gin-gonic/gin"
)

type prodiHandler struct {
	service prodi.Service
}

func NewProdiHandler(service prodi.Service) *prodiHandler {
	return &prodiHandler{service: service}
}

func (h *prodiHandler) GetProdis(c *gin.Context) {
	prodis, err := h.service.GetProdis()
	if err != nil {
		response := helper.APIResponse("Error, table prodi is empty", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := prodi.FormatProdis(prodis)
	response := helper.APIResponse("List of prodis", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *prodiHandler) GetProdi(c *gin.Context) {
	var input prodi.GetProdiDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error, fail to get params from uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	prodiDetail, err := h.service.GetProdiByID(input)
	if err != nil {
		response := helper.APIResponse("Error, fail to get data with ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := prodi.FormatProdi(prodiDetail)
	response := helper.APIResponse("Detail of prodi", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *prodiHandler) CreateProdi(c *gin.Context) {
	var input prodi.CreateProdiInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Error, fail create prodi", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProdi, err := h.service.CreateProdi(input)
	if err != nil {
		response := helper.APIResponse("Error, fail to create prodi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := prodi.FormatProdi(newProdi)
	response := helper.APIResponse("Success to create new prodi", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *prodiHandler) UpdateProdi(c *gin.Context) {
	var inputID prodi.GetProdiDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update prodi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData prodi.CreateProdiInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update prodi", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateProdi, err := h.service.UpdateProdi(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := prodi.FormatProdi(updateProdi)
	response := helper.APIResponse("Success to update prodi", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *prodiHandler) DeleteProdi(c *gin.Context) {
	var inputID prodi.GetProdiDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete prodi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateProdi, err := h.service.DeleteProdi(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := prodi.FormatProdi(updateProdi)
	response := helper.APIResponse("Success to delete prodi", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
