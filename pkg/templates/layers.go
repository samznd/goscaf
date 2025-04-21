package templates

import (
	"fmt"
)

func RepositoryTemplate(orm string) string {
	switch orm {
	case "gorm":
		return `
package repositories

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetMessage() (string, error)
}

type RepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepoImpl{db: db}
}

func (r *RepoImpl) GetMessage() (string, error) {
	var result struct {
		Message string
	}
	if err := r.db.Raw("SELECT 'data from repository' AS message").Scan(&result).Error; err != nil {
		return "", err
	}
	return result.Message, nil
}
		`
	case "xorm":
		return `
package repositories

import (
	"xorm.io/xorm"
)

type Repository interface {
	GetMessage() (string, error)
}

type RepoImpl struct {
	engine *xorm.Engine
}

func NewRepository(engine *xorm.Engine) Repository {
	return &RepoImpl{engine: engine}
}

func (r *RepoImpl) GetMessage() (string, error) {
	// Example XORM usage
	result, err := r.engine.QueryString("SELECT 'data from repository'")
	if err != nil || len(result) == 0 {
		return ""
	}
	return result[0]["'data from repository'"]
}
	`
	case "ent":
		return `
package repositories

import (
	"entgo.io/ent/dialect"
	"context"
)

type Repository interface {
	GetMessage(ctx context.Context) string
}

type RepoImpl struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &RepoImpl{client: client}
}

func (r *RepoImpl) GetMessage(ctx context.Context) string {
	// You can customize this once your Ent schema is defined
	return "data from repository (ent)"
}	
	`
	}
	return `
package repositories

import "database/sql"

type Repository interface {
	GetMessage() (string, error)
}

type RepoImpl struct{
	db *sql.DB
}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (r *RepoImpl) GetMessage() (string, error) {
	return " data from repository"
}
`
}

func ServiceTemplate(projectName string) string {
	return fmt.Sprintf(`
package services

import "%s/internal/repositories"

type Service interface {
	GetMessage() (string, error)
}

type ServiceImpl struct {
	repo repositories.Repository
}

func NewService(r repositories.Repository) Service {
	return &ServiceImpl{repo: r}
}

func (s *ServiceImpl) GetMessage() (string, error) {
	return s.repo.GetMessage()
}

`, projectName)
}

func HandlerTemplate(projectName, orm, framework string) string {
	useContext := orm == "ent"

	// Shared logic
	handlerStruct := `
type Handler struct {
	service services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{service: s}
}
`

	// Based on context usage
	var callService string
	if useContext {
		callService = `message, err := h.service.GetMessage(context.Background())
	if err != nil {
		%s
	}
	return %s`
	} else {
		callService = `message, err := h.service.GetMessage()
	if err != nil {
		%s
	}
	return %s`
	}

	// Framework-specific handler body
	switch framework {
	case "fiber":
		return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/gofiber/fiber/v3"
	"context"
)

%s

func (h *Handler) Get(c *fiber.Ctx) error {
	%s
`, projectName, handlerStruct,
			fmt.Sprintf(callService,
				`return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})`,
				`c.JSON(fiber.Map{"message": message})`))

	case "gin":
		return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/gin-gonic/gin"
	"context"
)

%s

func (h *Handler) Get(c *gin.Context) {
	%s
`, projectName, handlerStruct,
			fmt.Sprintf(callService,
				`c.JSON(500, gin.H{"error": err.Error()}); return`,
				`c.JSON(200, gin.H{"message": message})`))

	case "echo":
		return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"context"
)

%s

func (h *Handler) Get(c echo.Context) error {
	%s
`, projectName, handlerStruct,
			fmt.Sprintf(callService,
				`return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})`,
				`c.JSON(http.StatusOK, map[string]string{"message": message})`))

	case "chi":
		return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"encoding/json"
	"net/http"
	"context"
)

%s

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	%s
`, projectName, handlerStruct,
			fmt.Sprintf(callService,
				`http.Error(w, err.Error(), http.StatusInternalServerError); return`,
				`json.NewEncoder(w).Encode(map[string]string{"message": message})`))

	case "iris":
		return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/kataras/iris/v12"
	"context"
)

%s

func (h *Handler) Get(ctx iris.Context) {
	%s
`, projectName, handlerStruct,
			fmt.Sprintf(callService,
				`ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(map[string]string{"error": err.Error()})
	return`,
				`ctx.JSON(map[string]string{"message": message})`))

	default:
		return "// Unsupported framework"
	}
}

func SetupRoutesTemplate(projectName, framework string) string {
	switch framework {
	case "fiber":
		return fmt.Sprintf(`
package routes

import (
	"%s/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app fiber.Router, h *handlers.Handler) {
	api := app.Group("/api")
	api.Get("/message", h.Get)
}
`, projectName)
	case "gin":
		return fmt.Sprintf(`
package routes

import (
	"%s/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handler) {
	api := router.Group("/api")
	api.GET("/message", h.Get)
}
`, projectName)
	case "echo":
		return fmt.Sprintf(`
package routes

import (
	"%s/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, h *handlers.Handler) {
	api := e.Group("/api")
	api.GET("/message", h.Get)
}
`, projectName)
	case "chi":
		return fmt.Sprintf(`
package routes

import (
	"%s/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRoutes(r chi.Router, h *handlers.Handler) {
	r.Route("/api", func(api chi.Router) {
		api.Get("/message", func(w http.ResponseWriter, r *http.Request) {
			h.Get(w, r)
		})
	})
}
`, projectName)
	case "iris":
		return fmt.Sprintf(`
package routes

import (
	"%s/internal/handlers"
	"github.com/kataras/iris/v12"
)

func SetupRoutes(app *iris.Application, h *handlers.Handler) {
	api := app.Party("/api")
	api.Get("/message", h.Get)
}
`, projectName)
	default:
		return `// ❌ Unsupported framework`
	}
}
