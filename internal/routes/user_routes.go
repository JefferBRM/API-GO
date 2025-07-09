package routes

import (
    "api-go/internal/controllers"
    "net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    // Configurar CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Middleware de logging
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Servir archivos estáticos
    r.Static("/static", "./static")
    r.StaticFile("/", "./static/index.html")

    // Rutas de API
    api := r.Group("/api")
    {
        // Información de la API
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

        // Rutas de usuarios
        users := api.Group("/users")
        {
            users.GET("/", controllers.GetAllUsers)
            users.GET("/:id", controllers.GetUserByID)
            users.POST("/", controllers.CreateUser)
            users.PUT("/:id", controllers.UpdateUser)
            users.DELETE("/:id", controllers.DeleteUser)
        }
    }

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":    "OK",
            "timestamp": gin.H{},
        })
    })

    // Middleware para rutas no encontradas
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "message": "Endpoint no encontrado",
        })
    })
}