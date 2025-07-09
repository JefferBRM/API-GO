package main

import (
    "log"
    "api-go/internal/config"
    "api-go/internal/routes"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar variables de entorno
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Inicializar base de datos
    if err := config.InitDB(); err != nil {
        log.Fatal("Error conectando a la base de datos:", err)
    }
    defer config.CloseDB()

    // Configurar Gin
    if os.Getenv("GIN_MODE") == "release" {
        gin.SetMode(gin.ReleaseMode)
    }

    // Crear router
    r := gin.Default()

    // Configurar rutas
    routes.SetupRoutes(r)

    // Obtener puerto
    port := os.Getenv("PORT")
    if port == "" {
        port = "3001"
    }

    log.Printf(">> Servidor corriendo en puerto %s", port)
    log.Printf(">> API disponible en: http://localhost:%s", port)
    log.Println(">> Endpoints principales:")
    log.Println("   GET    /api/users          - Obtener todos los usuarios")
    log.Println("   GET    /api/users/:id      - Obtener usuario por ID")
    log.Println("   POST   /api/users          - Crear nuevo usuario")
    log.Println("   PUT    /api/users/:id      - Actualizar usuario")
    log.Println("   DELETE /api/users/:id      - Eliminar usuario")

    if err := r.Run(":" + port); err != nil {
        log.Fatal("Error iniciando servidor:", err)
    }
}