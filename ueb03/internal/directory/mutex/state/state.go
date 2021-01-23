package state

type State int

const (
    RELEASED State = iota
    WANTED
    HELD
)
