-- Script de creación de la base de datos para API de Gestión de Usuarios
-- PostgreSQL

-- Crear base de datos (ejecutar como superusuario)
CREATE DATABASE mini_api_usuarios;

-- Conectar a la base de datos
\c mini_api_usuarios;

-- Crear tabla de usuarios
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    telefono VARCHAR(20),
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear trigger para actualizar fecha_actualizacion automáticamente
CREATE OR REPLACE FUNCTION actualizar_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fecha_actualizacion = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_actualizar_fecha
    BEFORE UPDATE ON usuarios
    FOR EACH ROW
    EXECUTE FUNCTION actualizar_fecha_modificacion();

-- Insertar datos de ejemplo (opcional)
INSERT INTO usuarios (nombre, email, telefono) VALUES
('Jefferson Guerrero', 'jeffersonguerrero332@gmail.com', '+57 322 896 7226'),
('Maria Garcia', 'maria@example.com', '+57 301 987 6543'),
('Carlos Lopez', 'carlos@example.com', '+57 302 456 7890');
