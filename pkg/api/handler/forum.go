package handler

import (
	sdk_jwt "NurtureNest/pkg/api/sdk/jwt"
	"NurtureNest/pkg/api/sdk/response"
	"NurtureNest/pkg/domain"
	services "NurtureNest/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ForumHandler struct {
	forumUseCase services.ForumUseCase
}

func NewForumHandler(usecase services.ForumUseCase) *ForumHandler {
	return &ForumHandler{
		forumUseCase: usecase,
	}
}

func (cr *ForumHandler) PostForum(c *gin.Context) {
	var request domain.RequestPostForum
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "BAD REQUEST", err)
		return
	}

	claimID, _ := sdk_jwt.ClaimToken(c)
	userID := uint(claimID)

	forum, err := cr.forumUseCase.PostForum(c.Request.Context(), request, userID)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Fail post forum", err)
		return
	}
	response.Success(c, http.StatusCreated, "post forum succeed", forum)
}

func (cr *ForumHandler) GetAllForum(c *gin.Context) {
	forums, err := cr.forumUseCase.GetAllForum(c.Request.Context())
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed get all forums", err)
		return
	}
	response.Success(c, http.StatusCreated, "Get all forums success", forums)
}

func (cr *ForumHandler) FindForumByTitle(c *gin.Context) {
	query := c.Query("title")

	forums, err := cr.forumUseCase.FindForumByTitle(c.Request.Context(), query)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed get all forums", err)
		return
	}
	response.Success(c, http.StatusCreated, "Get all forums success", forums)
}

func (cr *ForumHandler) GetForumById(c *gin.Context) {
	params := c.Param("id")

	parsedID, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "fail parsing id", err)
		return
	}
	id := uint(parsedID)
	forums, err := cr.forumUseCase.GetForumById(c.Request.Context(), id)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed get forum", err)
		return
	}
	response.Success(c, http.StatusCreated, "Get forum success", forums)

}

func (cr *ForumHandler) PostComment(c *gin.Context) {
	var request domain.RequestPostComment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "BAD REQUEST", err)
		return
	}

	userID, _ := sdk_jwt.ClaimToken(c)

	comment, err := cr.forumUseCase.PostComment(c.Request.Context(), request, uint(userID))
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Post comment failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "Post comment succeed", comment)
}

func (cr *ForumHandler) PostLike(c *gin.Context) {
	param := c.Param("forum_id")
	parsedForumID, _ := strconv.ParseUint(param, 10, 64)

	userID, _ := sdk_jwt.ClaimToken(c)

	check, err := cr.forumUseCase.CheckLikeForum(c.Request.Context(), uint(parsedForumID), uint(userID))
	if err != nil {
		_, err := cr.forumUseCase.PostLikeForum(c.Request.Context(), uint(parsedForumID), uint(userID))
		if err != nil {
			response.FailOrError(c, http.StatusBadRequest, "failed create like forum", err)
			return
		}
		check, err = cr.forumUseCase.CheckLikeForum(c.Request.Context(), uint(parsedForumID), uint(userID))
		if err != nil {
			response.FailOrError(c, http.StatusBadRequest, "failed check like forum", err)
			return
		}
	}

	if check.IsLiked {
		likes, err := cr.forumUseCase.UnlikeForum(c.Request.Context(), check.ForumID, check.UserID)
		if err != nil {
			response.FailOrError(c, http.StatusBadRequest, "unlike failed", err)
			return
		}
		response.Success(c, http.StatusCreated, "unlike succeed", likes)
		return
	} else {
		likes, err := cr.forumUseCase.LikeForum(c.Request.Context(), check.ForumID, check.UserID)
		if err != nil {
			response.FailOrError(c, http.StatusBadRequest, "like failed", err)
			return
		}
		response.Success(c, http.StatusCreated, "like succeed", likes)
	}

}
