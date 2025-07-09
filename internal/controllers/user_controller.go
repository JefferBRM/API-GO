package controllers

import (
    "api-go/internal/models"
    "api-go/pkg/utils"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
    users, err := models.GetAllUsers()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Error interno del servidor", err)
        return
    }

    utils.SuccessResponseWithCount(c, http.StatusOK, "Usuarios obtenidos exitosamente", users, len(users))
}

func GetUserByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", nil)
        return
    }

    user, err := models.GetUserByID(id)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Error interno del servidor", err)
        return
    }

    if user == nil {
        utils.ErrorResponse(c, http.StatusNotFound, "Usuario no encontrado", nil)
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Usuario encontrado", user)
}

func CreateUser(c *gin.Context) {
    var input models.UserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Datos inválidos", err)
        return
    }

    // Validaciones adicionales
    if strings.TrimSpace(input.Nombre) == "" {
        utils.ErrorResponse(c, http.StatusBadRequest, "El nombre es obligatorio", nil)
        return
    }

    if strings.TrimSpace(input.Email) == "" {
        utils.ErrorResponse(c, http.StatusBadRequest, "El email es obligatorio", nil)
        return
    }

    // Limpiar datos
    input.Nombre = strings.TrimSpace(input.Nombre)
    input.Email = strings.TrimSpace(strings.ToLower(input.Email))
    if input.Telefono != nil {
        telefono := strings.TrimSpace(*input.Telefono)
        if telefono == "" {
            input.Telefono = nil
        } else {
            input.Telefono = &telefono
        }
    }

    user, err := models.CreateUser(input)
    if err != nil {
        if strings.Contains(err.Error(), "email ya está registrado") {
            utils.ErrorResponse(c, http.StatusConflict, "El email ya está registrado", nil)
            return
        }
        utils.ErrorResponse(c, http.StatusInternalServerError, "Error interno del servidor", err)
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "Usuario creado exitosamente", user)
}

func UpdateUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", nil)
        return
    }

    var input models.UserInput
    if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Datos inválidos", err)
        return
    }

    // Validaciones adicionales
    if strings.TrimSpace(input.Nombre) == "" {
        utils.ErrorResponse(c, http.StatusBadRequest, "El nombre es obligatorio", nil)
        return
    }

    if strings.TrimSpace(input.Email) == "" {
        utils.ErrorResponse(c, http.StatusBadRequest, "El email es obligatorio", nil)
        return
    }

    // Limpiar datos
    input.Nombre = strings.TrimSpace(input.Nombre)
    input.Email = strings.TrimSpace(strings.ToLower(input.Email))
    if input.Telefono != nil {
        telefono := strings.TrimSpace(*input.Telefono)
        if telefono == "" {
            input.Telefono = nil
        } else {
            input.Telefono = &telefono
        }
    }

    user, err := models.UpdateUser(id, input)
    if err != nil {
        if strings.Contains(err.Error(), "email ya está registrado") {
            utils.ErrorResponse(c, http.StatusConflict, "El email ya está registrado", nil)
            return
        }
        utils.ErrorResponse(c, http.StatusInternalServerError, "Error interno del servidor", err)
        return
    }

    if user == nil {
        utils.ErrorResponse(c, http.StatusNotFound, "Usuario no encontrado", nil)
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Usuario actualizado exitosamente", user)
}

func DeleteUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", nil)
        return
    }

    deleted, err := models.DeleteUser(id)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Error interno del servidor", err)
        return
    }

    if !deleted {
        utils.ErrorResponse(c, http.StatusNotFound, "Usuario no encontrado", nil)
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Usuario eliminado exitosamente", nil)
}