package mysql

import "errors"

var (
	ErrUserExist       = errors.New("用于已经存在")
	ErrUserNotExit     = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("用户名或密码错误")
	ErrInvalidID       = errors.New("id错误")
)
