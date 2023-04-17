package models

import (
	"context"

	"github.com/zxcfer/newz/util"
	"gorm.io/gorm"
)

func GetArticleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Article))
}

type Article struct {
	Id              int `gorm:"primary_key"`
	Title           string
	Slug            string
	Image           string
	Description     string
	Content         string
	Link            string
	Viewed          int `gorm:"default:0"`
	WebsiteId       int
	WebsiteSlug     string
	IsUpdateContent int `gorm:"default:0"`
	Category        string
	gorm.Model
}

func (article *Article) TableName() string {
	return "articles"
}

type ArticleResponse struct {
	Title           string `json:"title"`
	Snippet         string `json:"snippet"`
	Slug            string `json:"slug"`
	Image           string `json:"image"`
	Link            string `json:"link"`
	UpdateAt        string `json:"update_at"`
	CreatedAt       string `json:"created_at"`
	IsUpdateContent int    `json:"is_update_content"`
	Tags            string `json:"tags"`
}
