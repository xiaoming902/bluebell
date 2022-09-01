package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000, "服务内部错误")
	InvalidParams = NewError(10001, "入参错误")

	UnauthorizedAuthNotExist  = NewError(10003, "账户不存在")
	UnauthorizedAuthFailed    = NewError(10004, "账户密码错误")
	UnauthorizedTokenError    = NewError(10005, "鉴权失败，Token 错误或丢失")
	UnauthorizedTokenTimeout  = NewError(10006, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10007, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10008, "请求过多")
)
