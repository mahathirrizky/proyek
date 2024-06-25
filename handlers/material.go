package handlers

import (
	"net/http"
	"proyek/helper"
	"proyek/material"
	"strconv"

	"github.com/gin-gonic/gin"
)

type materialHandler struct {
	materialService material.Service
}

func NewMaterialHandler(materialService material.Service) *materialHandler {
	return &materialHandler{materialService}
}

func (h *materialHandler) CreateMaterial(c *gin.Context) {
	var input material.CreateMaterialInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newMaterial, err := h.materialService.CreateMaterial(input)
	if err != nil {
		response := helper.APIResponse("Failed to create material", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	var inputStok material.CreateStokInput

	inputStok.IDMaterial = newMaterial.IdMaterial
	inputStok.IDProyek = input.IDProyek
	inputStok.Jumlah = input.Jumlah

	_, err = h.materialService.CreateStok(inputStok)
	if err != nil {
		response := helper.APIResponse("Failed to create stok", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}


	response := helper.APIResponse("Material created successfully", http.StatusCreated, "success", material.FormatMaterial(newMaterial, input.Jumlah))
	c.JSON(http.StatusCreated, response)
}


func (h *materialHandler) GetAllMaterials(c *gin.Context) {
	materials, err := h.materialService.GetAllMaterials()
	if err != nil {
		response := helper.APIResponse("Failed to fetch materials", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Materials fetched successfully", http.StatusOK, "success", materials)
	c.JSON(http.StatusOK, response)
}


func (h *materialHandler) GetMaterialByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Invalid material ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	material, err := h.materialService.GetMaterialByID(id)
	if err != nil {
		response := helper.APIResponse("Material not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Material found", http.StatusOK, "success", material)
	c.JSON(http.StatusOK, response)
}

func (h *materialHandler) UpdateMaterial(c *gin.Context) {
	var input material.UpdateMaterialInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedMaterial, err := h.materialService.UpdateMaterial(id, input)
	if err != nil {
		response := helper.APIResponse("Failed to update material", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Material updated successfully", http.StatusOK, "success", material.FormatMaterial(updatedMaterial, input.Jumlah))
	c.JSON(http.StatusOK, response)
}