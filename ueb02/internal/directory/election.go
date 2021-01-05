package directory

type ElectionDirectory struct {
    Count       int
    Color       Color
    Initiator   string
    Predecessor string
}

type Color int

const (
    WHITE Color = iota
    RED
    GREEN
)

func NewElectionDirectory() *ElectionDirectory {
    return &ElectionDirectory{
        Count:       0,
        Color:       WHITE,
        Initiator:   "",
        Predecessor: "",
    }
}

func (ed *ElectionDirectory) IsNotInitiator(selfId string) bool {
    return selfId != ed.Initiator
}
