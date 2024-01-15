package domain

import "gorm.io/gorm"

type Forum struct {
	gorm.Model
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Creator Users     `json:"-" gorm:"foreignKey:UserID"`
	UserID  uint      `json:"user_id"`
	Comment []Comment `gorm:"foreignKey:UserID"`
	Likes   int       `json:"likes" gorm:"default:0"`
}

type RequestPostForum struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Comment struct {
	gorm.Model
	Comment string `json:"content"`
	UserID  uint   `json:"user_id"`
	ForumID uint   `json:"forum_id"`
}

type RequestPostComment struct {
	ForumID uint   `json:"forum_id"`
	Comment string `json:"comment"`
}

type LikeForum struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	ForumID uint `json:"forum_id"`
	IsLiked bool `json:"is_liked" gorm:"default:false"`
}

type RequestLike struct {
	UserID  uint `json:"user_id"`
	ForumID uint `json:"post_id"`
}
