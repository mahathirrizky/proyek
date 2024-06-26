package handlers

import (
	"net/http"
	"proyek/helper"
	"proyek/permintaan"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermintaanHandler struct {
	PermintaanService permintaan.PermintaanService
}

func NewPermintaanHandler(ps permintaan.PermintaanService) *PermintaanHandler {
	return &PermintaanHandler{
		PermintaanService: ps,
	}
}

func (ph *PermintaanHandler) CreatePermintaan(c *gin.Context) {
	var input permintaan.CreatePermintaanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	p, err := ph.PermintaanService.CreatePermintaan(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to create permintaan", http.StatusInternalServerError, "error", nil))
		return
	}

	formatter := permintaan.FormatPermintaan(p)
	response := helper.APIResponse("Permintaan created successfully", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (ph *PermintaanHandler) UpdatePermintaan(c *gin.Context) {
	var input permintaan.UpdatePermintaanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid permintaan ID", http.StatusBadRequest, "error", nil))
		return
	}

	p, err := ph.PermintaanService.UpdatePermintaan(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to update permintaan", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := permintaan.FormatPermintaan(p)
	response := helper.APIResponse("Permintaan updated successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (ph *PermintaanHandler) GetPermintaanByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid permintaan ID", http.StatusBadRequest, "error", nil))
		return
	}

	p, err := ph.PermintaanService.GetPermintaanByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to get permintaan", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := permintaan.FormatPermintaan(p)
	response := helper.APIResponse("Permintaan fetched successfully", http.StatusOK, "success",formatter)
	c.JSON(http.StatusOK, response)
}

func (ph *PermintaanHandler) GetPermintaanByProyekID(c *gin.Context) {
	proyekID, err := strconv.Atoi(c.Param("proyek_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid proyek ID", http.StatusBadRequest, "error", nil))
		return
	}

	p, err := ph.PermintaanService.GetPermintaanByProyekID(proyekID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to get permintaan by proyek ID", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := permintaan.FormatPermintaanList(p)

	response := helper.APIResponse("Permintaan fetched successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (ph *PermintaanHandler) GetAllPermintaan(c *gin.Context) {
  p, err := ph.PermintaanService.GetAllPermintaan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to get all permintaan", http.StatusInternalServerError, "error", nil))
		return
	}
	formatter := permintaan.FormatPermintaanList(p)

	response := helper.APIResponse("All permintaan fetched successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
