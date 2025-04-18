package rest

import (
	"CourseService/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct{
	usecases *usecase.Usecase
	
}
func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{
		usecases: usecases,
	}
}


func (h *Handler) badRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (h *Handler) notFound(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

func (h *Handler) ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (h *Handler) created(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, data)
}

func (h *Handler) internalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}