package directory

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "log"
    "sort"
    "strconv"
    "time"
)

type StatusDirectory struct {
    Busy   bool
    Ticker *time.Ticker
    fst    []*pb.Message
    snd    []*pb.Message
}

func NewStatusDirectory() *StatusDirectory {
    return &StatusDirectory{
    }
}

func (sd *StatusDirectory) AddStatus(message *pb.Message, expectedSize int) bool {
    if len(sd.fst) < expectedSize {
        sd.fst = append(sd.fst, message)
    } else if len(sd.snd) < expectedSize {
        sd.snd = append(sd.snd, message)
    }
    return len(sd.fst) == expectedSize && len(sd.snd) == expectedSize
}

func (sd *StatusDirectory) CheckStatesAndNumberOfMessages(selfSent, selfReceived int) bool {
    all := append(sd.fst, sd.snd...)

    // account for application messages the coordinator gets
    sumSent := selfSent * 2
    sumReceived := selfReceived * 2

    for _, message := range all {
        status := message.GetStatus()

        if status.GetState() == pb.Status_ACTIVE {
            log.Printf("Philosopher %s is active\n", message.GetSender())
            return false
        }

        sumSent += int(status.GetSent())
        sumReceived += int(status.GetReceived())
    }

    log.Println("No philosopher is active")
    log.Printf("Sum of sent messages: %d\n", sumSent)
    log.Printf("Sum of received messages: %d\n", sumReceived)

    return sumSent == sumReceived
}

func (sd *StatusDirectory) Restart() {
    sd.fst = nil
    sd.snd = nil
    sd.Ticker.Reset(1000 * time.Millisecond)
}

func (sd *StatusDirectory) GetAndPrintResults(selfPreferredTime int, selfId string) int {
    results := map[string]int{}

    // collect results in a map and track the common result
    results[selfId] = selfPreferredTime
    preferredTime := selfPreferredTime
    for _, message := range sd.snd {
        t := int(message.GetStatus().GetTime())
        results[message.GetSender()] = t
        if preferredTime != t {
            preferredTime = 0
        }
    }

    // sort
    keys := make([]string, 0, len(results))
    for k := range results {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(x, y int) bool {
        i, _ := strconv.Atoi(keys[x])
        j, _ := strconv.Atoi(keys[y])
        return i < j
    })

    // print
    for _, k := range keys {
        log.Println(k, "prefers", results[k])
    }
    if preferredTime == 0 {
        log.Println("The voted times do not match")
    } else {
        log.Printf("The voted time is %d\n", preferredTime)
    }

    return preferredTime
}
