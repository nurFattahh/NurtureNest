package handler

import (
	"NurtureNest/pkg/api/sdk/response"
	"NurtureNest/pkg/domain"
	"fmt"
	"net/http"
	"time"

	sdk_jwt "NurtureNest/pkg/api/sdk/jwt"
	services "NurtureNest/pkg/usecase/interface"

	"github.com/gin-gonic/gin"
)

type GoalTrackingHandler struct {
	goalTrackingUseCase services.GoalTrackingUseCase
}

func NewGoalTrackingHandler(usecase services.GoalTrackingUseCase) *GoalTrackingHandler {
	return &GoalTrackingHandler{
		goalTrackingUseCase: usecase,
	}
}

func (cr *GoalTrackingHandler) SetGoal(c *gin.Context) {
	var goal domain.RequestSetGoal
	err := c.ShouldBindJSON(&goal)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "BAD REQUEST", err)
		return
	}

	duration := fmt.Sprintf("%dh%dm%ds", goal.GoalHour, goal.GoalMinute, goal.GoalSecond)
	result, err := time.ParseDuration(duration)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "fail parsing duration", err)
		return
	}

	userID, _ := sdk_jwt.ClaimToken(c)

	timeResult := domain.SetGoal{
		Goal:      time.Now().Add(result),
		UserID:    uint(userID),
		LimitTime: time.Now().Add(time.Hour * 24),
	}

	SetGoal, err := cr.goalTrackingUseCase.SetGoal(c.Request.Context(), timeResult)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "set goal failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "set goal success", SetGoal)
}
