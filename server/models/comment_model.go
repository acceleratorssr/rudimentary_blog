package models

type CommentModels struct {
	MODEL
	//ID              uint         `json:"id"`
	AuthorID        uint          `json:"author_id"`
	PostID          uint          `json:"post_id"`
	ParentCommentID uint          `json:"parent_id"`
	Post            ArticleModels `gorm:"foreignKey:PostID" json:"-"`
	Author          UserModels    `gorm:"foreignKey:AuthorID" json:"Author"`

	SubComments        []*CommentModels `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`
	ParentCommentModel *CommentModels   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model"`

	Text string `json:"text"`

	Like       int `gorm:"default:0" json:"like"`
	SubComment int `gorm:"default:0" json:"sub_comment"`
}

// GetID returns model ID
func (c *CommentModels) GetID() uint {
	return c.ID
}
