package e

var MsgFlags = map[int]string{
	SUCCESS:                      "ok",
	ERROR:                        "fail",
	INVALID_PARAMS:               "请求参数错误",
	UNAUTHORIZED:                 "登录态已失效，请重新登录",
	ERROR_EXIST_DATA_FAIL:        "获取数据失败",
	ERROR_UPDATE_DATA_FAIL:       "更新数据失败",
	ERROR_READ_ONLY_NOT_EDIT:     "只读配置不能修改",
	ERROR_AUTH_CHECK_TOKEN_FAIL:  "认证失败",
	ERROR_AUTH_CHECK_ALL_SN_FAIL: "all sn auth failed",
	ERROR_AUTH_TOKEN_EXPIRED:     "登录态已失效，请重新登录",
	ERROR_AUTH_LOGIN_FAIL:        "登录失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
