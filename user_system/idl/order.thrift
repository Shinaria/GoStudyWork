namespace go GoStudyWork.userSystem.order

include "user.thrift"

service OrderService{
    user.UserInfo GetUserInfo(1: i64 userid)
}