package common

import "errors"

var (
	ErrorNoRows = errors.New("Error No Rows!")

	ErrUsernameOrPasswordInvalid = NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrUsernameExisted = NewCustomError(
		errors.New("username has already been existed"),
		"username has already been existed",
		"ErrUsernameExisted",
	)
)
