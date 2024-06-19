package user

import "time"

type UserData struct {
	UserName  string
	Password  string
	BirthDate string
	CreatedAt time.Time
}
