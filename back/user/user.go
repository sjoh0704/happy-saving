package user

import "fmt"


type User struct{
	Id          int64
	Mail 		string
	Name 		string
	Password    string
}

func (u User) String() string{
	return fmt.Sprintf("User<%d %s %s>", u.Id, u.Name, u.Mail)
}

