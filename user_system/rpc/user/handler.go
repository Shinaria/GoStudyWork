package main

import (
	"context"
	user "user_system/kitex_gen/GoStudyWork/userSystem/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, userid *user.UserID) (resp *user.UserInfo, err error) {
	// TODO: Your code here...
	return
}
