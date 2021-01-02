package directory

type RumorDirectory struct {
    received map[string]int
}

func NewRumorDirectory() *RumorDirectory {
    return &RumorDirectory{
        received: make(map[string]int),
    }
}

func (rd *RumorDirectory) RememberRumor(rumor string) {
    rd.received[rumor] = rd.received[rumor] + 1
}

func (rd *RumorDirectory) IsNewRumor(rumor string) bool {
    return rd.received[rumor] == 1
}

func (rd *RumorDirectory) GetRumorCount(rumor string) int {
    return rd.received[rumor]
}
