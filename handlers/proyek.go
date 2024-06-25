package handlers

import (
    "net/http"
    "proyek/helper"
    "proyek/proyek"


    "github.com/gin-gonic/gin"
)

type proyekHandler struct {
    service proyek.Service
}

func NewProyekHandler(service proyek.Service) *proyekHandler {
    return &proyekHandler{service}
}

func (h *proyekHandler) CreateProyek(c *gin.Context) {
    var input proyek.CreateProyekInput
    if err := c.ShouldBindJSON(&input); err != nil {
        errors := helper.FormatValidationError(err)
        errorMessage := gin.H{"errors": errors}
        response := helper.APIResponse("Create proyek failed", http.StatusUnprocessableEntity, "error", errorMessage)
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }

    newProyek, err := h.service.CreateProyek(input)
    if err != nil {
        response := helper.APIResponse("Create proyek failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    formatter := proyek.FormatProyek(newProyek)
    response := helper.APIResponse("Proyek created", http.StatusCreated, "success", formatter)
    c.JSON(http.StatusCreated, response)
}

func (h *proyekHandler) UpdateProyek(c *gin.Context) {
    var inputID proyek.GetProyekDetailInput
    if err := c.ShouldBindUri(&inputID); err != nil {
        response := helper.APIResponse("Update proyek failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    var inputData proyek.UpdateProyekInput
    if err := c.ShouldBindJSON(&inputData); err != nil {
        errors := helper.FormatValidationError(err)
        errorMessage := gin.H{"errors": errors}
        response := helper.APIResponse("Update proyek failed", http.StatusUnprocessableEntity, "error", errorMessage)
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }

    updatedProyek, err := h.service.UpdateProyek(inputID, inputData)
    if err != nil {
        response := helper.APIResponse("Update proyek failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    formatter := proyek.FormatProyek(updatedProyek)
    response := helper.APIResponse("Proyek updated", http.StatusOK, "success", formatter)
    c.JSON(http.StatusOK, response)
}

func (h *proyekHandler) GetProyekByID(c *gin.Context) {
    var input proyek.GetProyekDetailInput
    if err := c.ShouldBindUri(&input); err != nil {
        response := helper.APIResponse("Get proyek detail failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    proyekDetail, err := h.service.GetProyekByID(input)
    if err != nil {
        response := helper.APIResponse("Get proyek detail failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    formatter := proyek.FormatProyek(proyekDetail)
    response := helper.APIResponse("Proyek detail", http.StatusOK, "success", formatter)
    c.JSON(http.StatusOK, response)
}

func (h *proyekHandler) GetAllProyeks(c *gin.Context) {
    proyeks, err := h.service.GetAllProyeks()
    if err != nil {
        response := helper.APIResponse("Get proyeks failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    formatter := proyek.FormatProyeks(proyeks)
    response := helper.APIResponse("List of proyeks", http.StatusOK, "success", formatter)
    c.JSON(http.StatusOK, response)
}

func (h *proyekHandler) DeleteProyek(c *gin.Context) {
    var input proyek.GetProyekDetailInput
    if err := c.ShouldBindUri(&input); err != nil {
        response := helper.APIResponse("Delete proyek failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    err := h.service.DeleteProyek(input)
    if err != nil {
        response := helper.APIResponse("Delete proyek failed", http.StatusBadRequest, "error", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    response := helper.APIResponse("Proyek deleted", http.StatusOK, "success", nil)
    c.JSON(http.StatusOK, response)
}
