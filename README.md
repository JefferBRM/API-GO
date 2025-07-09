# API de Gestión de Usuarios - Go + PostgreSQL

API RESTful desarrollada en Go para la gestión de usuarios con operaciones CRUD completas, incluyendo un dashboard web para administración.

![image](https://github.com/user-attachments/assets/f86a8448-0208-4267-a8ff-23e07b298dbf)
![image](https://github.com/user-attachments/assets/11a60d50-daf7-4b3c-aae3-d59d15457163)

## Características

-  **CRUD Completo**: Crear, leer, actualizar y eliminar usuarios
-  **API RESTful**: Endpoints siguiendo estándares REST
-  **Dashboard Web**: Interfaz gráfica para administrar usuarios
-  **Validaciones**: Validación de datos de entrada robusta
-  **Manejo de Errores**: Respuestas consistentes y manejo de errores
-  **Base de Datos**: PostgreSQL con pool de conexiones optimizado
-  **Arquitectura Limpia**: Implementación de patrones MVC y Repository
-  **CORS**: Configurado para desarrollo y producción
-  **Logging**: Registro detallado de operaciones

## Tecnologías Utilizadas

- **Backend**: Go 1.24+ con Gin Web Framework
- **Base de Datos**: PostgreSQL 17+
- **Patrones**: MVC, Repository, Dependency Injection
- **Frontend**: HTML5, CSS3, JavaScript (Vanilla)

## Requisitos

- Go 1.24 o superior
- PostgreSQL 17 o superior
- Git

## Instalación

### 1. Clonar el repositorio
```bash
git clone https://github.com/JefferBRM/API-GO.git
cd api-go-users
```

### 2. Instalar dependencias
```bash
go mod download
```

### 3. Configurar base de datos
```bash
# Crear base de datos en PostgreSQL
psql -U postgres -c "CREATE DATABASE mini_api_usuarios;"

# Ejecutar script de creación de tablas
psql -U postgres -d mini_api_usuarios -f database.sql
```

### 4. Configurar variables de entorno
Crear archivo `.env` en la raíz del proyecto:
```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=mini_api_usuarios
DB_USER=postgres
DB_PASSWORD=****

PORT=3001
```

### 5. Ejecutar la aplicación
```bash
go run cmd/server/main.go
go run ./cmd/server
```

La aplicación estará disponible en:
- **API**: http://localhost:3001/api
- **Dashboard**: http://localhost:3001

## Documentación de la API

### Endpoints Principales

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/users/` | Obtener todos los usuarios |
| GET | `/api/users/:id` | Obtener usuario por ID |
| POST | `/api/users/` | Crear nuevo usuario |
| PUT | `/api/users/:id` | Actualizar usuario |
| DELETE | `/api/users/:id` | Eliminar usuario |
| GET | `/health` | Health check |

### Ejemplo de Uso

#### Crear Usuario
```bash
curl -X POST http://localhost:3001/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Pérez",
    "email": "juan@example.com",
    "telefono": "+57 300 1234567"
  }'
```

#### Obtener Todos los Usuarios

curl -X GET http://localhost:3001/api/users/


## Arquitectura del Proyecto

```
api-go/
├── cmd/
│   └── server/
│       └── main.go              # Punto de entrada
├── internal/
│   ├── config/
│   │   └── database.go          # Configuración de BD
│   ├── controllers/
│   │   └── user_controller.go   # Lógica de controladores
│   ├── models/
│   │   └── user_model.go        # Modelos y acceso a datos
│   └── routes/
│       └── user_routes.go       # Definición de rutas
├── pkg/
│   └── utils/
│       └── response.go          # Utilidades de respuesta
├── static/
│   └── index.html               # Dashboard web
├── database.sql                 # Script de creación de BD
├── .env                         # Variables de entorno
├── go.mod                       # Dependencias de Go
└── README.md
```

##  Patrones de Diseño Implementados

### 1. **MVC (Model-View-Controller)**
- **Modelos**: Manejo de datos y lógica de persistencia
- **Vistas**: Dashboard web y respuestas JSON
- **Controladores**: Lógica de negocio y manejo de requests

### 2. **Repository Pattern**
- Abstracción del acceso a datos
- Facilita testing y cambios de implementación
- Operaciones CRUD encapsuladas

### 3. **Dependency Injection**
- Inyección de dependencias de base de datos
- Facilita testing y desacoplamiento
- Configuración centralizada

### 4. **Singleton Pattern**
- Conexión única a base de datos
- Gestión eficiente de recursos
- Pool de conexiones optimizado

##  Dashboard Web

El proyecto incluye un dashboard web moderno con:
-  Lista de usuarios con paginación
-  Formulario de creación de usuarios
-  Edición en línea
-  Eliminación con confirmación
-  Diseño responsive
-  Validaciones client-side

##  Seguridad

- Validación de datos de entrada
- Prevención de SQL injection
- Sanitización de datos
- Manejo seguro de errores

## Autor

**Tu Nombre**
- GitHub: [@JefferBRM](https://github.com/tu-usuario)
- Email: jeffersonguerrero332@gmail.com

---
