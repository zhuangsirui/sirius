package webdav

import "errors"

var (
	ErrAuthenticatorNotExist = errors.New("authenticator not exist")
	ErrUserDirInaccessible   = errors.New("user directory inaccessible")
)
