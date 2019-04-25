package g_error

import "errors"

var (
	ErrUsernameCannotEmpty = errors.New("username can't be empty")
	ErrPasswordCannotEmpty = errors.New("password can't be empty")
)

var (
	ErrProjectNameCannotEmpty = errors.New("project name can't empty")
	ErrUserIDForProjectCannotEmpty = errors.New("user id for project can't empty")
)