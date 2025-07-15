package gamecore


import "errors"

var (
	errGameIdle = errors.New("Game is idle")
	errGameExceptionOccurs = errors.New("Unknown errors Happens")
)
