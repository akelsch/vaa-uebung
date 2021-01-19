package directory

//import "time"
//
//type ElectionDirectory struct {
//    Count        int
//    Color        Color
//    Initiator    string
//    Predecessor  string
//    VictoryTimer *time.Timer
//}
//
//type Color int
//
//const (
//    WHITE Color = iota
//    RED
//    GREEN
//)
//
//func NewElectionDirectory() *ElectionDirectory {
//    return &ElectionDirectory{
//        Count:       0,
//        Color:       WHITE,
//        Initiator:   "",
//        Predecessor: "",
//    }
//}
//
//func (ed *ElectionDirectory) Reset() {
//    ed.Count = 0
//    ed.Color = WHITE
//    ed.Initiator = ""
//    ed.Predecessor = ""
//    ed.VictoryTimer = nil
//}
//
//func (ed *ElectionDirectory) IsInitiator(selfId string) bool {
//    return ed.Initiator == selfId
//}
//
//func (ed *ElectionDirectory) IsCoordinator(selfId string) bool {
//    return ed.IsInitiator(selfId) && ed.Color == GREEN
//}
