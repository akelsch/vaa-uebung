package directory

type FloodingDirectory struct {
    handled map[string]bool
}

func NewFloodingDirectory() *FloodingDirectory {
    return &FloodingDirectory{
        handled: make(map[string]bool),
    }
}

func (fd *FloodingDirectory) MarkAsHandled(identifier string) {
    fd.handled[identifier] = true
}
