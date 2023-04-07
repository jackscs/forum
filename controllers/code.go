package controllers

//封装一些快捷状态吗。便于快速进行返回错误类型

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExit
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeInvalidArticle
	CodeArticleSuccess
	CodeInvalidAuthorID
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已经存在",
	CodeUserNotExit:     "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的token",
	CodeInvalidArticle:  "提交文章格式有误",
	CodeArticleSuccess:  "提交文章成功",
	CodeInvalidAuthorID: "获取用户id失败",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
