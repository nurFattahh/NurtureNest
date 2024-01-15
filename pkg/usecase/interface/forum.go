package interfaces

import (
	domain "NurtureNest/pkg/domain"
	"context"
)

type ForumUseCase interface {
	PostForum(ctx context.Context, request domain.RequestPostForum, user_id uint) (domain.Forum, error)
	GetAllForum(ctx context.Context) ([]domain.Forum, error)
	FindForumByTitle(ctx context.Context, title string) ([]domain.Forum, error)
	GetForumById(ctx context.Context, id uint) (domain.Forum, error)
	PostComment(ctx context.Context, request domain.RequestPostComment, user_id uint) (domain.Comment, error)
	CheckLikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.LikeForum, error)
	PostLikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.LikeForum, error)
	LikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.Forum, error)
	UnlikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.Forum, error)
}
