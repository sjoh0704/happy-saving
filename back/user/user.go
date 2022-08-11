package user

import (
	"fmt"
	"time"

	"github.com/sjoh0704/happysaving/post"
)

type User struct {
	ID        int64        `json:"id`
	Mail      string       `json:"mail"`
	Name      string       `json:"name"`
	Password  string       `json:"password"`
	Gender    Gender       `json:"gender"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Posts     []*post.Post `pg:"rel:has-many"`
}

func (u User) String() string {
	return fmt.Sprintf("User id: %d, mail: %s, name: %s", u.ID, u.Name, u.Mail)
}

func (u *User) SetPassword(pw string) {
	u.Password = pw
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetMail(mail string) {
	u.Mail = mail
}

func (u *User) UpdateTime() {
	u.UpdatedAt = time.Now()
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetMail() string {
	return u.Mail
}
