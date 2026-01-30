namespace go GoStudyWork.userSystem.user

struct UserID{
    1: i64 no
}

struct UserInfo{
    1: UserID id
    2: string name
}

service UserService{
    UserInfo GetUserInfo (1: UserID userid)
}