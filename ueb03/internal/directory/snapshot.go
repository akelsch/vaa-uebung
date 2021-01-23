package directory

type SnapshotDirectory struct {
    State    *State
    outgoing map[uint64][]int64
    incoming map[uint64][]int64
}

type State struct {
    Balance int64
    Changes []int64
}

func NewSnapshotDirectory(initialBalance int64) *SnapshotDirectory {
    return &SnapshotDirectory{
        State: &State{
            Balance: initialBalance,
            Changes: nil,
        },
        outgoing: make(map[uint64][]int64),
        incoming: make(map[uint64][]int64),
    }
}

func (sd *SnapshotDirectory) AddChange(c int64) {
    sd.State.Changes = append(sd.State.Changes, c)
}

//func (sd *SnapshotDirectory) Reset() {
//    sd.incoming = make(map[uint64][]int64)
//    sd.outgoing = make(map[uint64][]int64)
//}
