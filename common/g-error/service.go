package g_error

import "errors"

var (
	ErrInvalidProToken = errors.New("invalid project token")
	ErrProjectLocked = errors.New("project locked")
	ErrCountMoreThanMax = errors.New("count more than max")
	ErrInvalidNotifyType = errors.New("invalid notify type")
)

var (
	ErrUsernameOrPasswordNotRight = errors.New("username or password not right")
	ErrInvalidUserToken = errors.New("invalid user token")
)
