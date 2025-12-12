# Go Template Project

This is a template project demonstrating a clean architecture pattern with a simple health check API. The project includes mocked database and Redis implementations for development and testing purposes.

## Architecture Overview

The project follows a layered architecture pattern with clear separation of concerns:

```
Controller → Manager → Service → Client
```

### Flow Description

1. **Controller Layer** (`internal/app/controller/`)
   - Handles HTTP requests and responses
   - Validates input and formats output
   - Calls the Manager layer

2. **Manager Layer** (`internal/app/manager/`)
   - Orchestrates business logic
   - Acts as an intermediary between Controller and Service
   - Can combine multiple service calls if needed

3. **Service Layer** (`internal/app/service/`)
   - Contains core business logic
   - Interacts with database and Redis
   - Calls external clients/APIs

4. **Client Layer** (`internal/app/clients/`)
   - Handles external API calls
   - Abstracts third-party integrations

## Project Structure

```
template/
├── api/
│   └── rest/
│       ├── server.go              # HTTP server setup
│       └── template/
│           ├── container.go       # Dependency injection container
│           └── apis.go            # Route registration
├── internal/
│   ├── app/
│   │   ├── controller/            # HTTP controllers
│   │   ├── manager/               # Business logic managers
│   │   ├── service/               # Core services
│   │   ├── clients/               # External API clients
│   │   └── dto/                   # Data transfer objects
│   └── pkg/
│       └── db/
│           ├── db.go              # Mocked database
│           └── redis.go           # Mocked Redis
```

## Initialization Flow

### 1. Server Setup (`api/rest/server.go`)

The `BuildServer()` function creates a new Gin engine with CORS middleware:

```go
func BuildServer() *gin.Engine {
    server := gin.New()
    server.Use(cors.New(cors.Config{
        AllowAllOrigins: true,
        AllowCredentials: true,
        AllowHeaders: []string{"Content-Type", "Authorization"},
    }))
    template.RegisterRoutes(server)
    return server
}
```

The `HttpBuildServer()` function wraps the Gin engine in an HTTP server:

```go
func HttpBuildServer() *http.Server {
    server := BuildServer()
    s := &http.Server{
        Addr:    ":8080",
        Handler: server,
    }
    return s
}
```

### 2. Container Initialization (`api/rest/template/container.go`)

The `NewContainer()` function initializes all dependencies using dependency injection:

```go
func NewContainer() *Container {
    // Initialize infrastructure
    database := db.NewDB()
    redis := db.NewRedis()
    
    // Initialize clients
    healthClient := healthClient.NewHealthClient()
    
    // Initialize services (with dependencies)
    healthService := healthService.NewHealthService(healthClient, database, redis)
    
    // Initialize managers (with services)
    healthManager := healthManager.NewHealthManager(healthService)
    
    // Initialize controllers (with managers)
    healthController := controller.NewHealthController(healthManager)
    
    return &Container{
        HealthController: healthController,
    }
}
```

**Dependency Chain:**
- `DB` and `Redis` → Infrastructure layer
- `HealthClient` → External client
- `HealthService` → Depends on `HealthClient`, `DB`, and `Redis`
- `HealthManager` → Depends on `HealthService`
- `HealthController` → Depends on `HealthManager`

### 3. Route Registration (`api/rest/template/apis.go`)

Routes are registered when `RegisterRoutes()` is called:

```go
func RegisterRoutes(router *gin.Engine) *gin.RouterGroup {
    container := NewContainer()
    v1 := router.Group("/v1")
    externalRoutes(v1, container)
    return v1
}

func externalRoutes(routerGroup *gin.RouterGroup, container *Container) {
    routerGroup.GET("/health", container.HealthController.HealthCheck)
}
```

## Usage

This template can be used as a starting point for new Go projects. Simply:

1. Copy the template structure
2. Replace the health check implementation with your own features
3. Replace mocked DB/Redis with real implementations when ready
4. Add new controllers, managers, services, and clients following the same pattern

## Benefits of This Architecture

- **Separation of Concerns**: Each layer has a single responsibility
- **Testability**: Easy to mock dependencies at each layer
- **Maintainability**: Clear structure makes code easy to navigate
- **Scalability**: Easy to add new features following the same pattern
- **Dependency Injection**: Dependencies are explicitly injected, making testing easier

