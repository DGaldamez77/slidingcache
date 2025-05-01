package util

import "github.com/joomcode/errorx"

var (
	ErrNoKeyFound        = errorx.IllegalState.New("key not found")
	ErrKeyAlreadyExpired = errorx.IllegalState.New("key already expired")
)
