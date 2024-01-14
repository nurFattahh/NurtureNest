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
	duration = result.String()

	timeResult := domain.GoalTracking{
		Goal:      duration,
		UserID:    uint(userID),
		LimitTime: time.Now().Add(time.Hour * 24),
		Status:    "In progress",
		IsActive:  true,
		Progress:  "0h0m0s",
		Result:    false,
	}

	timeResult, err = cr.goalTrackingUseCase.SetGoal(c.Request.Context(), timeResult, uint(userID))
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "fail set goal", err)
		return
	}

	response.Success(c, http.StatusCreated, "set goal success", timeResult)
}

func (cr *GoalTrackingHandler) CheckGoalProgress(c *gin.Context) {
	var request domain.RequestGoalResult
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "BAD REQUEST", err)
		return
	}

}

func (cr *GoalTrackingHandler) GoalResult(c *gin.Context) {
	var request domain.RequestGoalResult
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "BAD REQUEST", err)
		return
	}

	claimID, _ := sdk_jwt.ClaimToken(c)
	user_id := uint(claimID)

	err = cr.goalTrackingUseCase.SaveProgress(c.Request.Context(), request, user_id)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "fail save progress", err)
		return
	}

	goalTrack, err := cr.goalTrackingUseCase.GetGoalTracking(c.Request.Context(), user_id)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "fail get user goal track", err)
		return
	}

	progress, _ := time.ParseDuration(goalTrack.Progress)
	goal, _ := time.ParseDuration(goalTrack.Goal)
	fmt.Println(goalTrack.LimitTime)
	limit := time.Until(goalTrack.LimitTime)
	fmt.Println(limit)

	var result bool
	var status string

	if limit <= time.Hour*24 {
		status = "In Progress"
		result = false
	} else {
		if progress <= goal {
			status = "final"
			result = true
		} else {
			status = "final"
			result = false
		}
	}

	getResult, err := cr.goalTrackingUseCase.GoalResult(c.Request.Context(), goalTrack, user_id, status, result)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "fail get result", err)
		return
	}

	response.Success(c, http.StatusCreated, "success get goal result", getResult)

}
