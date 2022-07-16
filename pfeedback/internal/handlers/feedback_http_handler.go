package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	domain "github.com/ifont21/product-feedback-service/internal/domain/feedback"
	"github.com/ifont21/product-feedback-service/internal/services"
	"github.com/labstack/echo/v4"
)

type FeedbackHttpHandler struct {
	feedbackService services.IFeedbackService
}

func NewFeedbackHttpHandler(service services.IFeedbackService) FeedbackHttpHandler {
	return FeedbackHttpHandler{
		feedbackService: service,
	}
}

func (h FeedbackHttpHandler) GetAllFeedbacks(c echo.Context) error {
	boardId := uuid.MustParse(c.Param("boardId"))
	feedbacks, err := h.feedbackService.GetAllFeedbacks(boardId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, feedbacks)
}

func (h FeedbackHttpHandler) GetFeedbackById(c echo.Context) error {
	postId, err := strconv.ParseInt(c.Param("id"), 0, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("error trying to get postId"))
	}

	feedback, err := h.feedbackService.GetFeedbackById(postId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, feedback)
}

func (h FeedbackHttpHandler) GetAllCategories(c echo.Context) error {
	categories, err := h.feedbackService.GetCategories()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, categories)
}

func (h FeedbackHttpHandler) CreatePost(c echo.Context) error {
	var payload domain.PostPayload
	boardId := uuid.MustParse(c.Param("boardId"))

	if boardId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("boardId is required"))
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.feedbackService.CreatePost(payload.Title, payload.Description, boardId, payload.CategoryID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("error creating post"))
	}

	return c.NoContent(http.StatusOK)
}

func (h FeedbackHttpHandler) UpdatePost(c echo.Context) error {
	var payload domain.PostUpdatePayload
	postId, err := strconv.ParseInt(c.Param("id"), 0, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("error trying to get postId"))
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	status, err := domain.ToStatusValue(payload.Status)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.feedbackService.UpdatePost(postId, payload.Title, payload.Description, status, payload.CategoryID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h FeedbackHttpHandler) UpVote(c echo.Context) error {
	postId, err := strconv.ParseInt(c.Param("id"), 0, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("error trying to get postId"))
	}

	err = h.feedbackService.UpVote(postId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)

}
