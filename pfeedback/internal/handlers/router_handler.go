package handlers

import (
	"database/sql"

	"github.com/ifont21/product-feedback-service/internal/infrastructure"
	"github.com/ifont21/product-feedback-service/internal/services"
	"github.com/labstack/echo/v4"
)

type RouterHandler struct {
	projectHttpHandler  ProjectHttpHandler
	boardHttpHandler    BoardHttpHandler
	feedbackHttpHandler FeedbackHttpHandler
}

func NewRouterHandleFactory(db *sql.DB) RouterHandler {
	projectService := services.NewProjectManagementService(infrastructure.NewProjectDBRepository(db))
	boardService := services.NewBoardManagementService(infrastructure.NewBoardDBRepository(db))
	feedbackService := services.NewFeedbackService(infrastructure.NewfeedbackDBRepository(db))

	return RouterHandler{
		projectHttpHandler:  NewProjectHTTPHandler(projectService),
		boardHttpHandler:    NewBoardtHTTPHandler(boardService),
		feedbackHttpHandler: NewFeedbackHttpHandler(feedbackService),
	}
}

func (h *RouterHandler) RegisterRoutes(v1 *echo.Group) {
	projects := v1.Group("/projects")
	projects.GET("", h.projectHttpHandler.GetAllProjects)
	projects.POST("", h.projectHttpHandler.CreateProject)

	projects.GET("/:projectId/boards", h.boardHttpHandler.GetAllBoards)
	projects.POST("/:projectId/boards", h.boardHttpHandler.CreateBoard)

	boards := v1.Group("/boards")
	boards.GET("/:id", h.boardHttpHandler.GetBoardById)

	boardFeedbacks := boards.Group("/:boardId/feedbacks")
	boardFeedbacks.GET("", h.feedbackHttpHandler.GetAllFeedbacks)
	boardFeedbacks.POST("", h.feedbackHttpHandler.CreatePost)

	categories := v1.Group("/categories")
	categories.GET("", h.feedbackHttpHandler.GetAllCategories)

	feedbacks := v1.Group("/feedbacks")
	feedbacks.GET("/:id", h.feedbackHttpHandler.GetFeedbackById)
	feedbacks.PUT("/:id", h.feedbackHttpHandler.UpdatePost)
	feedbacks.PATCH("/:id/vote", h.feedbackHttpHandler.UpVote)
}
