package utils

import (
    "github.com/gin-gonic/gin"
)

type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Count   int         `json:"count,omitempty"`
}

func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
    response := Response{
        Success: true,
        Message: message,
        Data:    data,
    }
    c.JSON(status, response)
}

func SuccessResponseWithCount(c *gin.Context, status int, message string, data interface{}, count int) {
    response := Response{
        Success: true,
        Message: message,
        Data:    data,
        Count:   count,
    }
    c.JSON(status, response)
}

func ErrorResponse(c *gin.Context, status int, message string, err error) {
    response := Response{
        Success: false,
        Message: message,
    }

    if err != nil {
        response.Error = err.Error()
    }

    c.JSON(status, response)
}