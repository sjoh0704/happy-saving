package user

import (
	"fmt"
	"time"
)


type User struct{
	Id          int64 `json:"id`
	Mail 		string	`json:"mail"`
	Name 		string	`json:"name"`
	Password    string	`json:"password"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

func (u User) String() string{
	return fmt.Sprintf("User<%d %s %s>", u.Id, u.Name, u.Mail)
}
