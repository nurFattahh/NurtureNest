package handler

import (
	"NurtureNest/pkg/domain"
	services "NurtureNest/pkg/usecase/interface"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TutorHandler struct {
	tutorUseCase services.TutorUseCase
}

type TutorResponse struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewTutorHandler(usecase services.TutorUseCase) *TutorHandler {
	return &TutorHandler{
		tutorUseCase: usecase,
	}
}

func (cr *TutorHandler) Save(c *gin.Context) {
	var user domain.Tutor

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := cr.tutorUseCase.Save(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}
}
