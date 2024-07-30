package response

import "time"

type UserResponse struct {
	Id 			string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	Email 		string
	Name 		string
	Surname 	string
	Lastname 	string
}

type ListUserResponseResponse struct {
	Users []*UserResponse
}
