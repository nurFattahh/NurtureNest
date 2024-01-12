package handler

import (
	"NurtureNest/cmd/api/gemini"
	"NurtureNest/pkg/api/sdk/response"
	"NurtureNest/pkg/domain"
	services "NurtureNest/pkg/usecase/interface"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	sdk_jwt "NurtureNest/pkg/api/sdk/jwt"

	"github.com/gin-gonic/gin"
)

type ChatBotHandler struct {
	chatBotUseCase services.ChatBotUseCase
}

func NewChatBotHandler(usecase services.ChatBotUseCase) *ChatBotHandler {
	return &ChatBotHandler{
		chatBotUseCase: usecase,
	}
}

func (cr *ChatBotHandler) GenerateChatBot(c *gin.Context) {
	var request domain.ChatBot
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	resp, err := gemini.GenerativeAI(c, request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed generate ai", err)
		return
	}

	content := resp.Candidates[0].Content.Parts[0]

	userID, _ := sdk_jwt.ClaimToken(c)

	answer, err := json.Marshal(content)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed json to string", err)
		return
	}

	generate := domain.ChatBot{
		CreatedAt: time.Now(),
		Question:  request.Question,
		Answer:    string(answer),
		UserID:    uint(userID),
	}

	result, err := cr.chatBotUseCase.GenerateChatBot(c.Request.Context(), generate)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed generate chat", err)
		return
	}

	response.Success(c, http.StatusCreated, "Success generate chat", result)
}

func (cr *ChatBotHandler) GetAllHistoryChatBot(c *gin.Context) {
	user_id, _ := sdk_jwt.ClaimToken(c)
	user_id_uint := uint(user_id)

	var result []domain.ChatBot
	result, err := cr.chatBotUseCase.GetAllHistoryChatBot(c.Request.Context(), user_id_uint)
	fmt.Println(user_id_uint)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed get all chat", err)
		return
	}
	response.Success(c, http.StatusCreated, "Success get all chat", result)

}
