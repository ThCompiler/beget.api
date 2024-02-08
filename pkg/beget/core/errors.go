package core

import (
	"github.com/pkg/errors"
)

var (
	ErrResolveBegetHost = errors.New("error resolve beget hostname : " + Host)
	ErrFromAPIMethod    = errors.New("beget api method return error")
)
