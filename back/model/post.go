package model

import (
	"fmt"
	"time"
)

type Post struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
	CoupleID  int64     `json:"couple_id"`
	Couple    *Couple   `json:"couple" pg:"rel:has-one"`
	Title     string	`json:"title"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p Post) String() string {
	return fmt.Sprintf("Post id: %d, user id: %d, content: %s, image url: %s", p.ID, p.AuthorID, p.Content, p.ImageURL)
}

func (p *Post) GetCoupleID() int64 {
	return p.CoupleID
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
