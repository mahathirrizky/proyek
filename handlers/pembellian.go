package handlers

import (
	"net/http"
	"proyek/helper"
	"proyek/pembelian"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PembelianHandler struct {
	PembelianService pembelian.PembelianService
}

func NewPembelianHandler(ps pembelian.PembelianService) *PembelianHandler {
	return &PembelianHandler{
		PembelianService: ps,
	}
}

func (ph *PembelianHandler) CreatePembelian(c *gin.Context) {
	var input pembelian.CreatePembelianInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	pembelian, err := ph.PembelianService.CreatePembelian(input.IdMaterial, input.Jumlah, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to create pembelian", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("Pembelian created successfully", http.StatusCreated, "success", pembelian)
	c.JSON(http.StatusCreated, response)
}

func (ph *PembelianHandler) UpdatePembelian(c *gin.Context) {
	var input pembelian.UpdatePembelianInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid pembelian ID", http.StatusBadRequest, "error", nil))
		return
	}

	pembelian, err := ph.PembelianService.UpdatePembelian(id, input.Jumlah, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to update pembelian", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("Pembelian updated successfully", http.StatusOK, "success", pembelian)
	c.JSON(http.StatusOK, response)
}

func (ph *PembelianHandler) DeletePembelian(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid pembelian ID", http.StatusBadRequest, "error", nil))
		return
	}

	err = ph.PembelianService.DeletePembelian(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to delete pembelian", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("Pembelian deleted successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (ph *PembelianHandler) GetPembelianByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid pembelian ID", http.StatusBadRequest, "error", nil))
		return
	}

	pembelian, err := ph.PembelianService.GetPembelianByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch pembelian", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("Pembelian details", http.StatusOK, "success", pembelian)
	c.JSON(http.StatusOK, response)
}

func (ph *PembelianHandler) GetAllPembelian(c *gin.Context) {
	pembelians, err := ph.PembelianService.GetAllPembelian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch pembelians", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("List of pembelians", http.StatusOK, "success", pembelians)
	c.JSON(http.StatusOK, response)
}
