package models

type CommentModel struct {
	MODEL
	//ID              uint         `json:"id"`
	AuthorID        uint         `json:"author_id"`
	PostID          uint         `json:"post_id"`
	ParentCommentID uint         `json:"parent_id"`
	Post            ArticleModel `gorm:"foreignKey:PostID" json:"-"`
	Author          UserModel    `gorm:"foreignKey:AuthorID" json:"Author"`

	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model"`

	Text string `json:"text"`

	Like       int `gorm:"default:0" json:"like"`
	SubComment int `gorm:"default:0" json:"sub_comment"`
}

// GetID returns model ID
func (c *CommentModel) GetID() uint {
	return c.ID
}
