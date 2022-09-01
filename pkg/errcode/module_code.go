package errcode

var (
	NoPermission = NewError(20007, "无权限执行该请求")
)
