package directory

import (
    "github.com/akelsch/vaa/ueb03/internal/directory/election/color"
    "time"
)

type ElectionDirectory struct {
    Count        int
    Color        color.Color
    Initiator    uint64
    Predecessor  uint64
    VictoryTimer *time.Timer
}

func NewElectionDirectory() *ElectionDirectory {
    return &ElectionDirectory{
        Count:       0,
        Color:       color.WHITE,
        Initiator:   0,
        Predecessor: 0,
    }
}

func (ed *ElectionDirectory) Reset() {
    ed.Count = 0
    ed.Color = color.WHITE
    ed.Initiator = 0
    ed.Predecessor = 0
    ed.VictoryTimer = nil
}

func (ed *ElectionDirectory) IsInitiator(selfId uint64) bool {
    return ed.Initiator == selfId
}

func (ed *ElectionDirectory) IsCoordinator(selfId uint64) bool {
    return ed.IsInitiator(selfId) && ed.Color == color.GREEN
}
