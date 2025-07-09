package models

import (
    "database/sql"
    "fmt"
    "api-go/internal/config"
    "time"
)

type User struct {
    ID                 int       `json:"id"`
    Nombre             string    `json:"nombre"`
    Email              string    `json:"email"`
    Telefono           *string   `json:"telefono"`
    FechaCreacion      time.Time `json:"fecha_creacion"`
    FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

type UserInput struct {
    Nombre   string  `json:"nombre" binding:"required"`
    Email    string  `json:"email" binding:"required,email"`
    Telefono *string `json:"telefono"`
}

func GetAllUsers() ([]User, error) {
    query := `
        SELECT id, nombre, email, telefono, fecha_creacion, fecha_actualizacion 
        FROM usuarios 
        ORDER BY fecha_creacion DESC
    `
    
    rows, err := config.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error ejecutando consulta: %w", err)
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(
            &user.ID,
            &user.Nombre,
            &user.Email,
            &user.Telefono,
            &user.FechaCreacion,
            &user.FechaActualizacion,
        )
        if err != nil {
            return nil, fmt.Errorf("error escaneando fila: %w", err)
        }
        users = append(users, user)
    }

    return users, nil
}

func GetUserByID(id int) (*User, error) {
    query := `
        SELECT id, nombre, email, telefono, fecha_creacion, fecha_actualizacion 
        FROM usuarios 
        WHERE id = $1
    `
    
    var user User
    err := config.DB.QueryRow(query, id).Scan(
        &user.ID,
        &user.Nombre,
        &user.Email,
        &user.Telefono,
        &user.FechaCreacion,
        &user.FechaActualizacion,
    )
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("error obteniendo usuario: %w", err)
    }

    return &user, nil
}

func CreateUser(input UserInput) (*User, error) {
    query := `
        INSERT INTO usuarios (nombre, email, telefono) 
        VALUES ($1, $2, $3) 
        RETURNING id, nombre, email, telefono, fecha_creacion, fecha_actualizacion
    `
    
    var user User
    err := config.DB.QueryRow(query, input.Nombre, input.Email, input.Telefono).Scan(
        &user.ID,
        &user.Nombre,
        &user.Email,
        &user.Telefono,
        &user.FechaCreacion,
        &user.FechaActualizacion,
    )
    
    if err != nil {
        if err.Error() == `pq: duplicate key value violates unique constraint "usuarios_email_key"` {
            return nil, fmt.Errorf("el email ya está registrado")
        }
        return nil, fmt.Errorf("error creando usuario: %w", err)
    }

    return &user, nil
}

func UpdateUser(id int, input UserInput) (*User, error) {
    query := `
        UPDATE usuarios 
        SET nombre = $1, email = $2, telefono = $3 
        WHERE id = $4 
        RETURNING id, nombre, email, telefono, fecha_creacion, fecha_actualizacion
    `
    
    var user User
    err := config.DB.QueryRow(query, input.Nombre, input.Email, input.Telefono, id).Scan(
        &user.ID,
        &user.Nombre,
        &user.Email,
        &user.Telefono,
        &user.FechaCreacion,
        &user.FechaActualizacion,
    )
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        if err.Error() == `pq: duplicate key value violates unique constraint "usuarios_email_key"` {
            return nil, fmt.Errorf("el email ya está registrado")
        }
        return nil, fmt.Errorf("error actualizando usuario: %w", err)
    }

    return &user, nil
}

func DeleteUser(id int) (bool, error) {
    query := `DELETE FROM usuarios WHERE id = $1`
    
    result, err := config.DB.Exec(query, id)
    if err != nil {
        return false, fmt.Errorf("error eliminando usuario: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false, fmt.Errorf("error obteniendo filas afectadas: %w", err)
    }

    return rowsAffected > 0, nil
}