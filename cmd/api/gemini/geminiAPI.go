package gemini

import (
	"NurtureNest/pkg/api/sdk/response"
	"NurtureNest/pkg/config"
	"NurtureNest/pkg/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerativeAI(c *gin.Context, request domain.ChatBot) (*genai.GenerateContentResponse, error) {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed load config", configErr)
		return nil, configErr
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GEMINI_API_KEY))
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed Getting API", err)
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
	}
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Jadilah mentor mentalhealth saya"),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Baik, saya mentor mental health"),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(ctx, genai.Text(request.Question))
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed request question", err)
		return nil, err
	}
	return resp, nil
}
