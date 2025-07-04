package common

import "errors"

var (
	ErrNoItems = errors.New("item must have at least one item")
)