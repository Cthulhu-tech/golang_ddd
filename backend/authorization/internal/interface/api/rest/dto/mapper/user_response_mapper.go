package mapper

import (
	"github.com/Cthulhu-tech/microservice/internal/application/common"
	"github.com/Cthulhu-tech/microservice/internal/interface/api/rest/dto/response"
)

func ToUserResponse(user *common.UserResult) *response.UserResponse {
	return &response.UserResponse{
		Id:			user.Id.String(),
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
		Email: 		user.Email,
		Name: 		user.Name,
		Surname: 	user.Surname,
		Lastname: 	user.Lastname,
	}
}

func ToUserListResponse(users []*common.UserResult) *response.ListUserResponseResponse {
	var responseList []*response.UserResponse

	for _, user := range users {
		responseList = append(responseList, ToUserResponse(user))
	}

	return &response.ListUserResponseResponse{Users: responseList}
}
