package usecase

import (
	"NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	services "NurtureNest/pkg/usecase/interface"
	"context"
)

type forumUseCase struct {
	forumRepo interfaces.ForumRepository
}

func NewForumUseCase(repo interfaces.ForumRepository) services.ForumUseCase {
	return &forumUseCase{
		forumRepo: repo,
	}
}

func (c *forumUseCase) PostForum(ctx context.Context, request domain.RequestPostForum, user_id uint) (domain.Forum, error) {
	forum, err := c.forumRepo.PostForum(ctx, request, user_id)

	return forum, err
}

func (c *forumUseCase) GetAllForum(ctx context.Context) ([]domain.Forum, error) {
	forums, err := c.forumRepo.GetAllForum(ctx)
	return forums, err
}

func (c *forumUseCase) FindForumByTitle(ctx context.Context, title string) ([]domain.Forum, error) {
	forums, err := c.forumRepo.FindForumByTitle(ctx, title)
	return forums, err
}

func (c *forumUseCase) GetForumById(ctx context.Context, id uint) (domain.Forum, error) {
	forums, err := c.forumRepo.GetForumById(ctx, id)
	return forums, err
}

func (c *forumUseCase) PostComment(ctx context.Context, request domain.RequestPostComment, user_id uint) (domain.Comment, error) {
	comment, err := c.forumRepo.PostComment(ctx, request, user_id)
	return comment, err
}

func (c *forumUseCase) CheckLikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.LikeForum, error) {
	forum, err := c.forumRepo.CheckLikeForum(ctx, forum_id, user_id)
	return forum, err
}

func (c *forumUseCase) PostLikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.LikeForum, error) {
	forum, err := c.forumRepo.PostLikeForum(ctx, forum_id, user_id)
	return forum, err
}

func (c *forumUseCase) LikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.Forum, error) {
	forum, err := c.forumRepo.LikeForum(ctx, forum_id, user_id)
	return forum, err
}

func (c *forumUseCase) UnlikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.Forum, error) {
	forum, err := c.forumRepo.UnlikeForum(ctx, forum_id, user_id)
	return forum, err
}
