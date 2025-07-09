package routes

import (
    "api-go/internal/controllers"
    "net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.Static("/static", "./static")
    r.StaticFile("/", "./static/index.html")

    api := r.Group("/api")
    {
        api.GET("/", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "API de Gestión de Usuarios",
                "version": "1.0.0",
                "endpoints": gin.H{
                    "users":       "/api/users",
                    "healthcheck": "/health",
                    "dashboard":   "/",
                },
            })
        })

        users := api.Group("/users")
        {
            users.GET("/", controllers.GetAllUsers)
            users.GET("/:id", controllers.GetUserByID)
            users.POST("/", controllers.CreateUser)
            users.PUT("/:id", controllers.UpdateUser)
            users.DELETE("/:id", controllers.DeleteUser)
        }
    }

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":    "OK",
            "timestamp": gin.H{},
        })
    })

    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "message": "Endpoint no encontrado",
        })
    })
}