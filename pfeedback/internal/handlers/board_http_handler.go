package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	domain "github.com/ifont21/product-feedback-service/internal/domain/onboarding"
	"github.com/ifont21/product-feedback-service/internal/services"
	"github.com/labstack/echo/v4"
)

type BoardHttpHandler struct {
	boardService services.BoardManagementService
}

func NewBoardtHTTPHandler(service services.BoardManagementService) BoardHttpHandler {
	return BoardHttpHandler{
		boardService: service,
	}
}

func (h BoardHttpHandler) GetAllBoards(c echo.Context) error {
	boards, err := h.boardService.GetAllBoards("")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, boards)
}

func (h BoardHttpHandler) GetBoardById(c echo.Context) error {
	boardId := uuid.MustParse(c.Param("id"))
	board, err := h.boardService.GetById(boardId)

	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, errors.New("unable to retrieve the board"))
	}

	return c.JSON(http.StatusOK, board)
}

func (h BoardHttpHandler) CreateBoard(c echo.Context) error {
	var payload domain.BoardPayload

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	projectId := c.Param("projectId")

	err := h.boardService.CreateBoard(payload.Name, payload.Description, uuid.MustParse(projectId))

	if err != nil {
		c.Logger().Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
