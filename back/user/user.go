package user

import "time"

type UserInfo struct{
	Id          int64
	Mail 		string
	Name 		string
	Password    string
	CreatedTime time.Time
	UpdatedTime time.Time
}