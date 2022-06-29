package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ifont21/product-feedback-service/internal/domain/onboarding"
	"github.com/ifont21/product-feedback-service/internal/services"
	"github.com/labstack/echo/v4"
)

type projectPayload struct {
	domain.ProjectValue
	OwnerId uuid.UUID `json:"owner_id"`
}

type ProjectHttpHandler struct {
	projectService services.ProjectManagementService
}

func NewProjectHTTPHandler(service services.ProjectManagementService) ProjectHttpHandler {
	return ProjectHttpHandler{
		projectService: service,
	}
}

func (h ProjectHttpHandler) GetAllProjects(c echo.Context) error {
	projects, err := h.projectService.GetAllProjects("")

	fmt.Println("Get Projects ...")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, projects)
}

func (h ProjectHttpHandler) GetById(c echo.Context) error {
	return nil
}

func (h ProjectHttpHandler) CreateProject(c echo.Context) error {
	var payload projectPayload

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.projectService.CreateProject(payload.Name, payload.Description, payload.OwnerId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
