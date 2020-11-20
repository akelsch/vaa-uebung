package errutil

import "log"

func HandleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
