package model

import (
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"text;not null" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"datetime;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	result := db.Debug().Model(&Post{}).Create(&p)
	if result.Error != nil {
		return &Post{}, result.Error
	}
	if p.ID != 0 {
		result = db.Debug().Select("id", "username", "email", "created_at", "updated_at").Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author)
		if result.Error != nil {
			return &Post{}, result.Error
		}
	}

	return p, nil
}
