package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirasildev/student/api/models"
	"github.com/mirasildev/student/storage/repo"
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

// @Router /students [get]
// @Summary Get all students
// @Description Get all students
// @Tags student
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllStudentsResponse
// @Failure 500 {object} models.ErrorResponse
func(h *handlerV1) GetAllStudents(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Student().GetAll(&repo.GetAllStudentsParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getStudentsResponse(result))
}

func getStudentsResponse(data *repo.GetAllStudentsResult) *models.GetAllStudentsResponse {
	response := models.GetAllStudentsResponse{
		Students: make([]*models.Student, 0),
		Count: data.Count,
	}

	for _, student := range data.Students {
		s := parseStudentModel(student)
		response.Students = append(response.Students, &s)
	}

	return &response
}

func parseStudentModel(student *repo.Student) models.Student {
	return models.Student{
		ID:              student.ID,
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		UserName:        student.UserName,
		PhoneNumber:     student.PhoneNumber,
		Email:           student.Email,
		CreatedAt:       student.CreatedAt,
	}
}
