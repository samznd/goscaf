package templates

import "fmt"

func fiberHandler(projectName, orm string) string {
	ctxLine := ""
	call := "h.service.GetMessage()"
	if orm == "ent" {
		ctxLine = "context \"context\""
		call = "h.service.GetMessage(context.Background())"
	}

	return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/gofiber/fiber/v3"
	%s
)

type Handler struct {
	service services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	message, err := %s
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": message})
}
`, projectName, ctxLine, call)
}

func ginHandler(projectName, orm string) string {
	ctxArg := ""
	call := "h.service.GetMessage()"

	if orm == "ent" {
		ctxArg = "c.Request.Context()"
		call = "h.service.GetMessage(" + ctxArg + ")"
	}

	return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Get(c *gin.Context) {
	message, err := %s
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": message})
}
`, projectName, call)
}

func echoHandler(projectName, orm string) string {
	call := "h.service.GetMessage()"
	if orm == "ent" {
		call = "h.service.GetMessage(c.Request().Context())"
	}

	return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Get(c echo.Context) error {
	message, err := %s
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": message})
}
`, projectName, call)
}

func irisHandler(projectName, orm string) string {
	call := "h.service.GetMessage()"
	if orm == "ent" {
		call = "h.service.GetMessage(ctx.Request().Context())"
	}

	return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"github.com/kataras/iris/v12"
)

type Handler struct {
	service services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Get(ctx iris.Context) {
	message, err := %s
	if err != nil {
		ctx.StopWithStatus(500)
		return
	}
	ctx.JSON(iris.Map{"message": message})
}
`, projectName, call)
}

func chiHandler(projectName, orm string) string {
	call := "h.service.GetMessage()"
	if orm == "ent" {
		call = "h.service.GetMessage(r.Context())"
	}

	return fmt.Sprintf(`
package handlers

import (
	"%s/internal/services"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	message, err := %s
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
`, projectName, call)
}

func HandlerGenerator(projectName, framework, orm string) string {
	switch framework {
	case "fiber":
		return fiberHandler(projectName, orm)
	case "gin":
		return ginHandler(projectName, orm)
	case "echo":
		return echoHandler(projectName, orm)
	case "chi":
		return chiHandler(projectName, orm)
	case "iris":
		return irisHandler(projectName, orm)
	default:
		return ""
	}
}
