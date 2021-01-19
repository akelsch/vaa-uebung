package directory

type FloodingDirectory struct {
    handled map[string]bool
    counter uint64
}

func NewFloodingDirectory() *FloodingDirectory {
    return &FloodingDirectory{
        handled: make(map[string]bool),
        counter: 1,
    }
}

func (fd *FloodingDirectory) IsHandled(identifier string) bool {
    return fd.handled[identifier] == true
}

func (fd *FloodingDirectory) MarkAsHandled(identifier string) {
    fd.handled[identifier] = true
}

func (fd *FloodingDirectory) NextSequence() uint64 {
    seq := fd.counter
    fd.counter++
    return seq
}
