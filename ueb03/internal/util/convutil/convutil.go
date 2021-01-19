package convutil

import (
    "github.com/akelsch/vaa/ueb03/internal/util/errutil"
    "strconv"
)

func StringToUint(str string) uint64 {
    num, err := strconv.ParseUint(str, 10, 0)
    errutil.HandleError(err)
    return num
}
