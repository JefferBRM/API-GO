package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
    var err error
    
    // Construir string de conexión
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    // Conectar a la base de datos
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error abriendo conexión: %w", err)
    }

    // Verificar conexión
    if err = DB.Ping(); err != nil {
        return fmt.Errorf("error conectando a PostgreSQL: %w", err)
    }

    // Configurar pool de conexiones
    DB.SetMaxOpenConns(20)
    DB.SetMaxIdleConns(10)

    log.Println("Conexión exitosa a PostgreSQL")
    return nil
}

func CloseDB() {
    if DB != nil {
        DB.Close()
    }
}