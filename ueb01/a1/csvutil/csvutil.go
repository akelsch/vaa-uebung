package csvutil

import (
    "encoding/csv"
    "github.com/akelsch/vaa/ueb01/a1/errutil"
    "os"
    "path"
    "strings"
)

func Parse(filename string) [][]string {
    var filepath string
    if strings.HasPrefix(filename, ".") {
        pwd, _ := os.Getwd()
        filepath = path.Join(pwd, filename)
    } else {
        filepath = filename
    }

    file, err := os.Open(filepath)
    errutil.HandleError(err)

    rows, err := csv.NewReader(file).ReadAll()
    errutil.HandleError(err)

    return rows
}
