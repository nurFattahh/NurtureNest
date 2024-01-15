package repository

import (
	"NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	"context"

	"gorm.io/gorm"
)

type forumDatabase struct {
	DB *gorm.DB
}

func NewForumRepository(DB *gorm.DB) interfaces.ForumRepository {
	return &forumDatabase{DB}
}

func (c *forumDatabase) PostForum(ctx context.Context, request domain.RequestPostForum, user_id uint) (domain.Forum, error) {
	var forum domain.Forum = domain.Forum{
		Title:   request.Title,
		Content: request.Content,
		UserID:  user_id,
		Likes:   0,
	}
	err := c.DB.Create(&forum).Error
	return forum, err
}

func (c *forumDatabase) GetAllForum(ctx context.Context) ([]domain.Forum, error) {
	var forums []domain.Forum
	err := c.DB.Find(&forums).Error
	return forums, err
}

func (c *forumDatabase) FindForumByTitle(ctx context.Context, title string) ([]domain.Forum, error) {
	var forums []domain.Forum
	err := c.DB.Where("title LIKE ?", "%"+title+"%").Find(&forums).Error
	return forums, err
}

func (c *forumDatabase) GetForumById(ctx context.Context, id uint) (domain.Forum, error) {
	var forums domain.Forum
	err := c.DB.Model(forums).Where("id =?", id).Preload("Comment").First(&forums).Error
	return forums, err
}

func (c *forumDatabase) PostComment(ctx context.Context, request domain.RequestPostComment, user_id uint) (domain.Comment, error) {
	var comment domain.Comment = domain.Comment{
		Comment: request.Comment,
		UserID:  user_id,
		ForumID: request.ForumID,
	}
	err := c.DB.Create(&comment).Error
	return comment, err
}

func (c *forumDatabase) CheckLikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.LikeForum, error) {
	var forums domain.LikeForum
	err := c.DB.Model(forums).Where("forum_id =? and user_id =?", forum_id, user_id).First(&forums).Error
	return forums, err
}

func (c *forumDatabase) PostLikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.LikeForum, error) {
	var forums domain.LikeForum = domain.LikeForum{
		UserID:  user_id,
		ForumID: forum_id,
		IsLiked: false,
	}
	err := c.DB.Model(forums).Create(&forums).Error
	return forums, err
}

func (c *forumDatabase) LikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.Forum, error) {
	var forum domain.Forum
	_ = c.DB.Model(&domain.LikeForum{}).Where("forum_id =? and user_id =?", forum_id, user_id).Update("is_liked", true).Error
	err := c.DB.Model(&domain.Forum{}).Where("id =?", forum_id).Update("likes", gorm.Expr("likes + ?", 1)).Find(&forum).Error
	return forum, err
}

func (c *forumDatabase) UnlikeForum(ctx context.Context, forum_id uint, user_id uint) (domain.Forum, error) {
	_ = c.DB.Model(&domain.LikeForum{}).Where("forum_id =? and user_id =?", forum_id, user_id).Update("is_liked", false).Error
	var forum domain.Forum
	err := c.DB.Model(&forum).Where("id =?", forum_id).Update("likes", gorm.Expr("likes - ?", 1)).Find(&forum).Error
	return forum, err
}
