package model

import (
	"fmt"
	"time"
)

type Post struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"user_id"`
	Author    *User     `pg:"rel:has-one"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p Post) String() string {
	return fmt.Sprintf("Post id: %d, user id: %d, content: %s, image url: %s", p.ID, p.AuthorID, p.Content, p.ImageURL)
}

func (p *Post) GetUserID() int64 {
	return p.AuthorID
}

func (p *Post) SetUserID(userID int64) {
	p.AuthorID = userID
}

func (p *Post) GetContent() string {
	return p.Content
}

func (p *Post) SetContent(content string) {
	p.Content = content
}

func (p *Post) GetImageURL() string {
	return p.ImageURL
}

func (p *Post) SetImageURL(imageURL string) {
	p.ImageURL = imageURL
}

func (p *Post) UpdateTime() {
	p.UpdatedAt = time.Now()
}
