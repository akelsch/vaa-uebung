package directory

import (
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/config"
)

type SnapshotDirectory struct {
    firstMarker bool
    Balance     int64
    Records     map[uint64]Record
    Responses   []*pb.Message
}

type Record struct {
    Changes   []int64
    Recording bool
}

func NewSnapshotDirectory() *SnapshotDirectory {
    return &SnapshotDirectory{
        firstMarker: true,
        Balance:     0,
        Records:     make(map[uint64]Record),
        Responses:   nil,
    }
}

func (sd *SnapshotDirectory) ChangesAsArray() []int64 {
    var changes []int64
    for _, record := range sd.Records {
        changes = append(changes, record.Changes...)
    }
    return changes
}

func (sd *SnapshotDirectory) IsFirstMarker() bool {
    firstMarker := sd.firstMarker
    sd.firstMarker = false
    return firstMarker
}

func (sd *SnapshotDirectory) RecordState(balance int64) {
    sd.Balance = balance
}

func (sd *SnapshotDirectory) MarkSenderAsEmpty(sender uint64) {
    record := &Record{Changes: nil, Recording: false}
    sd.Records[sender] = *record
}

func (sd *SnapshotDirectory) StartRecording(sender uint64, channels []*config.Node) {
    for _, channel := range channels {
        id := channel.Id
        if id != sender {
            record := &Record{Changes: nil, Recording: true}
            sd.Records[id] = *record
        }
    }
}

func (sd *SnapshotDirectory) StopRecording(sender uint64) {
    if record, ok := sd.Records[sender]; ok {
        record.Recording = false
        sd.Records[sender] = record
    }
}

func (sd *SnapshotDirectory) RecordChange(sender uint64, change int64) {
    if record, ok := sd.Records[sender]; ok {
        if record.Recording {
            record.Changes = append(record.Changes, change)
        }
        sd.Records[sender] = record
    }
}

func (sd *SnapshotDirectory) StoreResponse(message *pb.Message) {
    sd.Responses = append(sd.Responses, message)
}

func (sd *SnapshotDirectory) AreAllChannelsClosed(message *pb.Message) {
    sd.Responses = append(sd.Responses, message)
}

func (sd *SnapshotDirectory) Reset() {
    sd.firstMarker = true
    sd.Balance = 0
    sd.Records = make(map[uint64]Record)
    sd.Responses = nil
}
