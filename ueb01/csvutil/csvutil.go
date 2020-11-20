package csvutil

import (
    "encoding/csv"
    "github.com/akelsch/vaa/ueb01/errutil"
    "os"
    "path"
)

func Parse(filename string) [][]string {
    pwd, _ := os.Getwd()
    file, err := os.Open(path.Join(pwd, filename))
    errutil.HandleError(err)

    rows, err := csv.NewReader(file).ReadAll()
    errutil.HandleError(err)

    return rows
}
