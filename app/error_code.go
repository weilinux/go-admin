package app

// Predefined const error codes.
// message please see /resources/languages/*.ini

//goland:noinspection ALL
const (
	OK      = 0
	ERR     = 2
	SUCCESS = 200
	ERROR   = 500

	// code = 1000.. 用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WORNG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_TYPE     = 1007
	ERROR_USER_NOT_RIGHT = 1008

	baseNum = 10000

	ErrSvrNotAvailable = baseNum + 2100
	ErrServer          = baseNum + 2101
	ErrNoPermission    = baseNum + 2102
	ErrNotFound        = baseNum + 2104
	ErrNoResource      = baseNum + 2105
	ErrParams          = baseNum + 2106
	ErrMissingParams   = baseNum + 2107
	ErrNoRecord        = baseNum + 2108
	ErrOpFail          = baseNum + 2112
	ErrInvalidReq      = baseNum + 2113
	ErrInvalidParam    = baseNum + 2114

	ErrRepeatOp    = baseNum + 2202
	ErrDatabase    = baseNum + 2210
	ErrInvalidData = baseNum + 2211
	ErrCheckFail   = baseNum + 2212

	ErrDupRows    = baseNum + 2406
	ErrInsertFail = baseNum + 2401
	ErrUpdateFail = baseNum + 2403
	ErrDeleteFail = baseNum + 2404
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_PASSWORD_WORNG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_USER_NOT_RIGHT: "无权登录",
	ERROR_TOKEN_EXIST:    "Token鉴权失败",
	ERROR_TOKEN_RUNTIME:  "Token已过期",
	ERROR_TOKEN_WRONG:    "Token不正确",
	ERROR_TOKEN_TYPE:     "Token格式错误",
}

// GetErrMsg 输出错误信息返回
func GetErrMsg(code int) string {
	return codeMsg[code]
}
