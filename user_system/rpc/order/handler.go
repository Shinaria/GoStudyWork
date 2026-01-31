package main

import (
	"context"
	user "user_system/kitex_gen/GoStudyWork/userSystem/user"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// GetUserInfo implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetUserInfo(ctx context.Context, userid int64) (resp *user.UserInfo, err error) {
	// TODO: Your code here...
	
	return
}
