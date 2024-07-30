package query

import "github.com/Cthulhu-tech/microservice/internal/application/common"


type UserQueryResult struct {
	Result *common.UserResult
}

type UserQueryListResult struct {
	Result []*common.UserResult
}
