package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirasildev/student/api/models"
	// "github.com/mirasildev/student/storage/repo"
)

// @Router /student [post]
// @Summary Create a student
// @Description Create a student
// @Tags student
// @Accept json
// @Produce json
// @Param student body models.CreateStudent true "Student"
// @Success 201 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		req *models.CreateStudent
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Student().Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Successfully created",
	})
}
