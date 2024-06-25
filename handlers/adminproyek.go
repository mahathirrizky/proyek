// file: handlers/adminproyek_handler.go

package handlers

import (
	"net/http"
	"proyek/adminproyek"
	"proyek/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminProyekHandler struct {
    AdminProyekService adminproyek.AdminProyekService

}

func NewAdminProyekHandler(adminProyekService adminproyek.AdminProyekService) *AdminProyekHandler {
    return &AdminProyekHandler{adminProyekService}
}

func (ah *AdminProyekHandler) CreateAdminProyek(c *gin.Context) {
    var input adminproyek.CreateAdminProyekInput
    if err := c.ShouldBindJSON(&input); err != nil {
        errors := helper.FormatValidationError(err)
        c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
        return
    }

    adminProyek, err := ah.AdminProyekService.CreateAdminProyek(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to create admin proyek", http.StatusInternalServerError, "error", nil))
        return
    }

    response := helper.APIResponse("Admin proyek created successfully", http.StatusCreated, "success", adminProyek)
    c.JSON(http.StatusCreated, response)
}

func (ah *AdminProyekHandler) UpdateAdminProyek(c *gin.Context) {
    var input adminproyek.UpdateAdminProyekInput
    if err := c.ShouldBindJSON(&input); err != nil {
        errors := helper.FormatValidationError(err)
        c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
        return
    }

    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid admin proyek ID", http.StatusBadRequest, "error", nil))
        return
    }

    adminProyek, err := ah.AdminProyekService.UpdateAdminProyek(id, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to update admin proyek", http.StatusInternalServerError, "error", nil))
        return
    }

    response := helper.APIResponse("Admin proyek updated successfully", http.StatusOK, "success", adminProyek)
    c.JSON(http.StatusOK, response)
}

func (ah *AdminProyekHandler) DeleteAdminProyek(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid admin proyek ID", http.StatusBadRequest, "error", nil))
        return
    }

    if err := ah.AdminProyekService.DeleteAdminProyek(id); err != nil {
        c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to delete admin proyek", http.StatusInternalServerError, "error", nil))
        return
    }

    response := helper.APIResponse("Admin proyek deleted successfully", http.StatusOK, "success", nil)
    c.JSON(http.StatusOK, response)
}

func (ah *AdminProyekHandler) GetAdminProyekByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid admin proyek ID", http.StatusBadRequest, "error", nil))
        return
    }

    adminProyek, err := ah.AdminProyekService.GetAdminProyekByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, helper.APIResponse("Admin proyek not found", http.StatusInternalServerError, "error", nil))
        return
    }

    response := helper.APIResponse("Admin proyek details", http.StatusOK, "success", adminProyek)
    c.JSON(http.StatusOK, response)
}

func (ah *AdminProyekHandler) GetAllAdminProyeks(c *gin.Context) {
    adminProyeks, err := ah.AdminProyekService.GetAllAdminProyeks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch admin proyeks", http.StatusInternalServerError, "error", nil))
        return
    }

    response := helper.APIResponse("List of admin proyeks", http.StatusOK, "success", adminProyeks)
    c.JSON(http.StatusOK, response)
}
