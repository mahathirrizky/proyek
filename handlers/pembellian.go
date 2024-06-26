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

	p, err := ph.PembelianService.CreatePembelian(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to create pembelian", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := pembelian.FormatPembelian(*p)
	response := helper.APIResponse("Pembelian created successfully", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (ph *PembelianHandler) UpdatePembelian(c *gin.Context) {
	var input pembelian.UpdatePembelianInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	var inputID pembelian.GetPembelianDetailInput
    if err := c.ShouldBindUri(&inputID); err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid pembelian ID", http.StatusBadRequest, "error", nil))
		return
	}

	p, err := ph.PembelianService.UpdatePembelian(inputID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to update pembelian", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := pembelian.FormatPembelian(*p)

	response := helper.APIResponse("Pembelian updated successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (ph *PembelianHandler) DeletePembelian(c *gin.Context) {
	var inputID pembelian.GetPembelianDetailInput
    if err := c.ShouldBindUri(&inputID); err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid pembelian ID", http.StatusBadRequest, "error", nil))
		return
	}

	err := ph.PembelianService.DeletePembelian(inputID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to delete pembelian", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("Pembelian deleted successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (ph *PembelianHandler) GetPembelianByID(c *gin.Context) {
	var inputID pembelian.GetPembelianDetailInput
	if err := c.ShouldBindUri(&inputID); err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid pembelian ID", http.StatusBadRequest, "error", nil))
		return
	}

	p, err := ph.PembelianService.GetPembelianByID(inputID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch pembelian", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := pembelian.FormatPembelian(*p)

	response := helper.APIResponse("Pembelian details", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (ph *PembelianHandler) GetAllPembelian(c *gin.Context) {
	p, err := ph.PembelianService.GetAllPembelian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch pembelian list", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := pembelian.FormatPembelianList(p)

	response := helper.APIResponse("List of pembelian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (ph *PembelianHandler) GetPembelianByProyekID(c *gin.Context) {
	proyekID, err := strconv.Atoi(c.Param("proyek_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid proyek ID", http.StatusBadRequest, "error", nil))
		return
	}

	p, err := ph.PembelianService.GetPembelianByProyekID(proyekID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch pembelian list", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := pembelian.FormatPembelianList(p)

	response := helper.APIResponse("List of pembelian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}