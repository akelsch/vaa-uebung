package directory

import "time"

type StatusDirectory struct {
    Ticker *time.Ticker
}

func NewStatusDirectory() *StatusDirectory {
    return &StatusDirectory{
    }
}
