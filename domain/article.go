package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Article struct {
	bun.BaseModel `bun:"table:articles,alias:a"`
	ID            int64     `bun:"id,pk,autoincrement"`
	AuthorID      int64     `bun:"author_id,notnull"`
	Title         string    `bun:"title,notnull"`
	Slug          string    `bun:"slug,notnull"`
	Description   string    `bun:"description,notnull"`
	Body          string    `bun:"body,notnull"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Tag struct {
	bun.BaseModel `bu:"table:tags,alias:t"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
}

type ArticleTag struct {
	bun.BaseModel `bun:"table:article_tags,alias:at"`
	ArticleID     int64 `bun:"article_id,notnull"`
	TagID         int64 `bun:"tag_id,notnull"`
}
